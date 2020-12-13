package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/shuheitakada/eris/pkg"
	"github.com/spf13/cobra"
)

var (
	lineNumber int
	filepath   string
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a csv file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0600)
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
				line := ""
				for _, arg := range args {
					if strings.HasPrefix(arg, "/") && strings.HasSuffix(arg, "/") {
						str := strings.Trim(arg, "/")
						lastIndex := len(str) - 1
						digit, err := strconv.Atoi(str[3:lastIndex])
						if err != nil {
							fmt.Println(err)
						}
						if regexp.MustCompile(`\\d{\d+}`).MatchString(str) {
							line += strconv.Itoa(pkg.CreateRandomNumber(digit))
						} else if regexp.MustCompile(`\\w{\d+}`).MatchString(str) {
							line += pkg.CreateRandomString(digit)
						}
					} else {
						line += arg
					}
				}
				writer.Write([]string{line})
			}
			writer.Flush()
		},
	}
	cmd.Flags().StringVarP(&filepath, "file", "f", "tmp/example.csv", "file path where the csv file is created")
	cmd.Flags().IntVarP(&lineNumber, "line", "n", 100, "line number")
	return cmd
}
