package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/gofrs/uuid"
)

var letters = []rune("abcdefghijklmnopqrstuvwyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_-=+<>?/")

type Account struct {
	Id        string    `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccountWithTimestamp(login string, password string, urlString string) (*Account, error) {
	if len(login) < 4 {
		return nil, errors.New("login too short")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("incorrect URL: %s", err)
	}
	id, _ := uuid.NewV4()
	acc := Account{
		Id: id.String(),
		Login: login,
		Password: password,
		Url: urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		acc.generatePassowrd(12)
	}

	return &acc, nil
}

func (acc *Account) ToBytes() ([]byte, error){
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) OutputIndexed(index int) {
	result, err := json.MarshalIndent(acc, "", "\t")
	if err != nil { return }

	fmt.Println(string(result))
}

func (acc *Account) generatePassowrd(n int) {
	chars := make([]rune, n)
	for i := range chars {
		chars[i] = letters[rand.Int32N(int32(len(letters)))]
	}
	acc.Password = string(chars)
}
