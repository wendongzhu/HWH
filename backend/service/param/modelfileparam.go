package param

type ModelFileParam struct {
	ArchiveFilePath string `json:"ArchiveFilePath"`
	ArchiveFileName string `json:"ArchiveFileName"`
	ExportFilePath  string `json:"ExportFilePath"`
	ExportFileName  string `json:"ExportFileName"`
	SamplingTime    int    `json:"SamplingTime"`
}

type FileNameListParam struct {
	Data []FileNameParam `json:"data"`
}
type FileNameParam struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
