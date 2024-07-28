package taskController

type CreateOneRequest struct {
	Task CreateOneRequestTask `json:"task"`
}

type CreateOneRequestTask struct {
	RecipientEmail string `json:"recipient_email"`
	EmailSender    string `json:"email_sender"`
	Description    string `json:"description"`
}

type CreateOneResponse struct {
	Task Task `json:"task"`
}

type Task struct {
	Status string `json:"status"`
}
