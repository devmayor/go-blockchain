package main

import (
	"fmt"
	"github.com/devmayor/go-blockchain-learning/wallet"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func walletCmd() *cobra.Command {
	var walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Manages blockchain accounts and keys.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	walletCmd.AddCommand(walletNewAccountCmd())

	return walletCmd
}

func walletNewAccountCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new-account",
		Short: "Creates a new account with a new set of a elliptic-curve Private + Public keys.",
		Run: func(cmd *cobra.Command, args []string) {
			password := getPassPhrase("Please enter a password to encrypt the new wallet:", true)
			dataDir := getDataDirFromCmd(cmd)

			acc, err := wallet.NewKeystoreAccount(dataDir, password)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("New account created: %s\n", acc.Hex())
			fmt.Printf("Saved in: %s\n", wallet.GetKeystoreDirPath(dataDir))
		},
	}

	addDefaultRequiredFlags(cmd)

	return cmd
}

func getPassPhrase(prompt string, confirmation bool) string {
	password := readPassword(prompt)

	if confirmation {
		confirm := readPassword("Repeat password: ")
		if password != confirm {
			fmt.Fprintln(os.Stderr, "passwords do not match")
			os.Exit(1)
		}
	}

	return password
}

func readPassword(prompt string) string {
	fmt.Print(prompt + " ")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read password: %v\n", err)
		os.Exit(1)
	}

	return string(password)
}
