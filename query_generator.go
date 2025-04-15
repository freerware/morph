package morph

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"
)

type queryGenerator struct {
	table         *Table
	keyColumns    []Column
	nonKeyColumns []Column
}

// newQueryGenerator creates a new query generator for the given table.
func newQueryGenerator(table *Table, keyColumns, nonKeyColumns []Column) queryGenerator {
	return queryGenerator{
		table:         table,
		keyColumns:    keyColumns,
		nonKeyColumns: nonKeyColumns,
	}
}

// query generates a query for the table using the provided template and options.
func (generator *queryGenerator) query(tmpl *template.Template, options ...QueryOption) (string, error) {
	if err := generator.table.validate(); err != nil {
		return "", err
	}

	qo := &QueryOptions{}
	opts := append(DefaultQueryOptions, options...)
	for _, opt := range opts {
		opt(qo)
	}

	data := struct {
		Table   *Table
		Key     []Column
		NonKeys []Column
		Options *QueryOptions
		Data    EvaluationResult
	}{
		Table:   generator.table,
		Options: qo,
		Key:     generator.keyColumns,
		NonKeys: generator.nonKeyColumns,
	}

	if qo.OmitEmpty && qo.obj != nil {
		var err error
		if data.Data, err = generator.table.Evaluate(qo.obj); err != nil {
			return "", err
		}
	}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (generator *queryGenerator) queryWithArgs(namedQuery string, obj any, options ...QueryOption) (string, []any, error) {
	qo := &QueryOptions{}
	opts := append(DefaultQueryOptions, options...)
	for _, opt := range opts {
		opt(qo)
	}

	result, err := generator.table.Evaluate(obj)
	if err != nil {
		return "", nil, err
	}

	args := []any{}
	missing := []string{}

	count := 0
	query := namedParamRegExp.ReplaceAllStringFunc(namedQuery, func(match string) string {
		name := match[1:]

		if arg, ok := result[name]; ok {
			args = append(args, arg)
			if qo.Ordered {
				count += 1
				return qo.Placeholder + fmt.Sprintf("%d", count)
			}
			return qo.Placeholder
		}

		missing = append(missing, name)
		return match
	})

	if len(missing) > 0 {
		return "", nil, errors.New("morph: missing values for named parameters: " + strings.Join(missing, ", "))
	}

	return query, args, nil
}

// InsertQuery generates an INSERT query for the table.
func (generator queryGenerator) InsertQuery(options ...QueryOption) (string, error) {
	return generator.query(insertTmpl, options...)
}

// InsertQueryWithArgs generates an INSERT query for the table along with arguments
// derived from the provided object.
func (generator queryGenerator) InsertQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	query, err := generator.InsertQuery(opts...)
	if err != nil {
		return "", nil, err
	}

	return generator.queryWithArgs(query, obj, opts...)
}

// SelectQuery generates a SELECT query for the table.
func (generator queryGenerator) SelectQuery(options ...QueryOption) (string, error) {
	return generator.query(selectTmpl, options...)
}

// SelectQueryWithArgs generates a SELECT query for the table along with arguments
// derived from the provided object.
func (generator *queryGenerator) SelectQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	query, err := generator.SelectQuery(options...)
	if err != nil {
		return "", nil, err
	}

	return generator.queryWithArgs(query, obj, options...)
}

// UpdateQuery generates an UPDATE query for the table.
func (generator queryGenerator) UpdateQuery(options ...QueryOption) (string, error) {
	return generator.query(updateTmpl, options...)
}

// UpdateQueryWithArgs generates an UPDATE query for the table along with arguments
// derived from the provided object.
func (generator *queryGenerator) UpdateQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	query, err := generator.UpdateQuery(options...)
	if err != nil {
		return "", nil, err
	}

	return generator.queryWithArgs(query, obj, options...)
}

// DeleteQueryWithArgs generates a DELETE query for the table along with arguments
// derived from the provided object.
func (generator *queryGenerator) DeleteQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	query, err := generator.DeleteQuery(options...)
	if err != nil {
		return "", nil, err
	}

	return generator.queryWithArgs(query, obj, options...)
}

// DeleteQuery generates a DELETE query for the table.
func (generator queryGenerator) DeleteQuery(options ...QueryOption) (string, error) {
	return generator.query(deleteTmpl, options...)
}
