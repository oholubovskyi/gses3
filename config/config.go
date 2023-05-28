package config

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Routes struct {
		Rate       string `json:"rate"`
		Subscribe  string `json:"subscribe"`
		SendEmails string `json:"sendEmails"`
	} `json:"routes"`
	Storage struct {
		Subscriptions string `json:"subscriptions"`
	} `json:"storage"`
	Smtp struct {
		SmtpServer string `json:"smtpServer"`
		SmtpPort   int    `json:"smtpPort"`
		Sender     string `json:"sender"`
		Password   string `json:"password"`
	} `json:"smtp"`
}
