package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Khasmag06/gophkeeper/config"
	"github.com/Khasmag06/gophkeeper/internal/models"
	"io"
	"net/http"
	"os"
)

type client struct {
	*http.Client
	serverAddress string
	decoder       decoder
	token         string
}

func New(cfg config.ServerConfig, decoder decoder) *client {
	serverAddress := fmt.Sprintf("http://%s%s", cfg.Host, cfg.Port)
	return &client{
		Client:        &http.Client{},
		serverAddress: serverAddress,
		decoder:       decoder,
	}
}

type SuccessTokenResponse struct {
	Status string                 `json:"status"`
	Data   *models.TokensResponse `json:"data"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func (c *client) RegisterUser() error {
	fmt.Println("Введите логин и пароль в формате json для регистрации")
	registrationURL := fmt.Sprintf("%s/api/user/signup", c.serverAddress)
	resp, err := c.GetUserResponse(registrationURL)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	defer resp.Body.Close()
	return nil
}

func (c *client) AuthenticateUser() error {
	fmt.Println("Введите логин и пароль в формате json для аутентификации")
	authenticationURL := fmt.Sprintf("%s/api/user/login", c.serverAddress)
	resp, err := c.GetUserResponse(authenticationURL)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)

	var successTokenResp SuccessTokenResponse
	err = json.Unmarshal(body, &successTokenResp)
	if err != nil {
		return err
	}

	c.token = successTokenResp.Data.AccessToken
	fmt.Println(string(body))
	defer resp.Body.Close()
	return nil
}

func (c *client) AddRecord(path string) error {
	fmt.Println("Введите записи для добавления в формате json")
	url := fmt.Sprintf("%s%s", c.serverAddress, path)
	recordJSON := GetDataFromConsole()
	encryptedRecord, err := c.decoder.Encrypt(recordJSON)
	if err != nil {
		return err
	}
	reqRecord := models.Record{EncryptedData: encryptedRecord}
	reqRecordBytes, err := json.Marshal(&reqRecord)
	if err != nil {
		return err
	}
	requestBody := bytes.NewBuffer(reqRecordBytes)

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
	return nil
}

func (c *client) GetRecords(path string, records any) error {
	url := fmt.Sprintf("%s%s", c.serverAddress, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	var successResp SuccessResponse
	err = json.Unmarshal(body, &successResp)
	if err != nil {
		return err
	}

	if successResp.Status == "error" {
		fmt.Println(string(body))
		return nil
	}

	decryptedRecords, err := c.decoder.Decrypt(successResp.Data)
	if err != nil {
		return fmt.Errorf("error to decrypt body: %w", err)
	}

	err = json.Unmarshal(decryptedRecords, &records)
	if err != nil {
		return err
	}
	fmt.Println(records)

	defer resp.Body.Close()
	return nil
}

func (c *client) GetUserResponse(url string) (*http.Response, error) {
	userJSON := GetDataFromConsole()
	requestBody := bytes.NewBuffer(userJSON)

	resp, err := c.Client.Post(url, "application/json", requestBody)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetDataFromConsole() []byte {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Bytes()

	return data
}
