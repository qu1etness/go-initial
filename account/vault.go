package account

import (
	"encoding/json"
	"fmt"
	"go-initial/files"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() (*Vault, error) {
	data, err := files.ReadFile("data.json")

	if err != nil {
		color.Yellow("File not found, creating a new vault.")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}, err
	}

	var valterVault = Vault{}
	err = json.Unmarshal(data, &valterVault)

	if err != nil {
		color.Red("Something went wrong with unmarshalling, creating a new vault.")
		valterVault = Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &valterVault, err
}

func (v *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err == nil {
		return nil, err
	}
	return bytes, err
}

func (v *Vault) AddAccount(account Account) {
	v.Accounts = append(v.Accounts, account)
	v.UpdatedAt = time.Now()
}

func (v *Vault) FindAccount(researchingName string) {
	for i, acc := range v.Accounts {
		if acc.Name == researchingName {
			color.Green("Account found at index %d:", i)
			color.Cyan("Name: %s", acc.Name)
			color.Cyan("URL: %s", acc.Url)
			color.Cyan("Password: %s", acc.Password)
			return
		}
	}
	color.Yellow("Account with name '%s' not found.", researchingName)
}

func (v *Vault) RemoveAccount(accountName string) {

	for index, account := range v.Accounts {
		if account.Name == accountName {
			v.Accounts = append(v.Accounts[:index], v.Accounts[index+1:]...)
			v.UpdatedAt = time.Now()
			return
		}
	}
	color.Yellow("Account with name '%s' not found.", accountName)
}

func (v *Vault) ShowAllAccounts() {
	if len(v.Accounts) == 0 {
		color.Yellow("No accounts found in the vault.")
		return
	}

	for i, acc := range v.Accounts {
		color.Blue("------%v------", i+1)
		fmt.Println(acc, "\n")
	}
}
