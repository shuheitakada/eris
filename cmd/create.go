package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lineNumber int

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a csv file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.OpenFile("tmp/example.csv", os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			err = file.Truncate(0)
			if err != nil {
				fmt.Println(err)
			}

			writer := csv.NewWriter(file)
			for i := 0; i < lineNumber; i++ {
				writer.Write([]string{"hoge"})
			}
			writer.Flush()
		},
	}
	cmd.Flags().IntVarP(&lineNumber, "line", "n", 100, "line number")
	return cmd
}
