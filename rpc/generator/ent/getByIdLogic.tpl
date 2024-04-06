package {{.packageName}}

import (
	"context"

	"{{.projectPath}}{{.importPrefix}}/internal/svc"
	"{{.projectPath}}{{.importPrefix}}/internal/utils/dberrorhandler"
	"{{.projectPath}}{{.importPrefix}}/types/{{.projectName}}"

{{if .useUUID}}    "github.com/qmcloud/admin-common/utils/uuidx"
{{end}}{{if .HasCreated}}	"github.com/qmcloud/admin-common/utils/pointy"{{end}}
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *Get{{.modelName}}ByIdLogic) Get{{.modelName}}ById(in *{{.projectName}}.{{if .useUUID}}UU{{end}}ID{{.IdType}}Req) (*{{.projectName}}.{{.modelName}}Info, error) {
	result, err := l.svcCtx.DB.{{.modelName}}.Get(l.ctx, {{if .useUUID}}uuidx.ParseUUIDString({{end}}in.Id{{if .useUUID}}){{end}})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &{{.projectName}}.{{.modelName}}Info{
		Id:          {{if .useUUID}}pointy.GetPointer(result.ID.String()){{else}}&result.ID{{end}},{{if .HasCreated}}
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),{{end}}
{{.listData}}
	}, nil
}

