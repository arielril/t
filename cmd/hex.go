package cmd

import (
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/arielril/t/internal"
	"github.com/spf13/cobra"
)

var format = struct {
	isBin   bool
	isAscii bool
	isNum   bool
}{}

var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Hex encode/decode",
}

var hexDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Hex decode",
	Args:  cobra.ExactArgs(1),
	Run:   decodeHex,
}

var hexEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Hex encode",
	Run:   encodeHex,
}

func init() {
	hexCmd.PersistentFlags().BoolVar(&format.isBin, "bin", false, "Binary input")
	hexCmd.PersistentFlags().BoolVar(&format.isAscii, "ascii", false, "ASCII input")
	hexCmd.PersistentFlags().BoolVar(&format.isNum, "num", false, "Number input")

	hexCmd.AddCommand(hexDecodeCmd, hexEncodeCmd)
	rootCmd.AddCommand(hexCmd)
}

func getInput(cmd *cobra.Command, args []string) string {
	if len(args) == 0 {
		inputData, err := io.ReadAll(cmd.InOrStdin())
		cobra.CheckErr(err)
		return string(inputData)
	}
	return strings.Join(args, " ")
}

func encodeHex(cmd *cobra.Command, args []string) {
	// * need to know the input format bin/string/number/hex
	input := getInput(cmd, args)

	var result string

	if format.isBin {
		x := internal.Bin(input).FromString()
		result = hex.EncodeToString(x)
	}

	if format.isAscii {
		result = hex.EncodeToString([]byte(input))
	}

	if format.isNum {
		num, err := strconv.ParseInt(input, 10, 64)
		cobra.CheckErr(err)
		result = fmt.Sprintf("%X", num)
	}

	fmt.Printf("0x%s\n", strings.ToUpper(result))
}

func decodeHex(cmd *cobra.Command, args []string) {
	// ! input as hex -> output bin/string/number/hex (same as win dbg)
	input := getInput(cmd, args)
	if strings.Contains(input, "0x") {
		input = strings.Replace(input, "0x", "", 1)
	}

	bin, err := hex.DecodeString(input)
	cobra.CheckErr(err)

	fmt.Printf("HEX: %s\nBIN: %08b\nASCII: %s\nNUM: %d\n", input, bin, string(bin), bin)
}
