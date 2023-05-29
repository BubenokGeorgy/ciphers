package data

import (
	. "main/const"
	"os"
)

func GetProverb() string{
	data, _ := os.ReadFile(PathProverb)
	text := string(data)
	return text
}

func GetText() string{
		data, _ := os.ReadFile(PathText)
		text := string(data)
	return text
}
