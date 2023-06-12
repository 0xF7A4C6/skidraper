package reporter

type Report struct {
	MessageLink string
	Description string
	ReportType  string
	Subject     string

	FormID      string
	Email       string
}

type Session struct {
	CurrentSession struct {
		CsrfToken string `json:"csrf_token"`
	} `json:"current_session"`
}
