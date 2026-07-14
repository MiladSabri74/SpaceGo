package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~*=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	counter := 0
	for _, line := range lines {
		if re.MatchString(line) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))
	re := regexp.MustCompile(`User\s+(\S+)`)
	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			userName := matches[1]
			result[i] = fmt.Sprintf("[USR] %s %s", userName, line)
		} else {
			result[i] = line
		}
	}
	return result
}
