{{.groupName}}  rpc create{{.modelName}} ({{.modelName}}Info) returns (Base{{if .useUUID}}UU{{end}}ID{{.IdType}}Resp);
{{.groupName}}  rpc update{{.modelName}} ({{.modelName}}Info) returns (BaseResp);
{{.groupName}}  rpc get{{.modelName}}List ({{.modelName}}ListReq) returns ({{.modelName}}ListResp);
{{.groupName}}  rpc get{{.modelName}}ById ({{if .useUUID}}UU{{end}}ID{{.IdType}}Req) returns ({{.modelName}}Info);
{{.groupName}}  rpc delete{{.modelName}} ({{if .useUUID}}UU{{end}}IDs{{.IdType}}Req) returns (BaseResp);
