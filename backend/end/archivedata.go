package end

import (
	"github.com/wendongzhu/hwh/backend/service/global"
	"github.com/wendongzhu/hwh/backend/service/param"
	tool "github.com/wendongzhu/hwh/backend/service/tools"
)

func (b *Backend) ReadArchiveFileAPI(selectParam param.SelectDataParam) (res param.HWHArchiveDataListParam) {
	ha := tool.HWHArchiveTool{}
	res = ha.ReadArchiveFile(selectParam)
	return res
}

func (b *Backend) ReadArchiveFileNameAPI() (res param.FileNameListParam) {
	res = tool.ReadFileName(global.ModelFile.ArchiveFilePath+"/", "xcarch")
	return res
}

func (b *Backend) WriteExportFileAPI(selectParam param.SelectDataParam) (res param.ExportListParam) {
	res = tool.ReadExcelFile(selectParam)
	return res
}
