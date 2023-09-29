package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go/twiml"
)

func main() {
	router := gin.Default()

	router.POST("/sms", func(context *gin.Context) {
		message := &twiml.MessagingMessage{
			Body: "The Robots are coming! Head for the hills!",
		}
		fmt.Println(message)

		twimlResult, err := twiml.Messages([]twiml.Element{message})
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
		} else {
			context.Header("Content-Type", "text/xml")
			context.String(http.StatusOK, twimlResult)
		}
		fmt.Println(twimlResult)
	})

	router.Run(":3000")
}
