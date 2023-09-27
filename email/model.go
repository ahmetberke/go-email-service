package email

type Email struct {
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}
