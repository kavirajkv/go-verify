package api

import (
	"go-verify/models"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)


func twilioclient()*twilio.RestClient{
	client:=twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	return client
}

func SendOTP(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": "Invalid request"})
	}

	if user.PhoneNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"error": "Phone number is required"})
	}

	params := &openapi.CreateVerificationParams{}
	params.SetTo(user.PhoneNumber)
	params.SetChannel("sms")

	client:=twilioclient()

	res, err := client.VerifyV2.CreateVerification(os.Getenv("TWILIO_SERVICE_SID"), params)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	} else {
		return c.Status(200).JSON(&fiber.Map{"response": res})
	}
}


func VerifyOTP(c *fiber.Ctx) error {
	var verify models.OTP
	
	err:=c.BodyParser(&verify)
	if err!=nil{
		c.Status(400).SendString(err.Error())
	}

	if verify.PhoneNumber=="" || verify.OTP==""{
		return c.Status(400).SendString("Enter both phone and otp")
	}

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(verify.PhoneNumber)
	params.SetCode(verify.OTP)

	client:=twilioclient()

	_,error:= client.VerifyV2.CreateVerificationCheck(os.Getenv("TWILIO_SERVICE_SID"), params)

	if error!=nil{
		c.Status(400).SendString(err.Error())
	}

	return c.Status(200).SendString("OTP verification successful")

}