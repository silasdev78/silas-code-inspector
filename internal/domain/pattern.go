package domain

type Pattern struct {
	ID             string
	Title          string
	Regex          string
	Severity       Severity
	Description    string
	Recommendation string
}
