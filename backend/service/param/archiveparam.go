package param

type HWHArchiveDataListParam struct {
	Data  []HWHArchiveParam `json:"data"`
	Count int               `json:"count"`
}

type HWHArchiveParam struct {
	SpotTime               string    `json:"spot_time"`
	ModuleGroup            string    `json:"model_group"`
	ModuleName             string    `json:"model_name"`
	TypeId                 string    `json:"type_id"`
	SpotNumber             string    `json:"spot_number"`
	ProgramNumber          string    `json:"program_number"`
	SpotName               string    `json:"spot_name"`
	ControlMode            string    `json:"control_mode"`
	SetCurrent             string    `json:"set_current"`
	RealMeanCurrent        string    `json:"real_mean_current"`
	RealMeanVoltage        string    `json:"real_mean_voltage"`
	PreTime                string    `json:"pre_time"`
	KeepTime               string    `json:"keep_time"`
	SetTime                string    `json:"set_time"`
	RealTime               string    `json:"real_time"`
	RealEnergy             string    `json:"real_energy"`
	MillingCycleCounter    string    `json:"milling_cycle_counter"`
	MillingWeldSpotCounter string    `json:"milling_weld_spot_counter"`
	WeldSpotCounter        string    `json:"weld_spot_counter"`
	RMax                   string    `json:"r_max"`
	QValue                 string    `json:"q_value"`
	QThreshold             string    `json:"q_threshold"`
	SpotQuality            string    `json:"spot_quality"`
	Spatter                string    `json:"spatter"`
	GunNumber              string    `json:"gun_number"`
	GunName                string    `json:"gun_name"`
	ErrorCode              string    `json:"error_code"`
	ArchiveSeries          string    `json:"archive_series"`
	SKT                    []float32 `json:"skt"`
	IKV                    []float32 `json:"ikv"`
	UKV                    []float32 `json:"ukv"`
	RKV                    []float32 `json:"rkv"`
	QKV                    []float32 `json:"qkv"`
}
