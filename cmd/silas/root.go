package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
	"github.com/silasdev78/silas-code-inspector/internal/engine"
	"github.com/silasdev78/silas-code-inspector/internal/learner"
	"github.com/silasdev78/silas-code-inspector/internal/report"
)

var (
	langFlag    string
	outputFlag  string
	learnerFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "silas [file or directory]",
	Short: "Silas Code Inspector - multi-language security scanner",
	Long:  "Scans TON, Go, Docker, Web, and Go module files for vulnerabilities.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]

		var l *learner.Learner
		if learnerFlag {
			lr, err := learner.NewLearner(".silas-state.json")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error initializing learner: %v\n", err)
			} else {
				l = lr
			}
		}

		info, err := os.Stat(target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		var results []report.Result
		if info.IsDir() {
			results = scanDirectory(target, l)
		} else {
			res := scanFile(target, l)
			results = append(results, res)
		}

		switch outputFlag {
		case "json":
			if err := report.WriteJSON(results, os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing JSON: %v\n", err)
			}
		case "sarif":
			if err := report.WriteSARIF(results, os.Stdout); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing SARIF: %v\n", err)
			}
		default:
			// text output already printed by scanFile
		}

		if learnerFlag && l != nil {
			if err := l.Save(); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving learner state: %v\n", err)
			}
		}
	},
}

func scanFile(path string, l *learner.Learner) report.Result {
	lang := detectLang(path)
	if lang == "" {
		color.Yellow("Skipping unsupported file: %s\n", path)
		return report.Result{File: path}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
		return report.Result{File: path}
	}

	scanner, err := engine.NewScanner(lang)
	if err != nil {
		color.Yellow("Unsupported language for %s: %v\n", path, err)
		return report.Result{File: path}
	}

	issues := scanner.Scan(string(data))

	if l != nil {
		var filtered []domain.Issue
		for _, issue := range issues {
			weight := l.GetWeight(issue.Title)
			if weight >= 0.5 {
				filtered = append(filtered, issue)
			}
		}
		issues = filtered
	}

	printResults(path, issues)
	return report.Result{File: path, Issues: issues}
}

func scanDirectory(dir string, l *learner.Learner) []report.Result {
	var files []string
	extensions := map[string]string{
		".tact": "tact",
		".go":   "go",
		".html": "web",
		".js":   "web",
		".ts":   "web",
		".mod":  "gomod",
	}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		base := filepath.Base(path)
		if strings.EqualFold(base, "dockerfile") || strings.HasPrefix(strings.ToLower(base), "dockerfile") {
			files = append(files, path)
			return nil
		}
		if base == "go.mod" {
			files = append(files, path)
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if _, ok := extensions[ext]; ok {
			files = append(files, path)
		}
		return nil
	})

	var results []report.Result
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			res := scanFile(f, l)
			mu.Lock()
			results = append(results, res)
			mu.Unlock()
		}(file)
	}
	wg.Wait()
	return results
}

func detectLang(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	base := filepath.Base(path)
	switch {
	case base == "go.mod":
		return "gomod"
	case ext == ".tact":
		return "tact"
	case ext == ".go":
		return "go"
	case ext == ".html" || ext == ".js" || ext == ".ts":
		return "web"
	case strings.EqualFold(base, "dockerfile") || strings.HasPrefix(strings.ToLower(base), "dockerfile"):
		return "docker"
	default:
		return ""
	}
}

func printResults(path string, issues []domain.Issue) {
	if outputFlag != "text" && outputFlag != "" {
		return
	}
	if len(issues) == 0 {
		color.Green("✓ %s: No vulnerabilities found.\n", path)
		return
	}
	color.Red("✗ %s: %d issues found.\n", path, len(issues))
	for _, issue := range issues {
		printIssue(issue)
	}
}

func printIssue(issue domain.Issue) {
	sevStr := severityColor(issue.Severity)
	bold := color.New(color.Bold).SprintfFunc()
	fmt.Printf("  • %s\n", bold(issue.Title))
	fmt.Printf("    Severity: %s | Line: %d\n", sevStr, issue.Line)
	if issue.Snippet != "" {
		fmt.Printf("    Code: %s\n", color.YellowString(issue.Snippet))
	}
	fmt.Printf("    Fix: %s\n\n", color.GreenString(issue.Recommendation))
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
	rootCmd.Flags().StringVarP(&langFlag, "lang", "l", "", "Force language (go, docker, web, tact, gomod)")
	rootCmd.Flags().StringVarP(&outputFlag, "output", "o", "text", "Output format: text, json, sarif")
	rootCmd.Flags().BoolVar(&learnerFlag, "learner", false, "Enable adaptive learning (requires .silas-state.json)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
