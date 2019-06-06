package GoRequest

type InitRequestEntity struct {
	BaseUrl    string `json:"base_url"`
	UserAgent  string `json:"user_agent"`
	Device     string `json:"device"`
	SdkVersion string `json:"sdk_version"`
	Uuid       string `json:"uuid"`
	PhoneMode  string `json:"phone_mode"`
	PhoneOs    string `json:"phone_os"`
}
