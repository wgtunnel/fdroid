package md

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"metascoop/apps"
)

const (
	tableStart = "<!-- This table is auto-generated. Do not edit -->"
	tableEnd   = "<!-- end apps table -->"

	tableTmpl = `
| Icon | Name | Description | Version |
| --- | --- | --- | --- |{{range .Apps}}
| <a href="{{.sourceCode}}"><img src="fdroid/repo/{{.packageName}}/en-US/icon.png" alt="{{.name}} icon" width="36px" height="36px"></a> | [**{{.name}}**]({{.sourceCode}}) | {{.summary}} | {{.suggestedVersionName}} ({{.suggestedVersionCode}}) |{{end}}
` + tableEnd

	siteAppsStart = "<!-- This apps list is auto-generated. Do not edit -->"
	siteAppsEnd   = "<!-- end apps list -->"

	siteAppsHTML = `
        <div class="apps">{{range .Apps}}
          <a class="app-card" href="{{.sourceCode}}">
            <img
              src="fdroid/repo/{{.packageName}}/en-US/icon.png"
              alt="{{.name}} icon"
            />
            <div>
              <h3>{{.name}}</h3>
              <p>{{.summary}}</p>
            </div>
            <span class="version">{{.suggestedVersionName}}</span>
          </a>{{end}}
        </div>
` + siteAppsEnd
)

var (
	readmeTmpl = template.Must(template.New("readme").Parse(tableTmpl))
	siteTmpl   = template.Must(template.New("site").Parse(siteAppsHTML))
)

func RegenerateReadme(readMePath string, index *apps.RepoIndex) (err error) {
	content, err := os.ReadFile(readMePath)
	if err != nil {
		return
	}

	tableStartIndex := bytes.Index(content, []byte(tableStart))
	if tableStartIndex < 0 {
		return fmt.Errorf("cannot find table start in %q", readMePath)
	}

	tableEndIndex := bytes.Index(content, []byte(tableEnd))
	if tableEndIndex < 0 {
		return fmt.Errorf("cannot find table end in %q", readMePath)
	}

	var table bytes.Buffer
	table.WriteString(tableStart)
	if err = readmeTmpl.Execute(&table, index); err != nil {
		return err
	}

	newContent := append([]byte{}, content[:tableStartIndex]...)
	newContent = append(newContent, table.Bytes()...)
	newContent = append(newContent, content[tableEndIndex:]...)

	return os.WriteFile(readMePath, newContent, os.ModePerm)
}

// RegenerateSite updates the auto-generated apps list inside the GitHub Pages index.html.
func RegenerateSite(sitePath string, index *apps.RepoIndex) error {
	content, err := os.ReadFile(sitePath)
	if err != nil {
		return err
	}

	start := bytes.Index(content, []byte(siteAppsStart))
	if start < 0 {
		return fmt.Errorf("cannot find apps list start in %q", sitePath)
	}
	end := bytes.Index(content, []byte(siteAppsEnd))
	if end < 0 {
		return fmt.Errorf("cannot find apps list end in %q", sitePath)
	}

	var appsHTML bytes.Buffer
	appsHTML.WriteString(siteAppsStart)
	if err := siteTmpl.Execute(&appsHTML, index); err != nil {
		return err
	}

	// F-Droid index text often already includes entities; avoid ugly double-escaping.
	rendered := appsHTML.String()
	rendered = strings.ReplaceAll(rendered, "&amp;amp;", "&amp;")
	rendered = strings.ReplaceAll(rendered, "&#43;", "+")

	newContent := append([]byte{}, content[:start]...)
	newContent = append(newContent, []byte(rendered)...)
	newContent = append(newContent, content[end:]...)

	return os.WriteFile(sitePath, newContent, os.ModePerm)
}

// SitePathFromRepoDir returns repo-root index.html given the fdroid/repo directory.
func SitePathFromRepoDir(repoDir string) string {
	return filepath.Join(filepath.Dir(filepath.Dir(repoDir)), "index.html")
}
