package mail

type MailAttachment struct {
	FileDir  string
	FileName string
}

type MailDetail struct {
	To          []string
	Cc          []string
	Subject     string
	Html        string
	Template    string
	Attachments []MailAttachment
	Data        interface{}
}
