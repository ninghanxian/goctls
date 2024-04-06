package {{.packageName}}

import (
	"context"

	"{{.projectPath}}{{.importPrefix}}/internal/svc"
	"{{.projectPath}}{{.importPrefix}}/internal/utils/dberrorhandler"
	"{{.projectPath}}{{.importPrefix}}/types/{{.projectName}}"

{{if .useI18n}}    "github.com/qmcloud/admin-common/i18n"
{{else}}    "github.com/qmcloud/admin-common/msg/errormsg"
{{end}}{{if .hasUUID}}    "github.com/qmcloud/admin-common/utils/uuidx"
{{end}}
{{if .hasPointy}}	"github.com/qmcloud/admin-common/utils/pointy"
{{end}}	"github.com/zeromicro/go-zero/core/logx"
)

type Create{{.modelName}}Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreate{{.modelName}}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Create{{.modelName}}Logic {
	return &Create{{.modelName}}Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Create{{.modelName}}Logic) Create{{.modelName}}(in *{{.projectName}}.{{.modelName}}Info) (*{{.projectName}}.Base{{if .useUUID}}UU{{end}}ID{{.IdType}}Resp, error) {
    {{if not .hasSingle}}result, err{{else}}query{{end}} := l.svcCtx.DB.{{.modelName}}.Create(){{if .noNormalField}}.{{end}}
{{.setLogic}}

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &{{.projectName}}.Base{{if .useUUID}}UU{{end}}ID{{.IdType}}Resp{Id: result.ID{{if .useUUID}}.String(){{end}}, Msg: {{if .useI18n}}i18n.CreateSuccess{{else}}errormsg.CreateSuccess{{end}} }, nil
}
