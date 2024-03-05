package end

import (
	"github.com/wendongzhu/hwh/backend/service/global"
	"github.com/wendongzhu/hwh/backend/service/param"
	tool "github.com/wendongzhu/hwh/backend/service/tools"
)

func (b *Backend) CollectionDataAPI() interface{} {

	return ""
}

func (b *Backend) ReadFileNameAPI() (res param.FileNameListParam) {
	res = tool.ReadFileName(global.ModelFile.ExportFilePath+"/", "xlsx")
	return res
}

func (b *Backend) ReadExportFileAPI(selectParam param.SelectDataParam) (res param.ExportListParam) {
	res = tool.ReadExcelFile(selectParam)
	return res
}
