package morph

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"
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

	// ErrMissingForeignKeyColumns represents an error encountered when establishing a reference
	// between two tables but the foreign key columns are not present in the child table.
	ErrMissingForeignKeyColumns = errors.New("morph: specified foreign key columns are missing from table")

	// ErrMissingReference represents an error encountered when a reference is not found
	// based on the provided criteria.
	ErrMissingReference = errors.New("morph: no matching reference found")
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
	references     []Reference
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
func (t Table) TypeName() string {
	return t.typeName
}

// Name retrieves the the table name.
func (t Table) Name() string {
	return t.name
}

// SetName modifies the name of the table.
func (t *Table) SetName(name string) {
	t.name = strings.TrimSpace(name)
}

// Alias retrieves the alias for the table.
func (t Table) Alias() string {
	return t.alias
}

// ReferencesTo retrieves all of the tables that this table references.
func (t Table) ReferencesTo() []Table {
	refs := t.FindReferences(func(r Reference) bool {
		return r.child.Equals(t)
	})
	tables := []Table{}
	for _, ref := range refs {
		tables = append(tables, ref.parent)
	}
	return tables
}

// ReferenceTo retrieves the reference to the provided table.
func (t Table) ReferenceTo(table Table) (Reference, error) {
	ref, ok := t.FindReference(func(r Reference) bool {
		return r.child.Equals(t) && r.parent.Equals(table)
	})

	if !ok {
		return Reference{}, ErrMissingReference
	}

	return ref, nil
}

// HasReferenceTo checks if this table has a reference to the provided table.
func (t Table) HasReferenceTo(table Table) bool {
	_, ok := t.FindReference(func(r Reference) bool {
		return r.child.Equals(t) && r.parent.Equals(table)
	})
	return ok
}

// IsReferenced checks if this table is referenced by any other tables.
func (t Table) IsReferenced() bool {
	_, ok := t.FindReference(func(r Reference) bool {
		return r.parent.Equals(t)
	})
	return ok
}

// IsReferencedBy checks if this table is referenced by the provided table.
func (t Table) IsReferencedBy(table Table) bool {
	_, ok := t.FindReference(func(r Reference) bool {
		return r.child.Equals(table) && r.parent.Equals(t)
	})
	return ok
}

// ReferencedBy retrieves all of the tables that reference this table.
func (t Table) ReferencedBy() []Table {
	refs := t.FindReferences(func(r Reference) bool {
		return r.parent.Equals(t)
	})
	tables := []Table{}
	for _, ref := range refs {
		tables = append(tables, ref.child)
	}
	return tables
}

// FindReferences retrieves all of the references that match the provided predicate.
func (t Table) FindReferences(p func(Reference) bool) []Reference {
	refs := []Reference{}
	for _, ref := range t.references {
		if p(ref) {
			refs = append(refs, ref)
		}
	}
	return refs
}

// FindReference retrieves the first reference that matches the provided predicate,
// returning a boolean indicating if a match was found.
func (t Table) FindReference(p func(Reference) bool) (Reference, bool) {
	refs := t.FindReferences(p)
	if len(refs) > 0 {
		return refs[0], true
	}
	return Reference{}, false
}

func (t *Table) setReferences(refs []Reference) {
	if refs == nil {
		t.references = []Reference{}
		return
	}

	if t.references == nil {
		t.references = []Reference{}
	}

	r := make([]Reference, len(refs))
	copy(r, refs)
	t.references = r
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

// PrimaryKeyColumns retrieves all of the primary key columns for the table.
func (t *Table) PrimaryKeyColumns() []Column {
	return t.FindColumns(func(c Column) bool {
		return c.PrimaryKey()
	})
}

// NonPrimaryKeyColumns retrieves all of the non-primary key columns for the table.
func (t *Table) NonPrimaryKeyColumns() []Column {
	return t.FindColumns(func(c Column) bool {
		return !c.PrimaryKey()
	})
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

// InsertQuery generates an INSERT query for the table.
func (t *Table) InsertQuery(options ...QueryOption) (string, error) {
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.query(insertTmpl, options...)
}

// InsertQueryWithArgs generates an INSERT query for the table along with arguments
// derived from the provided object.
func (t *Table) InsertQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.InsertQueryWithArgs(obj, opts...)
}

// MustInsertQuery performs the same operation as InsertQuery but panics if an error occurs.
func (t *Table) MustInsertQuery(options ...QueryOption) string {
	return Must(t.InsertQuery(options...))
}

// UpdateQuery generates an UPDATE query for the table.
func (t *Table) UpdateQuery(options ...QueryOption) (string, error) {
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.query(updateTmpl, options...)
}

// UpdateQueryWithArgs generates an UPDATE query for the table along with arguments
// derived from the provided object.
func (t *Table) UpdateQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.UpdateQueryWithArgs(obj, opts...)
}

// MustUpdateQuery performs the same operation as UpdateQuery but panics if an error occurs.
func (t *Table) MustUpdateQuery(options ...QueryOption) string {
	return Must(t.UpdateQuery(options...))
}

// DeleteQuery generates a DELETE query for the table.
func (t *Table) DeleteQuery(options ...QueryOption) (string, error) {
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.query(deleteTmpl, options...)
}

// DeleteQueryWithArgs generates a DELETE query for the table along with arguments
// derived from the provided object.
func (t *Table) DeleteQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.DeleteQueryWithArgs(obj, opts...)
}

// MustDeleteQuery performs the same operation as DeleteQuery but panics if an error occurs.
func (t *Table) MustDeleteQuery(options ...QueryOption) string {
	return Must(t.DeleteQuery(options...))
}

// SelectQuery generates a SELECT query for the table.
func (t *Table) SelectQuery(options ...QueryOption) (string, error) {
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.query(selectTmpl, options...)
}

// SelectQueryWithArgs generates a SELECT query for the table along with arguments
// derived from the provided object.
func (t *Table) SelectQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	generator := newQueryGenerator(t, t.PrimaryKeyColumns(), t.NonPrimaryKeyColumns())
	return generator.SelectQueryWithArgs(obj, opts...)
}

// MustSelectQuery performs the same operation as SelectQuery but panics if an error occurs.
func (t *Table) MustSelectQuery(options ...QueryOption) string {
	return Must(t.SelectQuery(options...))
}

// References creates a reference between two tables, where the provided table
// is treated as the parent and the current table is treated as the child.
func (t *Table) References(table *Table, key ...Column) (Reference, error) {
	if err := t.validate(); err != nil {
		return Reference{}, err
	}

	if err := table.validate(); err != nil {
		return Reference{}, err
	}

	if len(t.FindColumns(func(c Column) bool { return slices.Contains(key, c) })) == 0 {
		return Reference{}, ErrMissingForeignKeyColumns
	}

	ref := Reference{parent: *table, child: *t, foreignKey: key}

	if !slices.ContainsFunc(t.references, func(r Reference) bool { return ref.equals(r) }) {
		t.setReferences(append(t.references, ref))
	}

	if !slices.ContainsFunc(table.references, func(r Reference) bool { return ref.equals(r) }) {
		table.setReferences(append(table.references, ref))
	}

	return Reference{parent: *table, child: *t, foreignKey: key}, nil
}

// equal checks if two tables are equal.
func (t Table) Equals(other Table) bool {
	sameTypeName := t.TypeName() == other.TypeName()
	sameName := t.Name() == other.Name()
	sameColumns := slices.EqualFunc(t.Columns(), other.Columns(), func(a, b Column) bool {
		return a.equals(b)
	})

	return sameTypeName && sameName && sameColumns
}
