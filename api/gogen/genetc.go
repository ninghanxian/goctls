package gogen

import (
	_ "embed"
	"fmt"

	"github.com/qmcloud/goctls/api/spec"
	"github.com/qmcloud/goctls/config"
	"github.com/qmcloud/goctls/util/format"
)

const (
	etcDir = "etc"
)

//go:embed etc.tpl
var etcTemplate string

func genEtc(dir string, cfg *config.Config, api *spec.ApiSpec, g *GenContext) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, api.Service.Name)
	if err != nil {
		return err
	}

	service := api.Service
	host := "0.0.0.0"
	port := g.Port

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          etcDir,
		filename:        fmt.Sprintf("%s.yaml", filename),
		templateName:    "etcTemplate",
		category:        category,
		templateFile:    etcTemplateFile,
		builtinTemplate: etcTemplate,
		data: map[string]any{
			"serviceName": service.Name,
			"host":        host,
			"port":        port,
			"useCasbin":   g.UseCasbin,
			"useI18n":     g.UseI18n,
			"useEnt":      g.UseEnt,
			"useCoreRpc":  g.UseCoreRpc,
		},
	})
}
