package telegram

import (
	"strconv"
)

func int64toUUID(chatID int64) string {
	hexStr := strconv.FormatInt(chatID, 16)
	// format hexStr to 16 length
	for len(hexStr) < 16 {
		hexStr = "0" + hexStr
	}
	return "58699cb7-14ba-4d96-" + hexStr[:4] + "-" + hexStr[4:]
}
