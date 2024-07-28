package models

type Task struct {
	BaseModel
	RecipientEmail string `json:"recipient_email"`
	EmailSender    string `json:"email_sender"`
	Description    string `json:"description"`
	IsViewed       bool
}

func (t Task) TableName() string {
	return "task"
}
