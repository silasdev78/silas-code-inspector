package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
	"github.com/silasdev78/silas-code-inspector/internal/engine/tact"
)

var rootCmd = &cobra.Command{
	Use:   "silas <file.tact>",
	Short: "Silas Code Inspector - TON smart contract scanner",
	Long:  "A static analysis tool for finding security vulnerabilities in Tact/TON smart contracts.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		if ext := strings.ToLower(filepath.Ext(path)); ext != ".tact" {
			fmt.Fprintf(os.Stderr, "Error: only .tact files are supported, got %q\n", ext)
			os.Exit(1)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}

		scanner := tact.NewScanner()
		issues := scanner.Scan(string(data))

		if len(issues) == 0 {
			color.Green("✓ No vulnerabilities found!")
			return
		}

		color.Red("✗ Found %d potential issues:\n", len(issues))

		for _, issue := range issues {
			printIssue(issue)
		}
	},
}

func printIssue(issue domain.Issue) {
	sevStr := severityColor(issue.Severity)
	bold := color.New(color.Bold).SprintfFunc()
	fmt.Printf("\n%s\n", bold(issue.Title))
	fmt.Printf("  Severity: %s\n", sevStr)
	fmt.Printf("  Line: %d\n", issue.Line)
	if issue.Snippet != "" {
		fmt.Printf("  Code: %s\n", color.YellowString(issue.Snippet))
	}
	fmt.Printf("  Issue: %s\n", issue.Description)
	fmt.Printf("  Fix: %s\n", color.GreenString(issue.Recommendation))
}

func severityColor(s domain.Severity) string {
	switch s {
	case domain.SeverityCritical:
		return color.MagentaString(string(s))
	case domain.SeverityHigh:
		return color.RedString(string(s))
	case domain.SeverityMedium:
		return color.YellowString(string(s))
	case domain.SeverityLow:
		return color.BlueString(string(s))
	default:
		return color.WhiteString(string(s))
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
