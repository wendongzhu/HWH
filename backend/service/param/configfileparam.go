package param

type ConfigFile struct {
	ArchiveData  Archive      `yaml:"ArchiveData"`
	ExportData   Export       `yaml:"ExportData"`
	SamplingTime SamplingTime `yaml:"SamplingTime"`
}

type Archive struct {
	FileName string `yaml:"FileName"`
	FilePath string `yaml:"FilePath"`
}

type Export struct {
	FileName string `yaml:"FileName"`
	FilePath string `yaml:"FilePath"`
}

type SamplingTime struct {
	Time int `yaml:"Time"`
}
