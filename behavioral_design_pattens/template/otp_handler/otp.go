package otphandler

type IOtp interface {
	GetRandomOtp(int) string
	SaveOtpToCache(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	IOtp IOtp
}

func (o *Otp) GetAndSendOTP(otpLength int) error {
	otp := o.IOtp.GetRandomOtp(otpLength)
	o.IOtp.SaveOtpToCache(otp)
	message := o.IOtp.GetMessage(otp)
	return o.IOtp.SendNotification(message)
}
