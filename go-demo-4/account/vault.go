package account

import (
	"demo/password/output"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)
type ByteReader interface {
	Read() ([]byte, error)
}
type ByteWriter interface {
	Write([]byte)
}
type DataBase interface {
	ByteReader
	ByteWriter
}
type Encrypted interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ManagedVault struct {
	Vault
	db     DataBase
	enc    Encrypted
}

func NewVault(db DataBase, enc Encrypted) (*ManagedVault, error) {
	file, err := db.Read()
	vault := &ManagedVault{
		Vault: Vault{
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		},
		db: db,
		enc: enc,
	}
	if err != nil {
		output.PrintError(fmt.Sprintf("Can't open existing vault: %v", err))
		return vault, nil
	}
	file = enc.Decrypt(file)
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError(fmt.Sprintf("Can't create vault: %v", err))

		return nil, err
	}
	return vault, nil
}

func (accs *Vault) ToBytes() ([]byte, error) {
	data, err := json.Marshal(accs)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (accs *ManagedVault) FindAccount(finder func(*Account) bool) ([]Account) {
	var result []Account
	for _, acc := range accs.Accounts {
		if finder(&acc) {
			result = append(result, acc)
		}
	}
	return result
}

func (accs *ManagedVault) DeleteAccountsByUrl(url string) ([]Account) {
	var result []Account
	var saved []Account
	for _, acc := range accs.Accounts {
		if strings.Contains(acc.Url, url) {
			result = append(result, acc)
		} else {
			saved = append(saved, acc)
		}
	}
	accs.Accounts = saved
	accs.Save()
	return result
}

func (accs *ManagedVault) AddAccount(acc Account) {
	accs.Accounts = append(accs.Accounts, acc)
	accs.Save()
}

func (accs *ManagedVault) Save() {
	accs.UpdatedAt = time.Now()

	data, err := accs.ToBytes()
	if err != nil { 
		output.PrintError(fmt.Sprintf("Can't add account: %v", err))
	}
	data = accs.enc.Encrypt(data)
	accs.db.Write(data)
	output.Success()
}
