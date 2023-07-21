package client

import (
	"bufio"
	"fmt"
	"github.com/Khasmag06/gophkeeper/internal/models"
	"os"
)

func (c *client) Run() error {
	signupPath := "/api/user/signup"
	loginPath := "/api/user/login"
	loginCredsPath := "/api/record/login-creds"
	addLoginCredsPath := "/api/record/login-creds/add"
	bankCardPath := "/api/record/bank-card"
	addBankCardPath := "/api/record/bank-card/add"
	binaryDataPath := "/api/record/binary"
	addBinaryDataPath := "/api/record/binary/add"
	textDataPath := "/api/record/text"
	addTextDataPath := "/api/record/text/add"

	paths := signupPath + "\n" +
		loginPath + "\n" +
		loginCredsPath + "\n" +
		addLoginCredsPath + "\n" +
		bankCardPath + "\n" +
		addBankCardPath + "\n" +
		binaryDataPath + "\n" +
		addBinaryDataPath + "\n" +
		textDataPath + "\n" +
		addTextDataPath
	fmt.Println(paths)
	fmt.Println("Введите один из указанных маршрутов")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	event := scanner.Text()

	switch event {
	case signupPath:
		if err := c.RegisterUser(); err != nil {
			return err
		}
	case loginPath:
		if err := c.AuthenticateUser(); err != nil {
			return err
		}
	case addLoginCredsPath:
		if err := c.AddRecord(addLoginCredsPath); err != nil {
			return err
		}
	case loginCredsPath:
		var record []*models.LoginCredentials
		if err := c.GetRecords(loginCredsPath, record); err != nil {
			return err
		}
	case addTextDataPath:
		if err := c.AddRecord(addTextDataPath); err != nil {
			return err
		}
	case textDataPath:
		var record []*models.TextData
		if err := c.GetRecords(textDataPath, record); err != nil {
			return err
		}

	case addBinaryDataPath:
		if err := c.AddRecord(addBinaryDataPath); err != nil {
			return err
		}
	case binaryDataPath:
		var record []*models.BinaryData
		if err := c.GetRecords(binaryDataPath, record); err != nil {
			return err
		}

	case addBankCardPath:
		if err := c.AddRecord(addBankCardPath); err != nil {
			return err
		}
	case bankCardPath:
		var record []*models.BankCard
		if err := c.GetRecords(bankCardPath, record); err != nil {
			return err
		}

	}
	return nil
}
