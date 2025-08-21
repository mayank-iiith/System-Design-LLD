package main

// Template Method is a behavioral design pattern that allows you to define a skeleton of an algorithm in a base class and let subclasses override the steps without changing the overall algorithm’s structure.

import (
	"fmt"
	otphandler "lld/behavioral_design_pattens/template/otp_handler"
	"log"
)

func main() {
	smsOtpHandler := &otphandler.Otp{
		IOtp: &otphandler.Sms{},
	}
	err := smsOtpHandler.GetAndSendOTP(4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	emailOtpHandler := &otphandler.Otp{
		IOtp: &otphandler.Email{},
	}
	err = emailOtpHandler.GetAndSendOTP(6)
	if err != nil {
		log.Fatal(err)
	}
}
