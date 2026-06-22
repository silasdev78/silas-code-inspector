package golang

import (
	"regexp"
	"strings"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
	gopatterns "github.com/silasdev78/silas-code-inspector/internal/knowledge/golang"
)

type Scanner struct {
	patterns []domain.Pattern
}

func NewScanner() *Scanner {
	return &Scanner{patterns: gopatterns.Patterns()}
}

func (s *Scanner) Scan(source string) []domain.Issue {
	var issues []domain.Issue
	lines := strings.Split(source, "\n")
	for _, pattern := range s.patterns {
		re, err := regexp.Compile(pattern.Regex)
		if err != nil {
			continue
		}
		matches := re.FindAllStringIndex(source, -1)
		for _, match := range matches {
			lineNum := lineNumber(source, match[0])
			snippet := ""
			if lineNum > 0 && lineNum <= len(lines) {
				snippet = strings.TrimSpace(lines[lineNum-1])
			}
			issues = append(issues, domain.Issue{
				Title:          pattern.Title,
				Description:    pattern.Description,
				Severity:       pattern.Severity,
				Line:           lineNum,
				Snippet:        snippet,
				Recommendation: pattern.Recommendation,
			})
		}
	}
	return issues
}

func lineNumber(source string, offset int) int {
	return strings.Count(source[:offset], "\n") + 1
}
