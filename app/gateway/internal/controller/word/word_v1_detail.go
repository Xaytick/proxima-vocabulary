package word

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"proxima-vocabulary/app/gateway/api/word/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
