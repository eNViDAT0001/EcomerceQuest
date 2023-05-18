package entities

type Smtp struct {
	ID          uint   `gorm:"column:id;primary_key" json:"id"`
	Subject     string `gorm:"column:subject" json:"subject"`
	Content     string `gorm:"column:content" json:"content"`
	ToEmails    string `gorm:"column:to_emails" json:"to"`
	Cc          string `gorm:"column:cc" json:"cc"`
	Bcc         string `gorm:"column:bcc" json:"bcc"`
	AttachFiles string `gorm:"column:attach_files" json:"attach_files"`
	Status      string `gorm:"column:status" json:"status"`
}

func (Smtp) TableName() string {
	return "Smtp"
}

const (
	EmailWaiting = "WAITING"
	EmailSending = "SENDING"
	EmailSuccess = "SUCCESS"
	EMailFailed  = "FAILED"
)
