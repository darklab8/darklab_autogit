package changelog

import (
	"autogit/parser/conventionalcommits"
	sGit "autogit/parser/semanticGit"
	"autogit/settings"
	"autogit/utils"
	"fmt"
	"strings"
	"time"

	_ "embed"
)

type Header struct {
	From    string
	To      string
	Version string
}

func (r *Header) Render() string {
	currentTime := time.Now()
	return fmt.Sprintf("## **%s** <sub><sub>%s ([%s...%s](%s))</sub></sub>", r.Version, currentTime.Format("2006-01-02"), r.From, r.To, utils.TmpRender(settings.Template.CommitRangeUrl, r))
}

func (r *Header) New(logs []conventionalcommits.ConventionalCommit, version string) *Header {
	r.From = logs[len(logs)-1].Hash
	r.To = logs[0].Hash
	r.Version = version
	return r
}

type commitRecord struct {
	Commit string
}

func (c commitRecord) Render(record conventionalcommits.ConventionalCommit) string {
	type IssueData struct {
		Issue string
	}

	var issue_rendered strings.Builder
	for _, issue_n := range record.Issue {
		issue_rendered.WriteString(fmt.Sprintf(", [#%s](%s)", issue_n, utils.TmpRender(settings.Template.IssueUrl, IssueData{Issue: issue_n})))
	}

	rendered_subject := record.Subject
	IssueMatch := conventionalcommits.IssueRegex.FindAllStringSubmatch(record.Subject, -1)
	for _, match := range IssueMatch {
		rendered_subject = strings.Replace(rendered_subject, match[0], fmt.Sprintf("[#%s](%s)", match[1], utils.TmpRender(settings.Template.IssueUrl, IssueData{Issue: match[1]})), -1)
	}

	formatted_url := utils.TmpRender(settings.Template.CommitUrl, commitRecord{Commit: record.Hash})
	return fmt.Sprintf("* %s ([%s](%s)%s)\n", rendered_subject, record.Hash, formatted_url, issue_rendered.String())
}

type ChangelogData struct {
	Header   string
	Features []string
	Fixes    []string
}

func (changelog ChangelogData) New(g *sGit.SemanticGit) ChangelogData {
	logs := g.GetChangelogByTag("", true)

	changelog.Header = (&Header{}).New(logs, g.GetNextVersion().ToString()).Render()

	for _, record := range logs {
		commit_formatted := commitRecord{Commit: record.Hash}.Render(record)
		if record.Type == "feat" {
			changelog.Features = append(changelog.Features, commit_formatted)
		} else if record.Type == "fix" {
			changelog.Fixes = append(changelog.Fixes, commit_formatted)
		}
	}

	return changelog
}

func (changelog ChangelogData) Render() string {
	return utils.TmpRender(settings.Template.Changelog, changelog)
}