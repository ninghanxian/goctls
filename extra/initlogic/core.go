// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package initlogic

import (
	_ "embed"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"

	"github.com/qmcloud/goctls/util/console"
)

//go:embed core.tpl
var coreTpl string

type CoreGenContext struct {
	Target      string
	ModelName   string
	Output      string
	Style       string
	ServiceName string
	RoutePrefix string
}

func GenCore(g *CoreGenContext) error {
	var coreString strings.Builder
	coreTemplate, err := template.New("init_core").Parse(coreTpl)
	if err != nil {
		return errors.Wrap(err, "failed to create core init template")
	}

	err = coreTemplate.Execute(&coreString, map[string]any{
		"modelName":      g.ModelName,
		"modelNameSnake": strcase.ToSnake(g.ModelName),
		"modelNameLower": strings.ToLower(g.ModelName),
		"modelNameUpper": strings.ToUpper(g.ModelName),
		"serviceName":    g.ServiceName,
		"routePrefix":    g.RoutePrefix,
	})
	if err != nil {
		return err
	}

	if g.Output != "" {
		absPath, err := filepath.Abs(g.Output)
		if err != nil {
			return errors.Wrap(err, "failed to find the output file")
		}

		if g.Output == "." {
			absPath = filepath.Join(absPath, "internal/logic/base/init_database_api_data.go")
			if !fileutil.IsExist(absPath) {
				return fmt.Errorf("failed to find the target file: %s", absPath)
			}
		}

		apiData, err := os.ReadFile(absPath)

		originalString := string(apiData)

		insertIndex := strings.Index(originalString, "err := l.svcCtx.DB.API.CreateBulk")

		if insertIndex == -1 {
			return errors.New("cannot find the insert place in the output file")
		} else {
			newString := originalString[:insertIndex] + coreString.String() + originalString[insertIndex:]

			err := os.WriteFile(absPath, []byte(newString), 0o666)
			if err != nil {
				return errors.Wrap(err, "failed to write data to output file")
			}
		}
	} else {
		console.Info(coreString.String())
	}

	return err
}
