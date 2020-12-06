package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a csv file",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("create is called")
		},
	}
	return cmd
}
