package domain

type Severity string

const (
	SeverityInfo     Severity = "INFO"
	SeverityLow      Severity = "LOW"
	SeverityMedium   Severity = "MEDIUM"
	SeverityHigh     Severity = "HIGH"
	SeverityCritical Severity = "CRITICAL"
)

type Issue struct {
	Title          string
	Description    string
	Severity       Severity
	Line           int
	Column         int
	Snippet        string
	Recommendation string
}
