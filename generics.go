package main

import (
	"fmt"
	"strconv"
	"strings"
)

// type want struct {
// 	name    string
// 	age     int
// 	contact struct {
// 		phone  string
// 		mobile string
// 		email  string
// 	}
// 	appointment time.Time
// }

const phoneNumberLength = 10

// func DataMap(data map[string]any) want {
// 	w := want{}
// 	return w
// }

// standard phone helper function US ONLY
// return phone number with format xxx-xxx-xxxx
func makePhone(number interface{}) (string, error) {
	var phone string
	// convert to string
	switch v := number.(type) {
	case int:
		phone = strconv.Itoa(v)
	default:
		phone = v.(string)
	}

	// clean up phone number
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	phone = strings.ReplaceAll(phone, "-", "")

	// check length of use base codes
	if len(phone) != phoneNumberLength {
		return "", fmt.Errorf("invalid phone number")
	}

	// TODO:do random area code analysis?

	phone = phone[:3] + "-" + phone[3:6] + "-" + phone[6:]
	return phone, nil
}
