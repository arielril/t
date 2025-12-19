package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "url commands",
}

var urlDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode URLs",
	Run: func(_cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")

		unescaped, err := url.QueryUnescape(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s\n", unescaped)
	},
}

var urlEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encode URLs",
	Run: func(_cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")
		fmt.Printf("%s\n", url.QueryEscape(input))
	},
}

func init() {
	urlCmd.AddCommand(urlDecodeCmd, urlEncodeCmd)
}
