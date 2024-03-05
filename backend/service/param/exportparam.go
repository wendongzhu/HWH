package param

type ExportListParam struct {
	Data  []ExportParam `json:"data"`
	Count int           `json:"count"`
}

type ExportParam struct {
	ModelGroup        string `json:"model_group"`
	ModelName         string `json:"model_name"`
	ProgramNumber     string `json:"program_number"`
	WeldTime          string `json:"weld_time"`
	SamplingTime      int    `json:"sampling_time"`
	ExtractCurrent    int    `json:"extract_current"`
	ExtractVoltage    int    `json:"extract_voltage"`
	ExtractResistance int    `json:"extract_resistance"`
	PeakCurrent       int    `json:"peak_current"`
	PeakVoltage       int    `json:"peak_voltage"`
	PeakResistance    int    `json:"peak_resistance"`
	ArchiveSeries     string `json:"archive_series"`
}
