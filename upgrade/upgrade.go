package upgrade

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/qmcloud/goctls/rpc/execx"
	"github.com/spf13/cobra"
)

// upgrade gets the latest goctl by
// go install github.com/qmcloud/goctls@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `go install github.com/qmcloud/goctls@latest`
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	color.Green.Println("Upgrade successfully")
	return nil
}
