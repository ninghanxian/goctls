package {{.packageName}}

import (
	"context"

	"{{.projectPath}}{{.importPrefix}}/internal/svc"
	"{{.projectPath}}{{.importPrefix}}/internal/types"
	"{{.projectPath}}{{.importPrefix}}/internal/utils/dberrorhandler"

{{if .useI18n}}    "github.com/qmcloud/admin-common/i18n"
{{else}}    "github.com/qmcloud/admin-common/msg/errormsg"
{{end}}{{if .useUUID}}    "github.com/qmcloud/admin-common/utils/uuidx"
{{end}}
{{if .HasPointy}}	"github.com/qmcloud/admin-common/utils/pointy"
{{end}}	"github.com/zeromicro/go-zero/core/logx"
)

type Get{{.modelName}}ByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGet{{.modelName}}ByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{.modelName}}ByIdLogic {
	return &Get{{.modelName}}ByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Get{{.modelName}}ByIdLogic) Get{{.modelName}}ById(req *types.{{if .useUUID}}UU{{end}}ID{{.IdType}}Req) (*types.{{.modelName}}InfoResp, error) {
	data, err := l.svcCtx.DB.{{.modelName}}.Get(l.ctx, {{if .useUUID}}uuidx.ParseUUIDString({{end}}req.Id{{if .useUUID}}){{end}})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.{{.modelName}}InfoResp{
	    BaseDataInfo: types.BaseDataInfo{
            Code: 0,
            Msg:  {{if .useI18n}}l.svcCtx.Trans.Trans(l.ctx, i18n.Success){{else}}errormsg.Success{{end}},
        },
        Data: types.{{.modelName}}Info{
{{if .HasCreated}}            Base{{if .useUUID}}UU{{end}}ID{{.IdType}}Info:    types.Base{{if .useUUID}}UU{{end}}ID{{.IdType}}Info{
				Id:          {{if .useUUID}}pointy.GetPointer(data.ID.String()){{else}}&data.ID{{end}},
				CreatedAt:    pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(data.UpdatedAt.UnixMilli()),
            },{{else}}			Id:  {{if .useUUID}}pointy.GetPointer(data.ID.String()){{else}}&data.ID{{end}}, {{end}}
{{.listData}}
        },
	}, nil
}

