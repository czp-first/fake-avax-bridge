package cmd

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"ShamirAlgorithm/shamir"
)

var shares string

var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Recostruct a secret from the parts read from stdin",
	Run:   runCombine,
}

func init() {
	RootCmd.AddCommand(combineCmd)

	combineCmd.Flags().StringVarP(&shares, "shares", "s", "", "shares split by comma")

}

func runCombine(cmd *cobra.Command, args []string) {
	splitShares := strings.Split(shares, ",")

	var byteParts [][]byte
	for _, v := range splitShares {
		b, _ := hex.DecodeString(v)

		byteParts = append(byteParts, b)
	}

	secret, err := shamir.Combine(byteParts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to combine secret: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(secret))
}
