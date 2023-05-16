package mail_template

var sender *gmailSender

type gmailSender struct {
	Name              string
	FromEmailAddress  string
	FromEmailPassword string
}

func GetGmailSender() *gmailSender {
	if sender != nil {
		return sender
	}

	sender = &gmailSender{
		Name:              "Shop Dummy",
		FromEmailAddress:  "dummyshophaha@gmail.com",
		FromEmailPassword: "ycjpeeqqfzvgcwnj",
	}

	return sender
}
