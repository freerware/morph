package morph

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"text/template"
)

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
)

// Table represents a mapping between an entity and a database table.
type Table struct {
	typeName       string
	name           string
	alias          string
	columnsByName  map[string]Column
	columnsByField map[string]Column
}

// SetType associates the entity type to the table.
func (t *Table) SetType(entity interface{}) {
	t.SetTypeName(fmt.Sprintf("%T", entity))
}

// SetTypeName modifies the entity type name for the table.
func (t *Table) SetTypeName(typeName string) {
	t.typeName = strings.Trim(typeName, " ")
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
	t.name = strings.Trim(name, " ")
}

// Alias retrieves the alias for the table.
func (t *Table) Alias() string {
	return t.alias
}

// SetAlias modifies the alias of the table.
func (t *Table) SetAlias(alias string) {
	t.alias = strings.Trim(alias, " ")
}

// ColumnNames retrieves all of the column names for the table.
func (t *Table) ColumnNames() []string {
	var columnNames []string
	for name := range t.columnsByName {
		columnNames = append(columnNames, name)
	}
	return columnNames
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
func (t *Table) Columns() []Column {
	var columns []Column
	for _, c := range t.columnsByName {
		columns = append(columns, c)
	}
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].Name() < columns[j].Name()
	})
	return columns
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

func (t *Table) Evaluate(obj interface{}) (map[string]interface{}, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}

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

	results := make(map[string]interface{})

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
			results[column.Name()] = valResults[0].Interface()
		}
	}

	return results, nil
}

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

	return nil
}

func (t *Table) query(tmpl *template.Template, options ...QueryOption) (string, error) {
	if err := t.validate(); err != nil {
		return "", err
	}

	qo := &QueryOptions{}
	WithPlaceholder("?", false)(qo)

	for _, option := range options {
		option(qo)
	}

	data := struct {
		Table   *Table
		Options *QueryOptions
	}{Table: t, Options: qo}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (t *Table) InsertQuery(options ...QueryOption) (string, error) {
	return t.query(insertTmpl, options...)
}

func (t *Table) UpdateQuery(options ...QueryOption) (string, error) {
	return t.query(updateTmpl, options...)
}

func (t *Table) DeleteQuery(options ...QueryOption) (string, error) {
	return t.query(deleteTmpl, options...)
}
