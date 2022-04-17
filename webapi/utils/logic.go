package utils

import "regexp"

type logic struct {
}

var Logic logic

func (logic) GetPasswordStrength(passwd string) (passwdLevel int) {
	counter := 0
	numberMatch := "[0-9]+"
	lowLetter := "[a-z]+"
	upLetter := "[A-Z]+"
	specialSymbol := `[*()~!@#$%^&*-+=_|:;'<>,.?/\[\]\{\}<>]+`
	if match, _ := regexp.MatchString(numberMatch, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(lowLetter, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(upLetter, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(specialSymbol, passwd); match {
		counter++
	}
	if len(passwd) < 8 || counter <= 1 {
		passwdLevel = 0
	} else if counter <= 2 {
		passwdLevel = 1
	} else if counter <= 3 {
		passwdLevel = 2
	} else {
		passwdLevel = 3
	}
	return
}
