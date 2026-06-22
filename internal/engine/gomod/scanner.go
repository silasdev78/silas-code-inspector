package gomod

import (
	"encoding/json"
	"os"
	"regexp"
	"strings"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
)

type cveEntry struct {
	Module      string `json:"module"`
	Version     string `json:"version"`
	CVE         string `json:"cve"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}

type Scanner struct {
	cves []cveEntry
}

func NewScanner() *Scanner {
	s := &Scanner{}
	data, err := os.ReadFile("internal/knowledge/cve/cve.json")
	if err == nil {
		json.Unmarshal(data, &s.cves)
	}
	return s
}

func (s *Scanner) Scan(source string) []domain.Issue {
	var issues []domain.Issue
	lines := strings.Split(source, "\n")
	requireRegex := regexp.MustCompile(`^\s*require\s*\(`)
	inBlock := false
	for _, line := range lines {
		if requireRegex.MatchString(line) {
			inBlock = true
			continue
		}
		if inBlock && strings.TrimSpace(line) == ")" {
			inBlock = false
			continue
		}
		if inBlock {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				module := parts[0]
				version := parts[1]
				for _, cve := range s.cves {
					if cve.Module == module {
						issues = append(issues, domain.Issue{
							Title:          "Vulnerable dependency: " + cve.CVE,
							Description:    cve.Description + " (module: " + module + " version: " + version + ")",
							Severity:       domain.Severity(cve.Severity),
							Line:           0,
							Snippet:        line,
							Recommendation: "Update " + module + " to version " + cve.Version,
						})
					}
				}
			}
		}
	}
	return issues
}
