package docker

import (
	_ "embed"

	"github.com/qmcloud/goctls/util/pathx"
)

const (
	category                     = "docker"
	dockerTemplateFile           = "docker.tpl"
	dockerLocalbuildTemplateFile = "docker_local_build.tpl"
)

//go:embed docker.tpl
var dockerTemplate string

//go:embed docker_local_build.tpl
var dockerLocalBuildTemplate string

// Clean deletes all templates files
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates creates docker template files
func GenTemplates() error {
	return initTemplate()
}

// Category returns the const string of docker category
func Category() string {
	return category
}

// RevertTemplate recovers the deleted template files
func RevertTemplate(name string) error {
	return pathx.CreateTemplate(category, name, dockerTemplate)
}

// Update deletes and creates new template files
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return initTemplate()
}

func initTemplate() error {
	return pathx.InitTemplates(category, map[string]string{
		dockerTemplateFile:           dockerTemplate,
		dockerLocalbuildTemplateFile: dockerLocalBuildTemplate,
	})
}
