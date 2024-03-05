package end

import (
	"context"
	tool "github.com/wendongzhu/hwh/backend/service/tools"
	"time"
)

type RealTime struct {
	ctx context.Context
}

func init() {
	tool.InitConfigFile()

	ah := tool.HWHArchiveTool{}
	go ah.ArchiveFileWatch()
}

func NewRealTime() *RealTime {
	return &RealTime{}
}

func (rt *RealTime) OnDomReady(ctx context.Context) {
	rt.ctx = ctx
	for {
		//if global.IsModbusConn {
		//	runtime.EventsEmit(ctx, "device/status", true)
		//} else {
		//	runtime.EventsEmit(ctx, "device/status", false)
		//}
		time.Sleep(time.Second)

	}

}
