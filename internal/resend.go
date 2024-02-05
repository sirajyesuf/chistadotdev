package internal

import (
	"github.com/resend/resend-go/v2"
)

type ReSendEmailRequest struct {
	Apikey string
	From    string   `json:"from"`
	To      []string `json:"to"`
	Html    string   `json:"html"`
	Subject string   `json:"subject"`
	Cc      []string `json:"cc"`
	Bcc     []string `json:"bcc"`
	ReplyTo string   `json:"replyto"`
}



func ResendHandler(req ReSendEmailRequest) (*resend.SendEmailResponse ,error) {

    var apiKey = req.Apikey

    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    req.From,
        To:      req.To,
        Html:    req.Html,
        Subject: req.Subject,
        Cc:      req.Cc,
        Bcc:     req.Bcc,
        ReplyTo: req.ReplyTo,
    }


    sent, err := client.Emails.Send(params)
    // if err != nil {
    //     fmt.Println(err.Error())
    //     return
    // }
    // fmt.Println(sent.Id)

    return sent,err
}
