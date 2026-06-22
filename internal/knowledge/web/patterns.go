package web

import "github.com/silasdev78/silas-code-inspector/internal/domain"

func Patterns() []domain.Pattern {
	return []domain.Pattern{
		{ID: "WEB-001", Title: "Inline script", Regex: `<script\b(?![^>]*\bsrc\s*=)[^>]*>`, Severity: domain.SeverityMedium, Description: "Inline JavaScript found; CSP may be weakened.", Recommendation: "Use external scripts with nonce/hash."},
		{ID: "WEB-002", Title: "Missing CSRF token", Regex: `<form\b[^>]*method\s*=\s*["']post`, Severity: domain.SeverityHigh, Description: "Form without CSRF token placeholder.", Recommendation: "Add a hidden CSRF token field."},
		{ID: "WEB-003", Title: "External HTTP resource", Regex: `http://[^'"]*\.(js|css|png)`, Severity: domain.SeverityLow, Description: "Resource loaded over HTTP (mixed content).", Recommendation: "Use HTTPS."},
		{ID: "WEB-004", Title: "Unsafe innerHTML", Regex: `\.innerHTML\s*=`, Severity: domain.SeverityCritical, Description: "Potential XSS via innerHTML.", Recommendation: "Use textContent or sanitize."},
	}
}
