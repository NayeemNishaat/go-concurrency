package lib

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"strings"
	"sync"
)

type Mail struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string
	From       string
	FromName   string
	Wait       *sync.WaitGroup
	MailerChan chan Message
	ErrorChan  chan error
	DoneChan   chan bool
}

type Message struct {
	From        string
	FromName    string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
	Data        any
	DataMap     map[string]any
	Template    string
}

type Sender struct {
	auth smtp.Auth
}

func (m *Mail) SendMail(msg Message, errorChan chan error) {
	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.From
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}

	msg.DataMap = data

	// formatterMessage, err := m.buildPlainMessage(msg)
	// if err != nil {
	// 	errorChan <- err
	// }

	plainMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		errorChan <- err
	}
	msg.Body = plainMessage

	m.send(&msg)
}

// func (m *Mail) buildPlainMessage(msg Message) (string, error) {
// 	return "", nil
// }

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	templateToRender := fmt.Sprintf("./template/%s.html.gohtml", msg.Template)
	t, err := template.New("email-plain").ParseFiles(templateToRender)

	if err != nil {
		return "", nil
	}

	var tpl bytes.Buffer

	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", nil
	}

	plainMessage := tpl.String()

	return plainMessage, nil
}

func (m *Mail) getSender() *Sender {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)
	return &Sender{auth}
}

func (m *Mail) send(msg *Message) error {
	return smtp.SendMail(fmt.Sprintf("%s:%d", m.Host, m.Port), m.getSender().auth, msg.From, msg.To, msg.ToBytes())
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	} else {
		buf.WriteString("Content-Type: text/html; charset=utf-8\n")
		// buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	}

	buf.WriteString(m.Body)

	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}

// Test:
/* m := lib.Mail{
	Host:       "app.debugmail.io",
	Port:       25,
	Username:   "80a6f626-3168-41c9-be7d-36c93eab65ef",
	Password:   "510378c4-c0c2-4df7-bc98-ce974bfda293",
	Encryption: "none",
	From:       "80a6f626-3168-41c9-be7d-36c93eab65ef",
	ErrorChan:  make(chan error),
}

msg := lib.Message{
	To:      []string{"recipient@example.com"},
	Subject: "Test",
	Data:    "Hello",
}

m.SendMail(msg, make(chan error))
*/
