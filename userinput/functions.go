package userinput

import (
	"strings"
)

func BuildIndexList(lineText string) string {
	outResult := ""

	if strings.Contains(lineText, "## ") && !strings.Contains(lineText, "### ") {

		lineText = strings.TrimSpace(lineText)
		lineText = strings.ToLower(lineText)
		lineText = strings.Replace(lineText, "## ", "", -1)

		linkName := "- [" + lineText + "]"

		linkAdress := strings.Replace(lineText, " ", "-", -1)
		linkAdress = "(#" + linkAdress + ")"

		linkForIndex := linkName + linkAdress

		outResult = outResult + linkForIndex + "\n"
	}

	return outResult
}
