package helpers

// EncryptRequest structure request comes from client
type EncryptRequest struct {
	Text string `json:"text"`
	Key string `json:"key"`
}

// EncryptResponse structure response going to the client
type EncryptResponse struct {
	Message string `json:"message"`
	Err string `json:"err"`
}

// DecryptRequest structure request comes from client
type DecryptRequest struct {
	Message string `json:"message"`
	Key string `json:"key"`
}

// DecryptResponse structure response going to the client
type DecryptResponse struct {
	Text string `json:"message"`
	Err string `json:"err"`
}