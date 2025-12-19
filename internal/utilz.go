package internal

import (
	"io"

	"github.com/spf13/cobra"
)

func GetCmdPositionalArgs(cmd *cobra.Command, args []string) []string {
	if len(args) == 0 {
		inputData, err := io.ReadAll(cmd.InOrStdin())
		cobra.CheckErr(err)
		args = append(args, string(inputData))
	}
	return args
}
