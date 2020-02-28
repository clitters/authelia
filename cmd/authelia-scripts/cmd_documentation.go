package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// DocumentationCmd serve authelia with the provided configuration
func DocumentationDeployFunc(cobraCmd *cobra.Command, args []string) {
	log.Info("Deploying documentation...")
	/* cmd := utils.CommandWithStdout(OutputDir+"/authelia", "--config", args[0])
	cmd.Env = append(os.Environ(), "PUBLIC_DIR=dist/public_html")
	utils.RunCommandUntilCtrlC(cmd)*/
}

var DocumentationDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy documentation to gh-pages branch",
	Run:   DocumentationDeployFunc,
	Args:  cobra.ExactArgs(0),
}
