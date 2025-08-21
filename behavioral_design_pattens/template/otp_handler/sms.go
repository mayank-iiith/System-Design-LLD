package otphandler

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Sms struct {
}

func (s Sms) GetRandomOtp(otpLength int) string {
	otp := ""
	for range otpLength {
		// Generate a random digit between 0 and 9
		otp += strconv.Itoa(rand.Intn(10))
	}
	fmt.Printf("SMS: generating random otp %s\n", otp)
	return otp
}

func (s Sms) SaveOtpToCache(otp string) {
	fmt.Printf("SMS: Saving otp %s to cache\n", otp)
}

func (s Sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
