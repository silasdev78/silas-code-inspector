package golang

import "github.com/silasdev78/silas-code-inspector/internal/domain"

func Patterns() []domain.Pattern {
	return []domain.Pattern{
		{ID: "GO-001", Title: "Hardcoded secret", Regex: `(?i)(password|secret|api_key|token)\s*=\s*"[^"]+"`, Severity: domain.SeverityCritical, Description: "Hardcoded sensitive data found.", Recommendation: "Use environment variables or a vault."},
		{ID: "GO-002", Title: "Insecure TLS config", Regex: `tls\.Config\s*\{.*InsecureSkipVerify\s*:\s*true`, Severity: domain.SeverityHigh, Description: "TLS certificate verification disabled.", Recommendation: "Remove InsecureSkipVerify in production."},
		{ID: "GO-003", Title: "Unchecked error", Regex: `err\s*:=\s*[^;]+[^;]+\n\s*if\s+err`, Severity: domain.SeverityMedium, Description: "Potential unhandled error (pattern may give false positives).", Recommendation: "Check all errors."},
		{ID: "GO-004", Title: "Use of math/rand for crypto", Regex: `rand\.Intn|rand\.Float64`, Severity: domain.SeverityHigh, Description: "math/rand is not suitable for cryptographic purposes.", Recommendation: "Use crypto/rand."},
		{ID: "GO-005", Title: "SQL injection via fmt.Sprintf", Regex: `fmt\.Sprintf\s*\([^)]*SELECT|INSERT|UPDATE|DELETE`, Severity: domain.SeverityCritical, Description: "Potential SQL injection.", Recommendation: "Use parameterized queries."},
		{ID: "GO-006", Title: "Open file without close", Regex: `os\.Open\s*\(`, Severity: domain.SeverityMedium, Description: "File opened but deferred close may be missing.", Recommendation: "Add defer f.Close()."},
		{ID: "GO-007", Title: "Insecure file permissions", Regex: `os\.Create|ioutil\.WriteFile\s*\([^,]*,\s*0[0-7]{3}`, Severity: domain.SeverityLow, Description: "File created with explicit permissions.", Recommendation: "Use 0600 for sensitive files."},
		{ID: "GO-008", Title: "Logging of sensitive data", Regex: `log\.Print(f|ln)?\s*\([^)]*password|secret|token`, Severity: domain.SeverityLow, Description: "Sensitive data may be logged.", Recommendation: "Remove logging of secrets."},
	}
}
