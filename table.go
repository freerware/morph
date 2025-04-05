package morph

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// Defines the various errors that can occur when interacting with tables.
var (
	// ErrMissingTypeName represents an error encountered when evaluation is attempted but the
	// table does not have a type name configured.
	ErrMissingTypeName = errors.New("morph: must have table type name to evaluate")

	// ErrMissingTableName represents an error encountered when evaluation is attempted but the
	// table does not have a table name configured.
	ErrMissingTableName = errors.New("morph: must have table name to evaluate")

	// ErrMissingTableAlias represents an error encountered when evaluation is attempted but the
	// table does not have a table alias configured.
	ErrMissingTableAlias = errors.New("morph: must have table alias to evaluate")

	// ErrMissingColumns represents an error encountered when evaluation is attempted but the
	// table does not have any columns configured.
	ErrMissingColumns = errors.New("morph: must have columns to evaluate")

	// ErrMismatchingTypeName represents an error encountered when evaluation is attempted but the
	// table type name does not match the type name of the object being evaluated.
	ErrMismatchingTypeName = errors.New("morph: must have matching type names to evaluate")

	// ErrMissingPrimaryKey represents an error encountered when a table does not have any primary key columns.
	ErrMissingPrimaryKey = errors.New("morph: table must have at least one primary key column")

	// ErrMissingNonPrimaryKey represents an error encountered when a table does not have any non-primary key columns.
	ErrMissingNonPrimaryKey = errors.New("morph: table must have at least one non-primary key column")
)

var (
	namedParamRegExp *regexp.Regexp = regexp.MustCompile(`:[a-zA-Z0-9_]+`)
)

// EvaluationResult represents the result of evaluating a table against an object.
type EvaluationResult map[string]any

// Empties retrieves all of the keys in the result that have nil values.
func (r EvaluationResult) Empties() []string {
	var empties []string
	for key, val := range r {
		if val == nil {
			empties = append(empties, key)
		}
	}
	return empties
}

// NonEmpties retrieves all of the keys in the result that have non-nil values.
func (r EvaluationResult) NonEmpties() []string {
	var nonEmpties []string
	for key, val := range r {
		if val != nil {
			nonEmpties = append(nonEmpties, key)
		}
	}
	return nonEmpties
}

// Table represents a mapping between an entity and a database table.
type Table struct {
	typeName       string
	name           string
	alias          string
	columnsByName  map[string]Column
	columnsByField map[string]Column
}

// SetType associates the entity type to the table.
func (t *Table) SetType(entity any) {
	t.SetTypeName(fmt.Sprintf("%T", entity))
}

// SetTypeName modifies the entity type name for the table.
func (t *Table) SetTypeName(typeName string) {
	t.typeName = strings.TrimSpace(typeName)
}

// TypeName retrieves the type name of the entity associated to the table.
func (t *Table) TypeName() string {
	return t.typeName
}

// Name retrieves the the table name.
func (t *Table) Name() string {
	return t.name
}

// SetName modifies the name of the table.
func (t *Table) SetName(name string) {
	t.name = strings.TrimSpace(name)
}

// Alias retrieves the alias for the table.
func (t *Table) Alias() string {
	return t.alias
}

// SetAlias modifies the alias of the table.
func (t *Table) SetAlias(alias string) {
	t.alias = strings.TrimSpace(alias)
}

// ColumnNames retrieves all of the column names for the table.
func (t *Table) ColumnNames() []string {
	var names []string
	for _, col := range t.Columns() {
		names = append(names, col.Name())
	}
	return names
}

// ColumnName retrieves the column name associated to the provide field name.
func (t *Table) ColumnName(field string) (string, error) {
	if t.columnsByField == nil {
		t.columnsByField = make(map[string]Column)
	}
	if column, ok := t.columnsByField[field]; ok {
		return column.Name(), nil
	}
	return "", fmt.Errorf("morph: no mapping for field %q", field)
}

// FieldName retrieves the field name associated to the provided column name.
func (t *Table) FieldName(name string) (string, error) {
	if t.columnsByName == nil {
		t.columnsByName = make(map[string]Column)
	}
	if column, ok := t.columnsByName[name]; ok {
		return column.Field(), nil
	}
	return "", fmt.Errorf("morph: no mapping for column %q", name)
}

// Columns retrieves all of the columns for the table.
func (t *Table) Columns() (columns []Column) {
	if t.columnsByName == nil {
		t.columnsByName = make(map[string]Column)
	}
	for _, c := range t.columnsByName {
		columns = append(columns, c)
	}
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].Name() < columns[j].Name()
	})
	return
}

