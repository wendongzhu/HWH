package tool

import (
	"github.com/wendongzhu/hwh/backend/service/global"
	"github.com/wendongzhu/hwh/backend/service/param"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func InitConfigFile() {
	// 1. 配置文件默认路径
	cfgFilePath := global.ConfigFilePath
	cfgFileName := global.ConfigFileName
	// 2. 配置文件默认
	defaultConfig := param.ConfigFile{
		ArchiveData: param.Archive{
			FileName: "CustomerArchive.xcarch",
			FilePath: "D:/Harms+Wende/SharedData/Archive/",
		},
		ExportData: param.Export{
			FileName: "HWH",
			FilePath: "D:/Harms+Wende/ExportData/",
		},
		SamplingTime: param.SamplingTime{
			Time: 10,
		},
	}
	// 3. 如果不存在则创建配置文件
	exists, err := PathExists(cfgFilePath)
	if err != nil {
		return
	}
	if exists {
		// 3. 如果不存在则创建配置文件
		_, err := os.Stat(cfgFilePath + cfgFileName)
		if os.IsNotExist(err) {
			data, err := yaml.Marshal(&defaultConfig)
			if err != nil {
				return
			}
			err = ioutil.WriteFile(cfgFilePath+cfgFileName, data, 0644)
			if err != nil {
				return
			}
		}
		// 4. 如果存在则读取配置文件
		ReadConfigFile(cfgFilePath + cfgFileName)
	}

}

func ReadConfigFile(configFilePath string) {
	_, err := os.Stat(global.ConfigFilePath)
	if !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return
		}

		var cfg param.ConfigFile
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			return
		}

		global.ModelFile.ArchiveFilePath = cfg.ArchiveData.FilePath
		global.ModelFile.ArchiveFileName = cfg.ArchiveData.FileName
		global.ModelFile.ExportFilePath = cfg.ExportData.FilePath
		global.ModelFile.ExportFileName = cfg.ExportData.FileName
		global.ModelFile.SamplingTime = cfg.SamplingTime.Time
	}

}

func WriteConfigFile(configFilePath string, cfg param.ConfigFile) {
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return
	}
}
