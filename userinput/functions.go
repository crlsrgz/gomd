package userinput

import (
	"strings"
)

func BuildIndexList(lineText string) string {
	outResult := ""

	if strings.Contains(lineText, "# ") {

		lineText = strings.TrimSpace(lineText)
		lineText = strings.ToLower(lineText)
		linkName := ""
		linkAdress := ""

		switch {
		case strings.Contains(lineText, "###### ") || strings.Contains(lineText, "##### ") || strings.Contains(lineText, "#### "):
			break
		case strings.Contains(lineText, "# ") && !strings.Contains(lineText, "## "):
			lineText = strings.Replace(lineText, "# ", "", -1)
			linkName = "- [" + lineText + "]"
			linkAdress = strings.Replace(lineText, " ", "-", -1)
			linkAdress = "(#" + linkAdress + ")"
		case strings.Contains(lineText, "## ") && !strings.Contains(lineText, "### "):
			lineText = strings.Replace(lineText, "## ", "", -1)
			linkName = "  - [" + lineText + "]"
			linkAdress = strings.Replace(lineText, " ", "-", -1)
			linkAdress = "(#" + linkAdress + ")"
		case strings.Contains(lineText, "### ") && !strings.Contains(lineText, "#### "):
			lineText = strings.Replace(lineText, "### ", "", -1)
			linkName = "    - [" + lineText + "]"
			linkAdress = strings.Replace(lineText, " ", "-", -1)
			linkAdress = "(#" + linkAdress + ")"

		}
		linkForIndex := linkName + linkAdress

		// length := strconv.Itoa(len(lineText))
		// outResult = outResult + linkForIndex + length + lineText + "\n"
		outResult = outResult + linkForIndex + "\n"
	}

	return outResult
}