// FindColumn retrieves the columns that matches the provided predicate.
func (t Table) FindColumns(p ColumnPredicate) []Column {
	cols := []Column{}
	for _, col := range t.Columns() {
		if p(col) {
			cols = append(cols, col)
		}
	}
	return cols
}

// AddColumn adds a column to the table.
func (t *Table) AddColumn(column Column) error {
	if t.columnsByName == nil {
		t.columnsByName = make(map[string]Column)
	}
	if _, ok := t.columnsByName[column.Name()]; ok {
		return fmt.Errorf(
			"morph: column with name %q already exists", column.Name())
	}
	if t.columnsByField == nil {
		t.columnsByField = make(map[string]Column)
	}
	if _, ok := t.columnsByField[column.Field()]; ok {
		return fmt.Errorf(
			"morph: column with field %q already exists", column.Field())
	}
	t.columnsByName[column.Name()],
		t.columnsByField[column.Field()] = column, column
	return nil
}

// AddColumns adds all of the provided columns to the table.
func (t *Table) AddColumns(columns ...Column) error {
	for _, column := range columns {
		if err := t.AddColumn(column); err != nil {
			return err
		}
	}
	return nil
}

// Evaluate applies the table to the provided object to produce a result
// containing the column names and their respective values. The result
// can then be subsequently used to execute queries.
func (t *Table) Evaluate(obj any) (EvaluationResult, error) {
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	// determine the type name for the pointer version of obj.
	ptrTypeName := fmt.Sprintf("%T", obj)
	if objType.Kind() != reflect.Ptr {
		ptrTypeName = fmt.Sprintf("%T", reflect.New(objVal.Type()).Interface())
	}

	// determine the type name for the value version of obj.
	valTypeName := fmt.Sprintf("%T", obj)
	if objType.Kind() == reflect.Ptr {
		valTypeName = fmt.Sprintf("%T", objVal.Elem().Interface())
	}

	// fail if the type name for the table doesn't match both the pointer and value type names.
	if t.typeName != ptrTypeName && t.typeName != valTypeName {
		return nil, ErrMismatchingTypeName
	}

	if err := t.validate(); err != nil {
		return nil, err
	}

	results := make(EvaluationResult)

	for fieldName, column := range t.columnsByField {
		if column.UsingStructFieldStrategy() {
			oVal := objVal
			if oVal.Kind() == reflect.Ptr {
				oVal = oVal.Elem()
			}
			val := oVal.FieldByName(fieldName)
			if !val.IsValid() {
				continue
			}

			if val.Kind() == reflect.Ptr && !val.IsZero() {
				if val.Elem().Kind() == reflect.Struct {
					continue
				}
				val = val.Elem()
			}

			if val.Kind() == reflect.Ptr && (val.IsZero() || val.IsNil()) {
				results[column.Name()] = nil
				continue
			}

			results[column.Name()] = val.Interface()
		}

		if column.UsingMethodStrategy() {
			oType := objType
			oVal := objVal
			if oVal.Kind() != reflect.Ptr {
				ptr := reflect.New(oVal.Type())
				ptr.Elem().Set(oVal)
				oVal = ptr
			}
			if oType.Kind() != reflect.Ptr {
				oType = reflect.PointerTo(oType)
			}
			method, ok := oType.MethodByName(fieldName)
			if !ok {
				continue
			}

			valResults := method.Func.Call([]reflect.Value{oVal})
			if len(valResults) == 0 || !valResults[0].IsValid() {
				continue
			}

			if valResults[0].Kind() == reflect.Ptr && (valResults[0].IsZero() || valResults[0].IsNil()) {
				results[column.Name()] = nil
				continue
			}
			results[column.Name()] = valResults[0].Interface()
		}
	}

	return results, nil
}

// MustEvaluate performs the same operation as Evaluate but panics if an error occurs.
func (t *Table) MustEvaluate(obj any) EvaluationResult {
	results, err := t.Evaluate(obj)
	if err != nil {
		panic(err)
	}
	return results
}

