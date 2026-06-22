package report

import (
	"encoding/json"
	"os"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
)

type Result struct {
	File   string         `json:"file"`
	Issues []domain.Issue `json:"issues"`
}

type SARIFReport struct {
	Schema  string     `json:"$schema"`
	Version string     `json:"version"`
	Runs    []SARIFRun `json:"runs"`
}

type SARIFRun struct {
	Tool    SARIFTool     `json:"tool"`
	Results []SARIFResult `json:"results"`
}

type SARIFTool struct {
	Driver SARIFDriver `json:"driver"`
}

type SARIFDriver struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type SARIFResult struct {
	RuleID    string          `json:"ruleId"`
	Level     string          `json:"level"`
	Message   SARIFMessage    `json:"message"`
	Locations []SARIFLocation `json:"locations"`
}

type SARIFMessage struct {
	Text string `json:"text"`
}

type SARIFLocation struct {
	PhysicalLocation SARIFPhysicalLocation `json:"physicalLocation"`
}

type SARIFPhysicalLocation struct {
	ArtifactLocation SARIFArtifactLocation `json:"artifactLocation"`
	Region           SARIFRegion           `json:"region"`
}

type SARIFArtifactLocation struct {
	URI string `json:"uri"`
}

type SARIFRegion struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
}

func WriteJSON(results []Result, file *os.File) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}

func WriteSARIF(results []Result, file *os.File) error {
	run := SARIFRun{
		Tool: SARIFTool{
			Driver: SARIFDriver{
				Name:    "Silas Code Inspector",
				Version: "0.2.0",
			},
		},
	}

	for _, res := range results {
		for _, issue := range res.Issues {
			sarifResult := SARIFResult{
				RuleID: issue.Title,
				Level:  severityToSARIFLevel(issue.Severity),
				Message: SARIFMessage{
					Text: issue.Description,
				},
				Locations: []SARIFLocation{
					{
						PhysicalLocation: SARIFPhysicalLocation{
							ArtifactLocation: SARIFArtifactLocation{URI: res.File},
							Region: SARIFRegion{
								StartLine:   issue.Line,
								StartColumn: issue.Column + 1,
							},
						},
					},
				},
			}
			run.Results = append(run.Results, sarifResult)
		}
	}

	sarif := SARIFReport{
		Schema:  "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
		Version: "2.1.0",
		Runs:    []SARIFRun{run},
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(sarif)
}

func severityToSARIFLevel(s domain.Severity) string {
	switch s {
	case domain.SeverityCritical, domain.SeverityHigh:
		return "error"
	case domain.SeverityMedium:
		return "warning"
	case domain.SeverityLow:
		return "note"
	default:
		return "none"
	}
}
