package errors

import (
	"github.com/lxn/walk"
	. "main/const"
	. "main/tools"
	"unicode/utf8"
)

func CheckErrors(text string, keys []*walk.TextEdit, errors CiphersErrors) (string, string){
	newCipherPackage :=CipherPackage{
		Text:   text,
		Keys:   keys,
		Cipher: errors,
	}
	var error string
	if error := GetErrorText(newCipherPackage); utf8.RuneCountInString(error)!=0{
		return "", error
	}
	result := errors.Cipher(text, keys)
	return result, error
}
