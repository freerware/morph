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
	ErrNotStruct = errors.New("morph: input must be a struct or pointer to a struct")
)

// Reflect observes the provided object and generates metadata from it
// using the provided options.
func Reflect(obj any, options ...ReflectOption) (Table, error) {
	t := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return Table{}, ErrNotStruct
	}
	pt := reflect.PointerTo(t)

	configuration := ReflectConfiguration{}
	opts := append(DefaultReflectOptions, options...)
	for _, opt := range opts {
		opt(&configuration)
	}

	tableName := configuration.TableName
	if configuration.IsInferredTableName {
		if configuration.SnakeCaseTableName() {
			tName := strcase.ToSnake(t.Name())
			tableName = &tName
		}

		if configuration.ScreamingSnakeCaseTableName() {
			tName := strings.ToUpper(strcase.ToSnake(t.Name()))
			tableName = &tName
		}

		if configuration.CamelCaseTableName() {
			tName := strcase.ToCamel(t.Name())
			tableName = &tName
		}

		if configuration.PluralTableName() {
			tName := pluralize.NewClient().Plural(*tableName)
			tableName = &tName
		}
	}

	tableAlias := configuration.TableAlias
	if configuration.IsInferredTableAlias && configuration.HasTableAliasLength() {
		if configuration.UppercaseTableAlias() {
			a := strings.ToUpper((*tableName)[:*configuration.TableAliasLength])
			tableAlias = &a
		}

		if configuration.LowercaseTableAlias() {
			a := strings.ToLower((*tableName)[:*configuration.TableAliasLength])
			tableAlias = &a
		}
	}

	columns := []Column{}
	columns = append(columns, fields(t, val, configuration)...)
	columns = append(columns, methods(pt, configuration)...)

	table := Table{}
	table.SetType(obj)
	table.SetName(*tableName)
	table.SetAlias(*tableAlias)
	if err := table.AddColumns(columns...); err != nil {
		return Table{}, err
	}

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
		if c.HasTag() {
			tagValue = strings.TrimSpace(field.Tag.Get(*c.Tag))
			if tagValue == "-" {
				continue
			}
		}

		if c.HasFieldExclusionPattern() {
			if regexp.MustCompile(*c.FieldExclusionPattern).MatchString(fieldName) {
				continue
			}
		}

		if c.HasFieldExclusions() && slices.Contains(c.FieldExclusions, field.Name) {
			continue
		}

		columnName := fieldName
		if tagValue != "" {
			columnName = tagValue
		} else if c.IsInferredColumnNames {
			if c.SnakeCaseColumnName() {
				columnName = strcase.ToSnake(columnName)
			}

			if c.ScreamingSnakeCaseColumnName() {
				columnName = strings.ToUpper(strcase.ToSnake(columnName))
			}

			if c.CamelCaseColumnName() {
				columnName = strcase.ToCamel(columnName)
			}

			if c.UppercaseColumnName() {
				columnName = strings.ToUpper(columnName)
			}

			if c.LowercaseColumnName() {
				columnName = strings.ToLower(columnName)
			}
		}

		var column Column
		column.SetField(fieldName)
		if c.HasColumnNameMappings() {
			if cName, ok := c.ColumnNameMappings[fieldName]; ok {
				columnName = cName
			}
		}
		column.SetName(columnName)
		if slices.Contains(c.PrimaryKeyColumns, columnName) {
			column.SetPrimaryKey(true)
		}
		column.SetFieldType(fieldType)
		column.SetStrategy(FieldStrategyStructField)
		columns = append(columns, column)
	}
	return columns
}

func methods(t reflect.Type, c ReflectConfiguration) []Column {
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

		if c.HasMethodExclusionPattern() {
			if regexp.MustCompile(*c.MethodExclusionPattern).MatchString(fieldName) {
				continue
			}
		}

		if c.HasMethodExclusions() && slices.Contains(c.MethodExclusions, fieldName) {
			continue
		}

		columnName := fieldName
		if c.IsInferredColumnNames {
			if c.SnakeCaseColumnName() {
				columnName = strcase.ToSnake(columnName)
			}

			if c.ScreamingSnakeCaseColumnName() {
				columnName = strings.ToUpper(strcase.ToSnake(columnName))
			}

			if c.CamelCaseColumnName() {
				columnName = strcase.ToCamel(columnName)
			}

			if c.UppercaseColumnName() {
				columnName = strings.ToUpper(columnName)
			}

			if c.LowercaseColumnName() {
				columnName = strings.ToLower(columnName)
			}
		}

		var column Column
		column.SetField(fieldName)
		if c.HasColumnNameMappings() {
			if cName, ok := c.ColumnNameMappings[fieldName]; ok {
				columnName = cName
			}
		}
		column.SetName(columnName)
		if slices.Contains(c.PrimaryKeyColumns, columnName) {
			column.SetPrimaryKey(true)
		}
		column.SetFieldType(fieldType)
		column.SetStrategy(FieldStrategyMethod)
		columns = append(columns, column)
	}
	return columns
}
