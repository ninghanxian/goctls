package docker

import (
	"fmt"
	"os"
	"text/template"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/gookit/color"
	"github.com/qmcloud/goctls/util"
	"github.com/qmcloud/goctls/util/pathx"
	"github.com/spf13/cobra"
)

const (
	dockerfileName = "Dockerfile"
	yamlEtx        = ".yaml"
)

// Docker describes a dockerfile
type Docker struct {
	Chinese     bool
	GoMainFrom  string
	GoRelPath   string
	GoFile      string
	ServiceName string
	ServiceType string
	BaseImage   string
	HasPort     bool
	Port        int
	Image       string
	HasTimezone bool
	Timezone    string
	Author      string
}

type GenContext struct {
	Home        string
	Image       string
	Remote      string
	Branch      string
	TimeZone    string
	Base        string
	Port        int
	ServiceName string
	ServiceType string
	China       bool
	Author      string
	LocalBuild  bool
}

// dockerCommand provides the entry for goctl docker
func dockerCommand(_ *cobra.Command, _ []string) (err error) {
	defer func() {
		if err == nil {
			fmt.Println(color.Green.Render("Done."))
		}
	}()

	home := varStringHome
	remote := varStringRemote
	if len(remote) > 0 {
		repo, _ := util.CloneIntoGitHome(remote, varStringBranch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}

	g := &GenContext{
		Home:        home,
		Image:       varStringImage,
		Remote:      remote,
		Branch:      varStringBranch,
		TimeZone:    varStringTZ,
		Base:        varStringBase,
		Port:        varIntPort,
		ServiceType: varServiceType,
		ServiceName: varServiceName,
		China:       varBoolChina,
		Author:      varStringAuthor,
		LocalBuild:  varBoolLocalBuild,
	}

	if err := generateDockerfile(g); err != nil {
		return err
	}

	return nil
}

func generateDockerfile(g *GenContext) error {
	color.Green.Println("Generating...")

	var projPath string
	var err error

	if len(projPath) == 0 {
		projPath = "."
	}

	if fileutil.IsExist(dockerfileName) {
		err = os.Remove(dockerfileName)
		if err != nil {
			return err
		}
	}

	out, err := os.Create(dockerfileName)
	if err != nil {
		return err
	}

	var text string

	if g.LocalBuild {
		text, err = pathx.LoadTemplate(category, dockerLocalbuildTemplateFile, dockerLocalBuildTemplate)
		if err != nil {
			return err
		}
	} else {
		text, err = pathx.LoadTemplate(category, dockerTemplateFile, dockerTemplate)
		if err != nil {
			return err
		}
	}

	t := template.Must(template.New("dockerfile").Parse(text))
	return t.Execute(out, Docker{
		Chinese:     g.China,
		GoRelPath:   projPath,
		ServiceName: g.ServiceName,
		ServiceType: g.ServiceType,
		BaseImage:   g.Base,
		HasPort:     g.Port > 0,
		Port:        g.Port,
		Image:       g.Image,
		HasTimezone: len(g.TimeZone) > 0,
		Timezone:    g.TimeZone,
		Author:      g.Author,
	})
}
