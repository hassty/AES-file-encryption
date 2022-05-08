package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	aes "github.com/hassty/aes/pkg"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt file using aes",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := os.Stdin
		if len(args) == 1 {
			var err error
			file, err = os.Open(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}

		key, err := os.ReadFile(Keyfile)
		if err != nil {
			if pathErr, ok := err.(*os.PathError); ok {
				log.Fatalf("invalid path to keyfile: %v\n", pathErr.Path)
			} else {
				log.Fatal(err)
			}
		}
		key = bytes.TrimSpace(key)

		msg, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		msg = bytes.TrimSpace(msg)

		encrypted, err := aes.Encrypt(key, msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encrypted)
	},
}
