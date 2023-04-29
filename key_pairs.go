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

func (client *Client) ListKeyPairs() (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path:   KeyPairsPath,
		Result: &ListKeyPairsResult{},
	})
}

func (client *Client) GetKeyPair(id int64) (*Response, error) {
	return client.Execute(&Request{
		Method: "GET",
		Path:   fmt.Sprintf("%s/%d", KeyPairsPath, id),
		Result: &GetKeyPairResult{},
	})
}
func (client *Client) CreateKeyPair(name string, publicKey string) (*Response, error) {
	req := &Request{
		Method: "POST",
		Path:   "/api/key-pairs",
		Body: map[string]interface{}{
			"keyPair": map[string]string{
				"name":      "test123",
				"publicKey": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDJzKbAk5Yu1YDNTcr4IxKOcvO3nZtCbRECons3Hl7SWCgX1+aj1BANJY0DTmpjbQ+rv+TkbWrQePwRy+HR4s6So4GJImlyOLmFMH5qJOJRMCoi3vUgyFa0cluT4G9MBK66ym1kYB3ZXh3RAMwmOY3jKmjRAA5282K3F+UT+bedihQZuuW0IIOrnLD1kFg39xB8XJsu0ysym1L1tEfh19i15bciQdESsRf0ClAn8ELgn6LZ/VYS5cOr03wa0VZojwF+GrxHEvYNP0KKvUb9Sgt7grlfieWf0f8LHdS6vbNNhumvzGJWHU9Ak6jNkvXO3aHaGhOUtoZYtXhYS96C8FlatZGHKnhMfkkRWcujW7lB4xujPs7R0yzOCwHZ65BRCWk1qn2yZti8OYhl4uZ4arUOMhXvOndLKOGBf1tShLrNxATs4Fe5L0G0B6iL0xJ8Xc4m1ppfot91hRwyFzn4QuVDnlr4zKyQI7JhOhwprwJDPreKi+iUPtRVxweXg98hOK0= stefan@console",
			},
		},
	}
	return client.Execute(req)
}
func (client *Client) DeleteKeyPair(id int64) (*Response, error) {
	return client.Execute(&Request{
		Path:   fmt.Sprintf("%s/%d", KeyPairsPath, id),
		Method: "DELETE",
		Result: &DeleteKeyPairResult{},
	})
}
