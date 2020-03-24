package profile

type HttpProfile struct {
	ReqMethod  string
	ReqTimeout int
	Endpoint   string
	// Deprecated, use Scheme instead
	Protocol string
	Scheme   string
	Debug    bool
}

func NewHttpProfile() *HttpProfile {
	return &HttpProfile{
		ReqMethod:  "POST",
		ReqTimeout: 60,
		Endpoint:   "",
		Scheme:     "HTTPS",
		Debug:      false,
	}
}
