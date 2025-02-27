package morph

import (
	"text/template"
)

type QueryOptions struct {
	Placeholder string
	Named       bool
	OmitEmpty   bool
}

type QueryOption func(*QueryOptions)

func WithPlaceholder(p string) QueryOption {
	return func(q *QueryOptions) {
		q.Placeholder = p
	}
}

func WithNamedParameters() QueryOption {
	return func(q *QueryOptions) {
		q.Named = true
	}
}

func WithoutEmptyValues() QueryOption {
	return func(q *QueryOptions) {
		q.OmitEmpty = true
	}
}

const insertSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  INSERT INTO {{$table.Name}} (
  {{- range $idx, $col := $table.Columns -}}
    {{$col.Name}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end -}}
  ) VALUES (
  {{- range $idx, $col := $table.Columns -}}
    {{param $col.Name $options}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end -}}
  );`

const updateSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  UPDATE {{$table.Name}} AS {{$table.Alias}} SET {{- if true}} {{end}}
  {{- range $idx, $col := $table.Columns -}}
    {{- if $col.PrimaryKey -}} {{continue}} {{- end -}}
    {{$table.Alias}}.{{.Name}} = {{param $col.Name $options}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end }} WHERE 1=1
  {{- range $idx, $col := $table.Columns -}}
    {{- if $col.PrimaryKey }} AND {{$table.Alias}}.{{.Name}} = {{param $col.Name $options}}{{- end -}}
  {{- end -}};`

const deleteSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  DELETE FROM {{$table.Name}} WHERE 1=1
  {{- range $idx, $col := $table.Columns -}}
    {{- if $col.PrimaryKey }} AND {{.Name}} = {{param $col.Name $options}}{{- end -}}
  {{- end -}};`

var (
	funcs = template.FuncMap{
		"param": func(columnName string, options *QueryOptions) string {
			if options.Named {
				return ":" + columnName
			}
			return options.Placeholder
		},
		"sub": func(a, b int) int {
			return a - b
		},
	}
	insertTmpl = template.Must(template.New("insertQuery").Funcs(funcs).Parse(insertSQL))
	updateTmpl = template.Must(template.New("updateQuery").Funcs(funcs).Parse(updateSQL))
	deleteTmpl = template.Must(template.New("deleteQuery").Funcs(funcs).Parse(deleteSQL))
)
