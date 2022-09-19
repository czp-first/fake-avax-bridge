package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"ShamirAlgorithm/shamir"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split an arbitrarily long secret from stdin",
	Run:   runSplit,
}

var (
	parts     int
	threshold int
	secret    string
)

func init() {
	RootCmd.AddCommand(splitCmd)

	splitCmd.Flags().IntVarP(&parts, "parts", "p", 5, "number of shares")
	splitCmd.Flags().IntVarP(&threshold, "threshold", "t", 3, "threshold of shares required to reconstruct the secret")
	splitCmd.Flags().StringVarP(&secret, "secret", "s", "", "secret to split")
}

func runSplit(cmd *cobra.Command, args []string) {
	byteSecret := []byte(secret)
	byteParts, err := shamir.Split(byteSecret, parts, threshold)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to split secret: %v\n", err)
		os.Exit(1)
	}
	for _, bytePart := range byteParts {
		fmt.Printf("%x\n", bytePart)
	}
}
