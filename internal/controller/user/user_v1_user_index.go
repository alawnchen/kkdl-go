package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) UserIndex(ctx context.Context, req *v1.UserIndexReq) (res *v1.UserIndexRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
