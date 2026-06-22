package engine

import (
	"fmt"

	"github.com/silasdev78/silas-code-inspector/internal/domain"
	"github.com/silasdev78/silas-code-inspector/internal/engine/docker"
	funcc "github.com/silasdev78/silas-code-inspector/internal/engine/func"
	"github.com/silasdev78/silas-code-inspector/internal/engine/golang"
	"github.com/silasdev78/silas-code-inspector/internal/engine/gomod"
	"github.com/silasdev78/silas-code-inspector/internal/engine/tact"
	"github.com/silasdev78/silas-code-inspector/internal/engine/web"
)

type Scanner interface {
	Scan(source string) []domain.Issue
}

func NewScanner(lang string) (Scanner, error) {
	switch lang {
	case "tact", "ton":
		return tact.NewScanner(), nil
	case "go", "golang":
		return golang.NewScanner(), nil
	case "docker", "dockerfile":
		return docker.NewScanner(), nil
	case "web", "html", "js":
		return web.NewScanner(), nil
	case "gomod", "go.mod":
		return gomod.NewScanner(), nil
	case "func", "fift", "fc":
		return funcc.NewScanner(), nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", lang)
	}
}
