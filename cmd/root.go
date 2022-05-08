package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	Keyfile  string
	Textfile string
	Iv       string
	rootCmd  = &cobra.Command{
		Use:   "aes",
		Short: "encrypt or decrypt files using aes",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	log.SetFlags(0)

	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)

	encryptCmd.Flags().StringVarP(&Keyfile, "keyfile", "k", "aes.key", "path to keyfile")
	encryptCmd.Flags().StringVar(&Iv, "iv", "", "initialization vector")
	decryptCmd.Flags().StringVarP(&Keyfile, "keyfile", "k", "aes.key", "path to keyfile")
	decryptCmd.Flags().StringVar(&Iv, "iv", "", "initialization vector")
}
