package helper

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-gomail/gomail"
)


func GenerateOtp() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendOtp(email, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("appEmail"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verification Code for Signup")
	m.SetBody("text/plain", "Your OTP for signup is: "+otp)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("appEmail"), os.Getenv("appPassword"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
