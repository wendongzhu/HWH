package tool

import (
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/wendongzhu/hwh/backend/service/global"
	"github.com/wendongzhu/hwh/backend/service/param"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type HWHArchiveTool struct {
}

func (ha *HWHArchiveTool) ArchiveFileWatch() {
	time.Sleep(time.Second * 5)
	filePath := global.ModelFile.ArchiveFilePath + "/" + global.ModelFile.ArchiveFileName
	configTail := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // true:一直监听(同tail -f) false:一次后即结束
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // true:文件不存在即退出 false:文件不存在即等待
		Poll:      true,
	}
	tails, err := tail.TailFile(filePath, configTail)
	if err != nil {
		//Logger("sys").Error().Msg("Tail file failed, err:" + err.Error())
		return
	}

	var (
		line *tail.Line
		ok   bool
	)

	//Logger("sys").Info().Msg("HWH Archive File Watch!")

	for {

		line, ok = <-tails.Lines
		if !ok {
			//Logger("sys").Error().Msg("tail file close reopen, filename: " + tails.Filename)
			time.Sleep(time.Second)
			continue
		}

		str1 := strings.Replace(line.Text, "/", " ", -1)
		str2 := strings.Replace(str1, "| |", "|_|", -1)
		str3 := strings.Replace(str2, "|", " ", -1)
		items := strings.Fields(str3)

		if len(items) > 15 {
			archiveData := ha.archiveParse(items)
			SamplingTime := global.ModelFile.SamplingTime
			if len(archiveData.IKV) > SamplingTime {
				exportData := param.ExportParam{
					ModelGroup:        archiveData.ModuleGroup,
					ModelName:         archiveData.ModuleName,
					ProgramNumber:     archiveData.ProgramNumber,
					WeldTime:          archiveData.SpotTime,
					SamplingTime:      SamplingTime,
					ExtractCurrent:    int(archiveData.IKV[SamplingTime]),
					ExtractVoltage:    int(archiveData.UKV[SamplingTime]),
					ExtractResistance: int(archiveData.RKV[SamplingTime]),
					PeakCurrent:       ha.getMax(archiveData.IKV),
					PeakVoltage:       ha.getMax(archiveData.UKV),
					PeakResistance:    ha.getMax(archiveData.RKV),
					ArchiveSeries:     archiveData.ArchiveSeries,
				}
				WriteExcelFile(exportData)
			}
		} else {
			//Logger("sys").Warn().Any("归档数据长度小于10，数据为：", items)
		}
	}

}

func (ha *HWHArchiveTool) ReadArchiveFile(selectParam param.SelectDataParam) (res param.HWHArchiveDataListParam) {
	fileName := selectParam.FileName
	pageSize := selectParam.PageSize
	currentPage := selectParam.CurrentPage
	filePath := global.ModelFile.ArchiveFilePath + "/" + fileName
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("read file fail", err)
		return
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return
	}

	for index, line := range strings.Split(string(fd), "\n") {
		str1 := strings.Replace(line, "/", " ", -1)
		str2 := strings.Replace(str1, "| |", "|_|", -1)
		str3 := strings.Replace(str2, "|", " ", -1)
		items := strings.Fields(str3)

		if len(items) > 15 && index > 1 && index > pageSize*(currentPage-1) && index <= pageSize*currentPage {
			archiveData := ha.archiveParse(items)
			res.Data = append(res.Data, archiveData)
		}
		res.Count = index
	}

	return res
}

