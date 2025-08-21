package otphandler

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Email struct {
}

func (e Email) GetRandomOtp(otpLength int) string {
	otp := ""
	for range otpLength {
		// Generate a random digit between 0 and 9
		otp += strconv.Itoa(rand.Intn(10))
	}
	fmt.Printf("Email: generating random otp %s\n", otp)
	return otp
}

func (e Email) SaveOtpToCache(otp string) {
	fmt.Printf("Email: Saving otp %s to cache\n", otp)
}

func (e Email) GetMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e Email) SendNotification(message string) error {
	fmt.Printf("Email: sending email: %s\n", message)
	return nil
}
