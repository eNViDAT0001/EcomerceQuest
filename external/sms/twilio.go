package sms

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/chat/v1"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

var twilioClient *twilio.RestClient

func Getclient() *twilio.RestClient {
	if twilioClient != nil {
		return twilioClient
	}
	restClient := twilio.ClientParams{
		Username: "AC49bec91da6a5766af6c96df0160f7939",
		Password: "f48e51e29a9a7fc86e15cf920e23a87f",
	}
	twilioClient = twilio.NewRestClientWithParams(restClient)
	return twilioClient
}

func SendMessage(msg string, to string) {
	params := openapi.CreateMessageParams{
		Body:       nil,
		From:       nil,
		Attributes: nil,
	}
}
func main() {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure

	client := twilio.NewRestClientWithParams(restClient)
	params := &verify.CreateServiceParams{}
	params.SetFriendlyName("My First Verify Service")

	resp, err := client.VerifyV2.CreateService(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}
