package morph

import (
	"errors"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

var (

	// ErrNotStruct is returned when the input being reflected is not a struct.
	// A struct is required to reflect the fields and methods to construct the
	// table metadata.
	ErrNotStruct = errors.New("morph: input must be a struct")
)

// Reflect observes the provided object and generates metadata from it
// using the provided options.
func Reflect(obj interface{}, options ...Option) (Table, error) {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return Table{}, ErrNotStruct
	}

	configuration := ReflectConfiguration{}
	WithInferredTableName(SnakeTableNameStrategy, true)(&configuration)
	WithInferredTableAlias(UppercaseTableAliasStrategy, DefaultTableAliasLength)(&configuration)
	WithInferredColumnNames(SnakeColumnNameStrategy)(&configuration)

	for _, option := range options {
		option(&configuration)
	}

	tableName := configuration.TableName
	if configuration.IsInferredTableName {
		if configuration.SnakeCaseTableName() {
			tName := strcase.ToSnake(t.Name())
			tableName = &tName
		}

		if configuration.CamelCaseTableName() {
			tName := strcase.ToCamel(t.Name())
			tableName = &tName
		}

		if configuration.IsTableNamePlural {
			tName := pluralize.NewClient().Plural(*tableName)
			tableName = &tName
		}
	}

	tableAlias := configuration.TableAlias
	if configuration.IsInferredTableAlias {
		if configuration.UppercaseTableAlias() {
			a := strings.ToUpper((*tableName)[:*configuration.TableAliasLength])
			tableAlias = &a
		}

		if configuration.LowercaseTableAlias() {
			a := strings.ToLower((*tableName)[:*configuration.TableAliasLength])
			tableAlias = &a
		}
	}

	val := reflect.ValueOf(obj)

	columns := []Column{}
	columns = append(columns, fields(t, val, configuration)...)
	columns = append(columns, methods(t, val, configuration)...)

	table := Table{}
	table.SetType(obj)
	table.SetName(*tableName)
	table.SetAlias(*tableAlias)
	table.AddColumns(columns...)

	return table, nil
}

func fields(t reflect.Type, v reflect.Value, c ReflectConfiguration) []Column {
	columns := []Column{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)
		fieldName := field.Name
		fieldType := field.Type.String()

		if !fieldVal.CanInterface() {
			continue
		}

		if fieldVal.Kind() == reflect.Ptr {
			fieldVal = fieldVal.Elem()
		}

		// ignore fields to other structs. if we want to retrieve
		// metadata for these fields, we should call Reflect on them directly.
		// NOTE: it would be nice to be able to recursively reflect struct
		// fields, but where things fall down is how we'd propogate the options
		// to each call in a way that makes sense.
		if fieldVal.Kind() == reflect.Struct {
			continue
		}

		var tagValue string
		if c.Tag != nil {
			tagValue = field.Tag.Get(*c.Tag)
		}

		if c.FieldExclusionPattern != nil && *c.FieldExclusionPattern != "" {
			if regexp.MustCompile(*c.FieldExclusionPattern).MatchString(fieldName) {
				continue
			}
		}

		if slices.Contains(c.FieldExclusions, field.Name) {
			continue
		}

		columnName := fieldName
		if tagValue != "" {
			columnName = tagValue
		} else if c.IsInferredColumnNames {
			if c.SnakeCaseColumnName() {
				columnName = strcase.ToSnake(columnName)
			}
			if c.CamelCaseColumnName() {
				columnName = strcase.ToCamel(columnName)
			}
		}

		var column Column
		column.SetField(fieldName)
		column.SetName(columnName)
		column.SetFieldType(fieldType)
		column.SetStrategy(FieldStrategyStructField)
		columns = append(columns, column)
	}
	return columns
}

func methods(t reflect.Type, v reflect.Value, c ReflectConfiguration) []Column {
	columns := []Column{}

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fieldName := method.Name

		if method.Type.NumOut() == 0 {
			continue
		}

		returnType := method.Type.Out(0)
		if returnType.Kind() == reflect.Ptr {
			returnType = returnType.Elem()
		}

		if returnType.Kind() == reflect.Struct && returnType != reflect.TypeOf(time.Time{}) {
			continue
		}
		fieldType := method.Type.Out(0).String()

		if c.MethodExclusionPattern != nil && *c.MethodExclusionPattern != "" {
			if regexp.MustCompile(*c.MethodExclusionPattern).MatchString(fieldName) {
				continue
			}
		}

		if slices.Contains(c.MethodExclusions, fieldName) {
			continue
		}

		columnName := fieldName
		if c.IsInferredColumnNames {
			if c.SnakeCaseColumnName() {
				columnName = strcase.ToSnake(columnName)
			}
			if c.CamelCaseColumnName() {
				columnName = strcase.ToCamel(columnName)
			}
		}

		var column Column
		column.SetField(fieldName)
		column.SetName(columnName)
		column.SetFieldType(fieldType)
		column.SetStrategy(FieldStrategyMethod)
		columns = append(columns, column)
	}
	return columns
}
