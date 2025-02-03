package morph

import (
	"errors"
	"reflect"
	"regexp"
	"slices"
	"strings"
)

var (
	ErrNotStruct = errors.New("morph: input must be a struct")
)

func Reflect(obj interface{}, options ...Option) (map[string]Table, error) {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	configuration := ReflectConfiguration{}
	WithInferredTableName(SnakeTableNameStrategy, true)(&configuration)
	WithInferredTableAlias(UppercaseTableAliasStrategy, 1)(&configuration)

	for _, option := range options {
		option(&configuration)
	}

	if configuration.TableName == nil {
		if configuration.TableNameStrategy == nil {
			return nil, errors.New("morph: table name strategy must be provided")
		}

		if *configuration.TableNameStrategy == SnakeTableNameStrategy {
			tableName := strings.ToLower(t.Name())
			configuration.TableName = &tableName
		}

		if *configuration.TableNameStrategy == CamelTableNameStrategy {
			tableName := strings.ToLower(t.Name())
			configuration.TableName = &tableName
		}
	}

	tableName := *configuration.TableName

	val := reflect.ValueOf(obj)

	columns := []Column{}
	columns = append(columns, fields(t, val, configuration)...)
	columns = append(columns, methods(t, val, configuration)...)

	table := Table{}
	table.SetTypeName(t.Name())
	table.SetName(tableName)
	table.SetAlias(strings.ToUpper(tableName[:1]))
	table.AddColumns(columns...)

	return map[string]Table{table.Name(): table}, nil
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

		if c.Tag != nil {
			fieldName = field.Tag.Get(*c.Tag)
			if fieldName == "" {
				fieldName = field.Name
			}
		}

		if c.FieldExclusionPattern != nil && *c.FieldExclusionPattern != "" {
			if regexp.MustCompile(*c.FieldExclusionPattern).MatchString(fieldName) {
				continue
			}
		}

		if slices.Contains(c.FieldExclusions, field.Name) {
			continue
		}

		var column Column
		column.SetField(fieldName)
		column.SetName(strings.ToLower(fieldName))
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
		columnName := strings.ToLower(method.Name)

		if method.Type.NumOut() == 0 {
			continue
		}

		returnType := method.Type.Out(0)
		if returnType.Kind() == reflect.Ptr {
			returnType = returnType.Elem()
		}

		if returnType.Kind() == reflect.Struct {
			continue
		}
		fieldType := method.Type.Out(0).String()

		if c.MethodExclusionPattern != nil && *c.FieldExclusionPattern != "" {
			if regexp.MustCompile(*c.MethodExclusionPattern).MatchString(fieldName) {
				continue
			}
		}

		if slices.Contains(c.MethodExclusions, fieldName) {
			continue
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
