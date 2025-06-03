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

const insertSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  INSERT INTO {{$table.Name}} (
  {{- range $idx, $col := $table.Columns -}}
    {{$col.Name}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end -}}
  ) VALUES (
  {{- range $idx, $col := $table.Columns -}}
    {{param $col.Name $options (add $idx 1)}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end -}}
  );`

const updateSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  {{- $data := .Data -}}
  {{- $nonPrimaryKeys := .NonKeys -}}
  UPDATE {{$table.Name}} AS {{$table.Alias}} SET
  {{- range $idx, $col := $nonPrimaryKeys -}}
    {{- if not (omit $data $col.Name) -}}
      {{- if ne $idx 0 -}}, {{end}}
      {{$table.Alias}}.{{.Name}} = {{param $col.Name $options (add $idx 1)}}
    {{- end -}}
  {{- end }} WHERE 1=1
  {{- range $idx, $col := .Key -}}
    AND {{$table.Alias}}.{{.Name}} = {{param $col.Name $options (add $idx 1)}}
  {{- end -}};`

const deleteSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  DELETE FROM {{$table.Name}} WHERE 1=1
  {{- range $idx, $col := .Key -}}
    AND {{.Name}} = {{param $col.Name $options (add $idx 1)}}
  {{- end -}};`

const selectSQL = `
  {{- $table := .Table -}}
  {{- $options := .Options -}}
  SELECT
  {{- range $idx, $col := $table.Columns -}}
    {{$table.Alias}}.{{$col.Name}}{{if ne $idx (sub (len $table.Columns) 1)}}, {{end}}
  {{- end -}}
  FROM {{$table.Name}} AS {{$table.Alias}} WHERE 1=1
  {{- range $idx, $col := .Key -}}
    AND {{$table.Alias}}.{{.Name}} = {{param $col.Name $options (add $idx 1)}}
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

	// insertTmpl is the parsed template used to generate an INSERT query.
	insertTmpl = template.Must(template.New("insertQuery").Funcs(funcs).Parse(insertSQL))

	// updateTmpl is the parsed template used to generate an UPDATE query.
	updateTmpl = template.Must(template.New("updateQuery").Funcs(funcs).Parse(updateSQL))

	// deleteTmpl is the parsed template used to generate a DELETE query.
	deleteTmpl = template.Must(template.New("deleteQuery").Funcs(funcs).Parse(deleteSQL))

	// selectTmpl is the parsed template used to generate a SELECT query.
	selectTmpl = template.Must(template.New("selectQuery").Funcs(funcs).Parse(selectSQL))
)
