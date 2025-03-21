package morph

import (
	"strconv"
	"text/template"
)

// DefaultPlaceholder represents the default placeholder value used for query generation.
const DefaultPlaceholder = "?"

// QueryOptions represents the options available for generating a query.
type QueryOptions struct {
	Placeholder string
	Ordered     bool
	Named       bool
	OmitEmpty   bool
	obj         any
}

// QueryOption represents a function that modifies the query options.
type QueryOption func(*QueryOptions)

// DefaultQueryOptions represents the default query options used for query generation.
var DefaultQueryOptions = []QueryOption{WithDefaultPlaceholder()}

// WithPlaceholder sets the placeholder value and whether the parameter should have
// a sequence number appended to it.
func WithPlaceholder(p string, o bool) QueryOption {
	return func(q *QueryOptions) {
		q.Placeholder = p
		q.Ordered = o
	}
}

func WithDefaultPlaceholder() QueryOption {
	return func(q *QueryOptions) {
		q.Placeholder = DefaultPlaceholder
		q.Ordered = false
	}
}

// WithNamedParameters sets the query to use named parameters.
func WithNamedParameters() QueryOption {
	return func(q *QueryOptions) {
		q.Named = true
	}
}

// WithoutEmptyValues indicates that columns with no value should be omitted from the query.
func WithoutEmptyValues(obj any) QueryOption {
	return func(q *QueryOptions) {
		q.OmitEmpty = true
		q.obj = obj
	}
}

// insertSQL is the raw template contents used to generate an insert query.
const insertSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  {{- $seq := 0 -}}
  INSERT INTO {{$table.Name}} (
  {{- range $idx, $col := $table.Columns -}}
    {{$col.Name}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end -}}
  ) VALUES (
  {{- range $idx, $col := $table.Columns -}}
    {{- $seq = add $seq 1 -}}
    {{param $col.Name $options $seq}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end -}}
  );`

// updateSQL is the raw template contents used to generate an update query.
const updateSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  {{- $seq := 0 -}}
  {{- $data := .Data -}}
  {{- $nonPrimaryKeys := .NonPrimaryKeys -}}
  UPDATE {{$table.Name}} AS {{$table.Alias}} SET {{- if true}} {{end}}
  {{- range $idx, $col := $nonPrimaryKeys -}}
    {{- if omit $data $col.Name -}} {{continue}} {{- end -}}
    {{- if ne $idx 0 -}} , {{end}}
    {{- $seq = add $seq 1 -}}
    {{$table.Alias}}.{{.Name}} = {{param $col.Name $options $seq}}
  {{- end }} WHERE 1=1
  {{- range $idx, $col := .PrimaryKeys -}}
    {{- $seq = add $seq 1 }} AND {{$table.Alias}}.{{.Name}} = {{param $col.Name $options $seq}}
  {{- end -}};`

// deleteSQL is the raw template contents used to generate a delete query.
const deleteSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  {{- $seq := 0 -}}
  DELETE FROM {{$table.Name}} WHERE 1=1
  {{- range $idx, $col := .PrimaryKeys -}}
    {{- $seq = add $seq 1 }} AND {{.Name}} = {{param $col.Name $options $seq}}
  {{- end -}};`

var (
	// funcs defines the custom functions leveraged within the query templates.
	funcs = template.FuncMap{
		"param": func(columnName string, options *QueryOptions, seq int) string {
			if options.Named {
				return ":" + columnName
			}

			p := options.Placeholder
			if options.Ordered {
				p += strconv.Itoa(seq)
			}

			return p
		},
		"omit": func(data EvaluationResult, columnName string) bool {
			for _, col := range data.Empties() {
				if col == columnName {
					return true
				}
			}

			return false
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
	}

	// insertTmpl is the parsed template used to generate an insert query.
	insertTmpl = template.Must(template.New("insertQuery").Funcs(funcs).Parse(insertSQL))

	// updateTmpl is the parsed template used to generate an update query.
	updateTmpl = template.Must(template.New("updateQuery").Funcs(funcs).Parse(updateSQL))

	// deleteTmpl is the parsed template used to generate a delete query.
	deleteTmpl = template.Must(template.New("deleteQuery").Funcs(funcs).Parse(deleteSQL))
)
