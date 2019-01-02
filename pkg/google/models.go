package mailapi

// Mail format
type Mail struct {
	UserID string `json:"userId"`
	Raw    []byte `json:"raw"`
}

// Header to be encoded
type Header struct {
	To      string
	Cc      string
	Subject string
}
