
{{template "base/base.html" .}}
{{define "meta"}}
{{end}}
{{define "head"}}
{{end}}
{{define "body"}}
{{.MarkdownContent |str2html}}
{{end}}


