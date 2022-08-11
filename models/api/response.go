package api

type URLCreated struct {
	ShortURL string `json:"short_url,omitempty"`
	Error    string `json:"error,omitempty"`
}
