package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/output"
	"strings"

	// "demo/password/cloud"
	"demo/password/files"
	"fmt"

	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.ManagedVault){
	"1": createAccount,
	"2": findAccountsByLogin,
	"3": findAccountsByURL,
	"4": deleteAccount,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Can't read .env file")
	}
	enc := encrypter.NewEncrypter()
	vault, err := account.NewVault(files.NewJsonDB("data.vault"), enc)
	// vault, err := account.NewVault(cloud.NewCloudDB("https://localhost:3000"))
	if err != nil { return }

	for {
		menuChoice := takeChoice()
		menuFunc := menu[menuChoice]
		if menuFunc == nil {
			break
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.ManagedVault) {
    login     := promtData("Login")
	password  := promtData("Password")
	urlString := promtData("Url")

	acc, err := account.NewAccountWithTimestamp(login, password, urlString)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault.AddAccount(*acc)
}

func findAccountsByURL(vault *account.ManagedVault) {
	value := promtData("url")
	finder := func(acc *account.Account) bool {
		return strings.Contains(acc.Url, value)
	}
	findAccounts(vault, finder)
}

func findAccountsByLogin(vault *account.ManagedVault) {
	value := promtData("login")
	finder := func(acc *account.Account) bool {
		return strings.Contains(acc.Login, value)
	}
	findAccounts(vault, finder)
}

func findAccounts(vault *account.ManagedVault, finder func(*account.Account) bool) {
	accounts := vault.FindAccount(finder)
	if len(accounts) == 0 {
		fmt.Println("Accounts not found.")
		return
	}
	for index, acc := range accounts {
		acc.OutputIndexed(index)
	}
}

func deleteAccount(vault *account.ManagedVault) {
	url := promtData("url")

	deletedAccs := vault.DeleteAccountsByUrl(url)
	if len(deletedAccs) == 0 { return }

	fmt.Println()
	for index, acc := range deletedAccs {
		acc.OutputIndexed(index)
	}
}

func takeChoice() string {
	return promtData(
		"1 - Add account",
		"2 - Find account by login",
		"3 - Find account by url",
		"4 - Delete account by url",
		"Any other - Exit",
		"Enter choice",
	)
}

func promtData(prompt ...any) string {
	var result string
	for i, line := range prompt {
		if i + 1 == len(prompt) {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	fmt.Scanln(&result)
	return result
}
