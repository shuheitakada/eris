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
				if strings.HasPrefix(args[0], "/") && strings.HasSuffix(args[0], "/") {
					str := strings.Trim(args[0], "/")
					if regexp.MustCompile(`\\d{\d+}`).MatchString(str) {
						lastIndex := len(str) - 1
						digit, err := strconv.Atoi(str[3:lastIndex])
						if err != nil {
							fmt.Println(err)
						}
						writer.Write([]string{strconv.Itoa(pkg.CreateRandomNumber(digit))})
					}
				}
			}
			writer.Flush()
		},
	}
	cmd.Flags().IntVarP(&lineNumber, "line", "n", 100, "line number")
	return cmd
}
