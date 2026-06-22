package docker

import "github.com/silasdev78/silas-code-inspector/internal/domain"

func Patterns() []domain.Pattern {
	return []domain.Pattern{
		{ID: "DOCK-001", Title: "Root user", Regex: `^USER\s+root|^USER\s+0`, Severity: domain.SeverityHigh, Description: "Container runs as root.", Recommendation: "Create a non-root user."},
		{ID: "DOCK-002", Title: "No healthcheck", Regex: `HEALTHCHECK`, Severity: domain.SeverityMedium, Description: "No HEALTHCHECK instruction.", Recommendation: "Add HEALTHCHECK to monitor container health."},
		{ID: "DOCK-003", Title: "Latest tag used", Regex: `FROM\s+\S+:latest`, Severity: domain.SeverityLow, Description: "Using latest tag can cause unpredictable builds.", Recommendation: "Pin a specific version."},
		{ID: "DOCK-004", Title: "Unnecessary ports exposed", Regex: `EXPOSE\s+\d+`, Severity: domain.SeverityInfo, Description: "Ports exposed; ensure they are intentional.", Recommendation: "Review exposed ports."},
	}
}
