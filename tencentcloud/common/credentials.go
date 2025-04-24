package common

var creErr = "ClientError.CredentialError"

type CredentialIface interface {
	GetSecretId() string

	GetToken() string

	GetSecretKey() string

	// GetCredential retrieves all credential components (SecretId, SecretKey, Token) as a tuple.
	// This method is designed to fetch all three credential values (SecretId, SecretKey, Token) at once.
	// This is beneficial for atomic operations, preventing potential issues that could arise from
	// separate calls where one credential element might be updated while others are not, leading to inconsistency.
	GetCredential() (string, string, string)

	// needRefresh() bool
	// refresh()
}

// Credential is a basic implementation of the CredentialIface.
// It stores the SecretId, SecretKey, and Token directly.
// This struct is suitable for simple cases where credentials are known beforehand.
type Credential struct {
	SecretId  string // The SecretId for authentication.
	SecretKey string // The SecretKey for signing requests.
	Token     string // The security token (optional).
}

func (c *Credential) needRefresh() bool {
	return false
}

func (c *Credential) refresh() {
}

// NewCredential creates a new Credential instance with the given SecretId and SecretKey.
// This constructor is used when you don't have a security token.
func NewCredential(secretId, secretKey string) *Credential {
	return &Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}

// NewTokenCredential creates a new Credential instance with the given SecretId, SecretKey, and Token.
// This constructor is used when you have a security token.
func NewTokenCredential(secretId, secretKey, token string) *Credential {
	return &Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
		Token:     token,
	}
}

func (c *Credential) GetSecretKey() string {
	return c.SecretKey
}

func (c *Credential) GetSecretId() string {
	return c.SecretId
}

func (c *Credential) GetToken() string {
	return c.Token
}

func (c *Credential) GetCredential() (string, string, string) {
	return c.SecretId, c.SecretKey, c.Token
}
