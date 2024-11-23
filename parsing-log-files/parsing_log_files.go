package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(DBG|ERR|FTL|INF|TRC|WRN)\] `).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*-=]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {
	re := regexp.MustCompile(`".*(?i)(password).*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\s*User \s*(\S+)`)
	for i, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			lines[i] = fmt.Sprintf("[USR] %s %s", match[1], line)
		}
	}
	return lines
}
