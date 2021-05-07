package morpheus

import (
	"fmt"
)

// globals

var (
	KeyPairsPath = "/api/key-pairs"
)

type KeyPair struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
	// PrivateKey string `json:"privateKey"`
}

type ListKeyPairsResult struct {
	KeyPairs *[]KeyPair  `json:"keyPairs"`
	Meta     *MetaResult `json:"meta"`
}

type GetKeyPairResult struct {
	KeyPair *KeyPair `json:"keyPair"`
}

type CreateKeyPairResult struct {
	Success bool              `json:"success"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
	KeyPair *KeyPair          `json:"keyPair"`
}

type UpdateKeyPairResult struct {
	CreateKeyPairResult
}

type DeleteKeyPairResult struct {
	DeleteResult
}

// request types

type KeyPairPayload struct {
	Name       string `json:"name"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

type CreateKeyPairPayload struct {
	KeyPair *KeyPairPayload `json:"keyPair"`
}

type UpdateKeyPairBody struct {
	CreateKeyPairPayload
}

func (client *Client) ListKeyPairs(req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path:   KeyPairsPath,
		Result: &ListKeyPairsResult{},
	})
}

func (client *Client) GetKeyPair(id int64, req *Request) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path:   fmt.Sprintf("%s/%d", KeyPairsPath, id),
		Result: &GetKeyPairResult{},
	})
}
