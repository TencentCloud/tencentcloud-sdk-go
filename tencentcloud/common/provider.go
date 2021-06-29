package common

// Provider interface
type Provider interface {
	GetCredential() (CredentialIface, error)
}
