package tact

import (
	"regexp"
	"strings"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
	"github.com/silasdev78/silas-code-inspector/internal/knowledge"
)

type Scanner struct {
	patterns []domain.Pattern
}

func NewScanner() *Scanner {
	return &Scanner{patterns: knowledge.TactPatterns()}
}

func (s *Scanner) Scan(source string) []domain.Issue {
	cleanSource := removeComments(source)
	lines := strings.Split(source, "\n")
	var issues []domain.Issue

	for _, pattern := range s.patterns {
		re, err := regexp.Compile(pattern.Regex)
		if err != nil {
			continue
		}
		matches := re.FindAllStringIndex(cleanSource, -1)
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
				Column:         match[0],
				Snippet:        snippet,
				Recommendation: pattern.Recommendation,
			})
		}
	}
	return issues
}

func removeComments(source string) string {
	// remove /* ... */ comments
	block := regexp.MustCompile(`/\*[\s\S]*?\*/`)
	out := block.ReplaceAllString(source, "")
	// remove // line comments
	line := regexp.MustCompile(`//.*`)
	out = line.ReplaceAllString(out, "")
	return out
}

func lineNumber(source string, offset int) int {
	return strings.Count(source[:offset], "\n") + 1
}
