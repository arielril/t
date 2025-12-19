package cmd

import (
	"fmt"
	"html"
	"strings"

	"github.com/spf13/cobra"
)

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Escape/Unescape HTML strings",
}

var escapeHtmlCmd = &cobra.Command{
	Use:   "escape",
	Short: "Escape HTML strings",
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")
		fmt.Printf("%s\n", html.EscapeString(input))
	},
}

var unescapeHtmlCmd = &cobra.Command{
	Use:   "unescape",
	Short: "Unescape HTML strings",
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")
		fmt.Printf("%s\n", html.UnescapeString(input))
	},
}

func init() {
	htmlCmd.AddCommand(escapeHtmlCmd, unescapeHtmlCmd)
}
