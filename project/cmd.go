package project

import (
	"github.com/qmcloud/goctls/internal/cobrax"
	"github.com/qmcloud/goctls/project/upgrade"
)

var (
	Cmd        = cobrax.NewCommand("project")
	upgradeCmd = cobrax.NewCommand("upgrade", cobrax.WithRunE(upgrade.UpgradeProject))
)

func init() {
	upgradeCmdFlag := upgradeCmd.Flags()

	upgradeCmdFlag.BoolVarP(&upgrade.VarBoolUpgradeMakefile, "makefile", "m")

	Cmd.AddCommand(upgradeCmd)
}
