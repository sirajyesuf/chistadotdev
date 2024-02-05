package internal

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/chista.dev/pkg"
	"github.com/gin-gonic/gin"
)

type ChistaRequest struct {
	From   struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	To []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"to"`
	HTML    string   `json:"HTML"`
	Subject string   `json:"subject"`
	Cc      []string `json:"cc"`
	Bcc     []string `json:"bcc"`
	ReplyTo string   `json:"replyto"`
	Text    string   `json:"text"`
}

type ChistaResponse struct { 
	sent struct {
		ServiceName string `json:"service_name"`
		StatusCode  int   `json:"status_code"`
	}
	err []ErrorResponse
}

type ErrorResponse struct {
	ServiceName string `json:"service_name"`
	Msg         string  `json:"msg"`
}



func Chista(c *gin.Context)  {
	
	var request ChistaRequest
	var response ChistaResponse
	var DB = pkg.GetRepo()

	if c.ShouldBind(&request) == nil {

		userid,_ := c.Get("userId")
		// apikey,_ := c.Get("apiKey")

		user ,_ := DB.GetUserServices(userid.(uint))

		for _,service := range user.Services {

			userService,err:= DB.GetUserService(service.ID)

			if err != nil {
				continue
			}
			
			if service.Name == "resend" {

				var emails []string
				for _, entry := range request.To {
						emails = append(emails, entry.Email)
				}

				var payload = ReSendEmailRequest{

					// Apikey: "re_KQvtzmbz_BN4sJgGdmEPhWR3TQFK9JtPm",
					Apikey: userService.CreadA,
					From:fmt.Sprintf("%s <%s>",request.From.Name,request.From.Email),
					To: emails,
					Html: request.HTML,
					Subject: request.Subject,
					Cc: request.Cc,
					Bcc: request.Bcc,
					ReplyTo: request.ReplyTo,
				}


				sent,err := ResendHandler(payload)

				fmt.Println("resend")
				fmt.Println(sent)
				fmt.Println(err)

				//succesfully sent
				if err == nil {
					response.sent.ServiceName = service.Name
					num,_ := strconv.Atoi(sent.Id)
					response.sent.StatusCode  = num
					break
				}

				response.err = append(response.err,ErrorResponse{
					ServiceName: service.Name,
					Msg: err.Error(),
				})

				fmt.Print(response)
			}





			if(service.Name == "mailersend"){

				var payload = MailerSendPayload{
					
					// MAILERSEND_API_KEY: "mlsn.c15f030cbd99c7dd968a0f16a0fb1d314af4f9ed702e5d02ef4bc3722c7f6d20",
					MAILERSEND_API_KEY: userService.CreadA,
					From: request.From,
					To: request.To,
					Subject: request.Subject,
					Text: request.Text,
					Html: request.HTML,
					
				}
				
				sent,err := MailerSendHandler(payload)
				fmt.Println("mailersend")
				fmt.Println(sent)
				fmt.Println(err)
				
				//succesfully sent
				if err == nil {
					response.sent.ServiceName = service.Name
					response.sent.StatusCode  = sent.StatusCode
					break
				}

				response.err = append(response.err,ErrorResponse{
					ServiceName: service.Name,
					Msg: err.Error(),
				})

				fmt.Println(response)

				
			}
		}

		// fmt.Println("loop end")
		// fmt.Println(response)

		if response.sent.ServiceName == ""{

			c.JSON(http.StatusOK, gin.H{
				"err":response.err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"sent": response.sent,
		})



	}
}