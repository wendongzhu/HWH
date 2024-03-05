package end

import (
	"fmt"
	"github.com/wendongzhu/hwh/backend/service/global"
	"github.com/wendongzhu/hwh/backend/service/param"
	tool "github.com/wendongzhu/hwh/backend/service/tools"
)

func (b *Backend) ReadSystemSettingAPI() param.ModelFileParam {
	var modelFileParam param.ModelFileParam
	tool.ReadConfigFile(global.ConfigFilePath + global.ConfigFileName)

	modelFileParam.ArchiveFilePath = global.ModelFile.ArchiveFilePath
	modelFileParam.ArchiveFileName = global.ModelFile.ArchiveFileName
	modelFileParam.ExportFilePath = global.ModelFile.ExportFilePath
	modelFileParam.ExportFileName = global.ModelFile.ExportFileName
	modelFileParam.SamplingTime = global.ModelFile.SamplingTime
	return modelFileParam
}

func (b *Backend) WriteSystemSettingAPI(data param.ModelFileParam) param.ModelFileParam {
	var modelFileParam param.ModelFileParam
	global.ModelFile.ArchiveFilePath = data.ArchiveFilePath
	global.ModelFile.ArchiveFileName = data.ArchiveFileName
	global.ModelFile.ExportFilePath = data.ExportFilePath
	global.ModelFile.ExportFileName = data.ExportFileName
	global.ModelFile.SamplingTime = data.SamplingTime

	fmt.Println("modelFile data", data)
	tool.WriteConfigFile(global.ConfigFilePath+global.ConfigFileName, param.ConfigFile{
		ArchiveData: param.Archive{
			FileName: data.ArchiveFileName,
			FilePath: data.ArchiveFilePath,
		},
		ExportData: param.Export{
			FileName: data.ExportFileName,
			FilePath: data.ExportFilePath,
		},
		SamplingTime: param.SamplingTime{
			Time: data.SamplingTime,
		},
	})

	modelFileParam = b.ReadSystemSettingAPI()
	fmt.Println("modelFileParam", modelFileParam)
	return modelFileParam
}
