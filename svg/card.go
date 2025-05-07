package svg

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/dustin/go-humanize"
	"github.com/tico88612/devstats-card/models"
)

func GenerateSVG(data models.CardData) string {
	tmplPath := filepath.Join("svg", "card.svg.tmpl")
	tmplBytes, err := os.ReadFile(tmplPath)
	if err != nil {
		return "<svg><text>Template not found</text></svg>"
	}
	funcMap := template.FuncMap{
		"formatNumber": func(n int) string {
			return humanize.Comma(int64(n))
		},
	}
	tmpl, err := template.New("svg").Funcs(funcMap).Parse(string(tmplBytes))
	if err != nil {
		return "<svg><text>Template error</text></svg>"
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "<svg><text>Render error</text></svg>"
	}

	return buf.String()
}

//func GenerateSVG(username string, prCount, rank int) string {
//	return fmt.Sprintf(`
//<svg xmlns="http://www.w3.org/2000/svg" width="300" height="80">
//  <rect width="100%%" height="100%%" fill="#0d1117" rx="10" />
//  <text x="20" y="30" fill="#58a6ff" font-size="16">DevStats for @%s</text>
//  <text x="20" y="55" fill="#c9d1d9" font-size="14">PRs: %d | Rank: #%d</text>
//  <text x="20" y="70" fill="#c9d1d9" font-size="14">PRs: %d | Rank: #%d</text>
//</svg>
//`, username, prCount, rank)
//}
