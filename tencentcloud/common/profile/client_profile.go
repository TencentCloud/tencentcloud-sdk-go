package profile

type ClientProfile struct {
	HttpProfile     *HttpProfile
	SignMethod      string
	UnsignedPayload bool
}

func NewClientProfile() *ClientProfile {
	return &ClientProfile{
		HttpProfile:     NewHttpProfile(),
		SignMethod:      "HmacSHA256",
		UnsignedPayload: false,
	}
}