func (ha *HWHArchiveTool) archiveParse(slice []string) (ap param.HWHArchiveParam) {
	ap.SpotTime = slice[4] + " " + slice[3]
	ap.ArchiveSeries = slice[5]
	ap.ModuleGroup = slice[0]
	ap.ModuleName = slice[1]

	if i := ha.find(slice, "14409"); i != -1 {
		ap.TypeId = slice[i+1]
	} else {
		ap.TypeId = "0"
	}
	if i := ha.find(slice, "14408"); i != -1 {
		ap.SpotNumber = slice[i+1]
	} else {
		ap.SpotNumber = "0"
	}
	if i := ha.find(slice, "3528"); i != -1 {
		ap.ProgramNumber = slice[i+1]
	} else {
		ap.ProgramNumber = "0"
	}
	if i := ha.find(slice, "10 "); i != -1 {
		ap.SpotName = slice[i+1]
	} else {
		ap.SpotName = "0"
	}
	if i := ha.find(slice, "14049"); i != -1 {
		switch slice[i+1] {
		case "1":
			ap.ControlMode = "SKT"
		case "2":
			ap.ControlMode = "KSR"
		case "4":
			ap.ControlMode = "IQR"
		case "5":
			ap.ControlMode = "AMC/DCM"
		case "9":
			ap.ControlMode = "AMF"
		default:
			ap.ControlMode = "0"
		}
	} else {
		ap.ControlMode = "0"
	}

	if i := ha.find(slice, "8077"); i != -1 {
		ap.SetCurrent = slice[i+1]
	} else {
		ap.SetCurrent = "0"
	}
	if i := ha.find(slice, "14181"); i != -1 {
		ap.RealMeanCurrent = slice[i+1]
	} else {
		ap.RealMeanCurrent = "0"
	}
	if i := ha.find(slice, "14185"); i != -1 {
		ap.RealMeanVoltage = slice[i+1]
	} else {
		ap.RealMeanVoltage = "0"
	}
	if i := ha.find(slice, "151"); i != -1 {
		ap.PreTime = slice[i+1]
	} else {
		ap.PreTime = "0"
	}
	if i := ha.find(slice, "152"); i != -1 {
		ap.KeepTime = slice[i+1]
	} else {
		ap.KeepTime = "0"
	}
	if i := ha.find(slice, "14219"); i != -1 {
		ap.SetTime = slice[i+1]
	} else {
		ap.SetTime = "0"
	}
	if i := ha.find(slice, "14219"); i != -1 {
		ap.RealTime = slice[i+1]
	} else {
		ap.RealTime = "0"
	}
	if i := ha.find(slice, "8073"); i != -1 {
		ap.RealEnergy = slice[i+1]
	} else {
		ap.RealEnergy = "0"
	}
	if i := ha.find(slice, "608"); i != -1 {
		ap.MillingCycleCounter = slice[i+1]
	} else {
		ap.MillingCycleCounter = "0"
	}
	if i := ha.find(slice, "14031"); i != -1 {
		ap.MillingWeldSpotCounter = slice[i+1]
	} else {
		ap.MillingWeldSpotCounter = "0"
	}
	if i := ha.find(slice, "14030"); i != -1 {
		ap.WeldSpotCounter = slice[i+1]
	} else {
		ap.WeldSpotCounter = "0"
	}
	if i := ha.find(slice, "8041"); i != -1 {
		ap.RMax = slice[i+1]
	} else {
		ap.RMax = "0"
	}

	if i := ha.find(slice, "14287"); i != -1 {
		ap.QValue = slice[i+1]
	} else {
		ap.QValue = "0"
	}
	if i := ha.find(slice, "14285"); i != -1 {
		ap.QThreshold = slice[i+1]
	} else {
		ap.QThreshold = "0"
	}

	if ap.QValue >= ap.QThreshold {
		ap.SpotQuality = "OK"
	} else {
		ap.SpotQuality = "NOK"
	}

	if i := ha.find(slice, "14459 "); i != -1 {
		ap.Spatter = slice[i+1]
	} else {
		ap.Spatter = "0"
	}

	if i := ha.find(slice, "3529"); i != -1 {
		ap.GunNumber = slice[i+1]
	} else {
		ap.GunNumber = "0"
	}
	if i := ha.find(slice, "625"); i != -1 {
		ap.GunName = slice[i+1]
	} else {
		ap.GunName = "0"
	}
	if i := ha.find(slice, "14512"); i != -1 {
		if slice[i+1] != "_" {
			ap.ErrorCode = slice[i+1]
		}
	} else {
		ap.ErrorCode = "0"
	}

	if i := ha.find(slice, "14095"); i != -1 {
		ap.SKT = ha.curveParse(slice[i+1])
	} else {
		ap.SKT = []float32{}
	}
	if i := ha.find(slice, "14055"); i != -1 {
		ap.IKV = ha.curveParse(slice[i+1])
	} else {
		ap.IKV = []float32{}
	}
	if i := ha.find(slice, "14057"); i != -1 {
		ap.UKV = ha.curveParse(slice[i+1])
	} else {
		ap.UKV = []float32{}
	}
	if i := ha.find(slice, "14740"); i != -1 {
		ap.RKV = ha.curveParse(slice[i+1])
	} else {
		ap.RKV = []float32{}
	}
	if i := ha.find(slice, "14281"); i != -1 {
		ap.QKV = ha.curveParse(slice[i+1])
	} else {
		ap.QKV = []float32{}
	}

	return ap
}

func (ha *HWHArchiveTool) find(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func (ha *HWHArchiveTool) curveParse(data string) []float32 {
	var i uint64
	var measuredValue float32
	var result []float32
	if len(data) > 50 && data != "null" && data != "0" && data != "000" {
		factor, _ := strconv.ParseUint(string(data[30])+string(data[31])+string(data[28])+string(data[29])+
			string(data[26])+string(data[27])+string(data[24])+string(data[25]), 16, 0)
		sampleNum, _ := strconv.ParseUint(string(data[34])+string(data[35])+string(data[32])+string(data[33]),
			16, 0)
		for range data {
			if i+1 < sampleNum {
				value, _ := strconv.ParseUint(string(data[38+i*4])+string(data[39+i*4])+string(data[36+i*4])+
					string(data[37+i*4]), 16, 0)
				measuredValue = (float32(value) * float32(factor)) / float32(65536)
				result = append(result, measuredValue)
			}
			i++
		}
	} else {
		result = append(result, 0)
	}

	return result
}

func (ha *HWHArchiveTool) getMax(arr []float32) int {
	if len(arr) == 0 {
		return 0
	}

	maxNum := arr[0]
	for _, num := range arr {
		if num > maxNum {
			maxNum = num
		}
	}

	return int(maxNum)
}
