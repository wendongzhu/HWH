package end

import (
	"context"
)

func init() {
	//rs := drives.Modbus{}
	//if !global.IsModbusDriveStatus {
	//	go rs.Start()
	//}

}

type Backend struct {
	ctx context.Context
}

func NewBackend() *Backend {
	return &Backend{}
}
