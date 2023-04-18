package utils

import (
	"regexp"
	"strings"
)

func ExtractCommand(response string) string {
	if !strings.Contains(response, "```") {
		return ""
	}
	re := regexp.MustCompile("(?s)```(.*?)```")
	matches := re.FindStringSubmatch(response)

	if len(matches) > 1 {
		return strings.TrimSpace(strings.Trim(matches[0], "`"))
	}

	return ""
}
