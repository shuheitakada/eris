package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
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
		for i := 0; i < 100; i++ {
			writer.Write([]string{"hoge"})
		}
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
