package detector

import (
	"strings"

	"github.com/go-rod/rod"
)

var fieldKeywords = map[string]string{
	"name":     "fullName",
	"email":    "email",
	"phone":    "phone",
	"address":  "address",
	"city":     "city",
	"state":    "state",
	"zip":      "zipCode",
	"linkedin": "linkedin",
	"github":   "github",
}

func DetectFields(page *rod.Page) []*rod.Element {
	// finds all possible elements to enter in the form
	return page.MustElements("input, textarea, select")
}

func IdentifyFieldType(element *rod.Element) string {
	// extract common attributes
	id, _ := element.Attribute("id")
	name, _ := element.Attribute("name")
	placeholder, _ := element.Attribute("placeholder")

	// match each attribute against our keywords
	for keyword, fieldType := range fieldKeywords {
		if containsIgnoreCase(*id, keyword) ||
			containsIgnoreCase(*name, keyword) ||
			containsIgnoreCase(*placeholder, keyword) {
			return fieldType
		}
	}

	return ""
}

func containsIgnoreCase(s, substr string) bool {
	if s == "" || substr == "" {
		return false
	}
	return strings.Contains(
		strings.ToLower(s),
		strings.ToLower(substr),
	)
}
