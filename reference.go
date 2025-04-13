package morph

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strings"
	"text/template"
)

// Reference represents a reference between two tables via a foreign key.
type Reference struct {
	parent     Table
	child      Table
	foreignKey []Column
}

// Parent returns the parent table of the reference.
func (r *Reference) Parent() Table {
	return r.parent
}

// Child returns the child table of the reference.
func (r *Reference) Child() Table {
	return r.child
}

// ForeignKey returns the foreign key columns of the reference.
func (r *Reference) ForeignKey() []Column {
	key := make([]Column, len(r.foreignKey))
	copy(key, r.foreignKey)
	return key
}

// equal checks if two references are equal.
func (r Reference) equals(other Reference) bool {
	sameParent := r.Parent().Equals(other.Parent())
	sameChild := r.Child().Equals(other.Child())
	sameForeignKey := slices.EqualFunc(r.ForeignKey(), other.ForeignKey(), func(a, b Column) bool {
		return a.equals(b)
	})

	return sameParent && sameChild && sameForeignKey
}

// query generates a query for the table using the provided template and options.
func (r *Reference) query(tmpl *template.Template, options ...QueryOption) (string, error) {
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
		Table:   &r.child,
		Options: qo,
		Key:     r.foreignKey,
		NonKeys: r.child.FindColumns(func(c Column) bool { return !c.PrimaryKey() }),
	}

	if qo.OmitEmpty && qo.obj != nil {
		var err error
		if data.Data, err = r.child.Evaluate(qo.obj); err != nil {
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

func (r *Reference) queryWithArgs(namedQuery string, obj any, options ...QueryOption) (string, []any, error) {
	qo := &QueryOptions{}
	opts := append(DefaultQueryOptions, options...)
	for _, opt := range opts {
		opt(qo)
	}

	result, err := r.child.Evaluate(obj)
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

// SelectQuery generates a SELECT query for the reference.
func (r *Reference) SelectQuery(options ...QueryOption) (string, error) {
	return r.query(selectTmpl, options...)
}

// SelectQueryWithArgs generates a SELECT query with arguments for the reference.
func (r *Reference) SelectQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	query, err := r.SelectQuery(opts...)
	if err != nil {
		return "", nil, err
	}

	return r.queryWithArgs(query, obj, opts...)
}

// DeleteQuery generates a DELETE query for the reference.
func (r *Reference) DeleteQuery(options ...QueryOption) (string, error) {
	return r.query(deleteTmpl, options...)
}

// DeleteQueryWithArgs generates a DELETE query with arguments for the reference.
func (r *Reference) DeleteQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	query, err := r.DeleteQuery(opts...)
	if err != nil {
		return "", nil, err
	}

	return r.queryWithArgs(query, obj, opts...)
}
