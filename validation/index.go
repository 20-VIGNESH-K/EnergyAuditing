package validation

import "regexp"

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}`)
	return emailRegex.MatchString(email)
}


func IsValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := regexp.MustCompile(`^[0-9]{10}$`)
	return phoneRegex.MatchString(phoneNumber)
}


func IsValidName(name string) bool {
	bankRegex := regexp.MustCompile(`^[a-zA-Z0-9.,_-]+(?:\s[a-zA-Z0-9]+)*$`)
	return bankRegex.MatchString(name)
}