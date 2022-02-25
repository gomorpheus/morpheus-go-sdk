package morpheus

import (
	"fmt"
)

var (
	// KeyPairsPath is the API endpoint for key pairs
	KeyPairsPath = "/api/key-pairs"
)

// KeyPair structures for use in request and response payloads
type KeyPair struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
	// PrivateKey string `json:"privateKey"`
}

// ListKeyPairsResult structure parses the list key pairs response payload
type ListKeyPairsResult struct {
	KeyPairs *[]KeyPair  `json:"keyPairs"`
	Meta     *MetaResult `json:"meta"`
}

// GetKeyPairResult structure parses the get key pair response payload
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
