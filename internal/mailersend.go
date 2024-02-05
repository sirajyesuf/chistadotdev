package internal

import (
	"context"
	// "fmt"
	// "os"
	// "time"

	// "fmt"

	"github.com/mailersend/mailersend-go"
)

type MailerSendPayload struct {

	MAILERSEND_API_KEY string
	From   struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	To  []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"to"`
	Subject string   `json:"subject"`
	Text string
	Html string
	ReplyTo string   `json:"replyto"`
	// Tags []string
    // Cc []struct {
	// 	Name string
	// 	Email string
	// }
	// Bcc []struct{
	// 	Name string
	// 	Email string
	// }

	Cc      []string `json:"cc"`
	Bcc     []string `json:"bcc"`


}

func MailerSendHandler(payload MailerSendPayload) (*mailersend.Response,error) {
	
	// Create an instance of the mailersend client
	ms := mailersend.NewMailersend(payload.MAILERSEND_API_KEY)

	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	subject := payload.Subject
	text := payload.Text
	html := payload.Html

	from := mailersend.From{
		Name:  payload.From.Name,
		Email: payload.From.Email,
	}

	// recipients := []mailersend.Recipient{
	// 	{
	// 		Name:  "Your Client",
	// 		Email: "your@client.com",
	// 	},
	// }

	var recipients []mailersend.Recipient
	for _, to := range payload.To {
		recipients = append(recipients,mailersend.Recipient{
			Email: to.Email,
			Name:  to.Name,
		})
	}
	
	
	// Send in 5 minute
	// sendAt := time.Now().Add(time.Minute * 5).Unix()

	// tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	// message.SetTags(tags)
	// message.SetSendAt(sendAt)
	message.SetInReplyTo("client-id")

	res, err := ms.Email.Send(ctx, message)

	// fmt.Printf(res.Header.Get("X-Message-Id"))

	return res,err

}
