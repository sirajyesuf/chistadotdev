package internal

import (
	mailjet "github.com/mailjet/mailjet-apiv3-go/v4"
)

//mailjet
// if value.CreadA !="" && value.CreadB !=""{

// 	var payload = MailJetEmailRequest{
// 		MJ_APIKEY_PUBLIC: "6841ad7d0d220b3dd0091e0193a25d3d",
// 		MJ_APIKEY_PRIVATE: "54d9b77950efff425028359d0a9d6e19",
// 		From: request.From,
// 		To: request.To,
// 		Text: request.Text,
// 		Html: request.HTML,
// 		ReplyTo: request.ReplyTo,
// 		Cc: request.Cc,
// 		Bcc: request.Bcc,
// 	}

// 	sent,err := MailJetHandler(payload)

// 	c.JSON(200,gin.H{
// 		"sentID":sent,
// 		"err":err,
// 		"service":"mailjet",

// 	})
// }
			
type MailJetEmailRequest struct {

	MJ_APIKEY_PUBLIC string 
	MJ_APIKEY_PRIVATE string
	From   struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	To []struct {
		Email  string `json:"email"`
		Name string `json:"name"`
	} `json:"to"`

	Text    string
	Html    string  
	Subject string   
	Cc      []string 
	Bcc     []string 
	ReplyTo string
}






func MailJetHandler(payload MailJetEmailRequest) (*mailjet.ResultsV31,error){

	var recipients []mailjet.RecipientV31
	for _, to := range payload.To {
		recipients = append(recipients, mailjet.RecipientV31{
			Email: to.Email,
			Name:  to.Name,
		})
	}
	
	mailjetClient := mailjet.NewMailjetClient(payload.MJ_APIKEY_PUBLIC,payload.MJ_APIKEY_PRIVATE)
	messagesInfo := []mailjet.InfoMessagesV31 {
     {
        From: &mailjet.RecipientV31{
			Email: payload.From.Email,
			Name: payload.From.Name,
		},
        To:(*mailjet.RecipientsV31)(&recipients),
        Subject: payload.Subject,
        TextPart: payload.Text,
        HTMLPart: payload.Html,
	},
    }
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := mailjetClient.SendMailV31(&messages)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Data: %+v\n", res)


	return res,err
}