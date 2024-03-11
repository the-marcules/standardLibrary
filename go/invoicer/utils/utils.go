package utils

import (
	"time"
)

func GetCurrentDateAndTimeString() string {
	return time.Now().String()
}

func GenerateInvoiceNumber() string {
	return "123aabbcc"
}
