package cmd

import (
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/bpcli/config"
	"github.com/spf13/cobra"
)

var setAccountCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the BindPlane account to be used",
	Run: func(cmd *cobra.Command, args []string) {
		set()
	},
}

func init() {
	accountCmd.AddCommand(setAccountCmd)
	setAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
	setAccountCmd.MarkFlagRequired("name")
}

func set() {
	if err := config.SetCurrent(accountName); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(accountName + " has successfully been set as current")
}
