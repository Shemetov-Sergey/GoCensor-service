package textParser

import (
	"strings"

	"github.com/Shemetov-Sergey/GoCensor-service/pkg/models"
)

func CheckCensored(text string, censored []*models.CensoredWords) bool {
	words := strings.Split(text, " ")
	for _, word := range words {
		for _, cw := range censored {
			if word == cw.Word {
				return true
			}
		}
	}

	return false
}