// validate ensures that the table is properly configured.
func (t *Table) validate() error {
	if len(t.typeName) == 0 {
		return ErrMissingTypeName
	}

	if len(t.name) == 0 {
		return ErrMissingTableName
	}

	if len(t.alias) == 0 {
		return ErrMissingTableAlias
	}

	if len(t.columnsByName) == 0 {
		return ErrMissingColumns
	}

	if len(t.FindColumns(func(c Column) bool { return c.PrimaryKey() })) == 0 {
		return ErrMissingPrimaryKey
	}

	if len(t.FindColumns(func(c Column) bool { return !c.PrimaryKey() })) == 0 {
		return ErrMissingNonPrimaryKey
	}

	return nil
}

// query generates a query for the table using the provided template and options.
func (t *Table) query(tmpl *template.Template, options ...QueryOption) (string, error) {
	if err := t.validate(); err != nil {
		return "", err
	}

	qo := &QueryOptions{}
	opts := append(DefaultQueryOptions, options...)
	for _, opt := range opts {
		opt(qo)
	}

	data := struct {
		Table          *Table
		PrimaryKeys    []Column
		NonPrimaryKeys []Column
		Options        *QueryOptions
		Data           EvaluationResult
	}{
		Table:          t,
		Options:        qo,
		PrimaryKeys:    t.FindColumns(func(c Column) bool { return c.PrimaryKey() }),
		NonPrimaryKeys: t.FindColumns(func(c Column) bool { return !c.PrimaryKey() }),
	}

	if qo.OmitEmpty && qo.obj != nil {
		var err error
		if data.Data, err = t.Evaluate(qo.obj); err != nil {
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

func (t *Table) queryWithArgs(namedQuery string, obj any, options ...QueryOption) (string, []any, error) {
	qo := &QueryOptions{}
	opts := append(DefaultQueryOptions, options...)
	for _, opt := range opts {
		opt(qo)
	}

	result, err := t.Evaluate(obj)
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

// InsertQuery generates an insert query for the table.
func (t *Table) InsertQuery(options ...QueryOption) (string, error) {
	return t.query(insertTmpl, options...)
}

// InsertQueryWithArgs generates an insert query for the table along with arguments
// derived from the provided object.
func (t *Table) InsertQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	query, err := t.InsertQuery(append(options, WithNamedParameters())...)
	if err != nil {
		return "", nil, err
	}

	return t.queryWithArgs(query, obj, options...)
}

// MustInsertQuery performs the same operation as InsertQuery but panics if an error occurs.
func (t *Table) MustInsertQuery(options ...QueryOption) string {
	return Must(t.InsertQuery(options...))
}

// UpdateQuery generates an update query for the table.
func (t *Table) UpdateQuery(options ...QueryOption) (string, error) {
	return t.query(updateTmpl, options...)
}

// UpdateQueryWithArgs generates an update query for the table along with arguments
// derived from the provided object.
func (t *Table) UpdateQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	query, err := t.UpdateQuery(append(options, WithNamedParameters())...)
	if err != nil {
		return "", nil, err
	}

	return t.queryWithArgs(query, obj, options...)
}

// MustUpdateQuery performs the same operation as UpdateQuery but panics if an error occurs.
func (t *Table) MustUpdateQuery(options ...QueryOption) string {
	return Must(t.UpdateQuery(options...))
}

// DeleteQuery generates a delete query for the table.
func (t *Table) DeleteQuery(options ...QueryOption) (string, error) {
	return t.query(deleteTmpl, options...)
}

// DeleteQueryWithArgs generates a delete query for the table along with arguments
// derived from the provided object.
func (t *Table) DeleteQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	query, err := t.DeleteQuery(append(options, WithNamedParameters())...)
	if err != nil {
		return "", nil, err
	}

	return t.queryWithArgs(query, obj, options...)
}

// MustDeleteQuery performs the same operation as DeleteQuery but panics if an error occurs.
func (t *Table) MustDeleteQuery(options ...QueryOption) string {
	return Must(t.DeleteQuery(options...))
}

// SelectQuery generates a SELECT query for the table.
func (t *Table) SelectQuery(options ...QueryOption) (string, error) {
	return t.query(selectTmpl, options...)
}

// SelectQueryWithArgs generates a SELECT query for the table along with arguments
// derived from the provided object.
func (t *Table) SelectQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	query, err := t.SelectQuery(append(options, WithNamedParameters())...)
	if err != nil {
		return "", nil, err
	}

	return t.queryWithArgs(query, obj, options...)
}

// MustSelectQuery performs the same operation as SelectQuery but panics if an error occurs.
func (t *Table) MustSelectQuery(options ...QueryOption) string {
	return Must(t.SelectQuery(options...))
}
