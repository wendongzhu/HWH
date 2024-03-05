package param

type SelectDataParam struct {
	FileName    string `json:"file_name"`
	PageSize    int    `json:"page_size"`
	CurrentPage int    `json:"current_page"`
}
