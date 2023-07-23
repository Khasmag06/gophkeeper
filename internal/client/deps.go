package client

type decoder interface {
	Encrypt(data []byte) (string, error)
	Decrypt(encrypted string) ([]byte, error)
}
