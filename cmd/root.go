package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "t",
	Short: "Hacking Toolbelt",
}

func Execute() error {
	cobra.CheckErr(rootCmd.Execute())
	return nil
}

func init() {
	rootCmd.AddCommand(urlCmd, htmlCmd)
}
