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
		case strings.Contains(lineText, "##### "):
			break
		case strings.Contains(lineText, "#### "):
			break
		case strings.Contains(lineText, "### "):
			lineText = strings.Replace(lineText, "### ", "", -1)
			linkName = "    - [" + lineText + "]"
			linkAdress = strings.Replace(lineText, " ", "-", -1)
			linkAdress = "(#" + linkAdress + ")"
		case strings.Contains(lineText, "## "):
			lineText = strings.Replace(lineText, "## ", "", -1)
			linkName = "  - [" + lineText + "]"
			linkAdress = strings.Replace(lineText, " ", "-", -1)
			linkAdress = "(#" + linkAdress + ")"
		case strings.Contains(lineText, "# "):
			lineText = strings.Replace(lineText, "# ", "", -1)
			linkName = "- [" + lineText + "]"
			linkAdress = strings.Replace(lineText, " ", "-", -1)
			linkAdress = "(#" + linkAdress + ")"

		}
		// lineText = strings.Replace(lineText, "## ", "", -1)

		// linkName := "- [" + lineText + "]"

		linkForIndex := linkName + linkAdress

		outResult = outResult + linkForIndex + "\n"
	}

	return outResult
}
