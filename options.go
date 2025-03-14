package morph

import "strings"

// CaseStrategy is an enumeration of the available case strategies.
type CaseStrategy string

const (
	// SnakeCaseStrategy is the snake case strategy.
	SnakeCaseStrategy CaseStrategy = "snake"

	// ScreamingSnakeCaseStrategy is the screaming (uppercase) snake case strategy.
	ScreamingSnakeCaseStrategy CaseStrategy = "screaming_snake"

	// CamelCaseStrategy is the camel case strategy.
	CamelCaseStrategy CaseStrategy = "camel"

	// LowerCaseStrategy is the lowercase strategy.
	LowerCaseStrategy CaseStrategy = "lower"

	// UpperCaseStrategy is the uppercase strategy.
	UpperCaseStrategy CaseStrategy = "upper"

	// DefaultTableAliasLength is the default table alias length.
	DefaultTableAliasLength = 1

	// DefaultTableNameStrategy is the default table name strategy.
	DefaultTableNameStrategy = SnakeCaseStrategy

	// DefaultTableAliasStrategy is the default table alias strategy.
	DefaultTableAliasStrategy = UpperCaseStrategy

	// DefaultColumnNameStrategy is the default column name strategy.
	DefaultColumnNameStrategy = SnakeCaseStrategy
)

// ReflectOption is a function that configures the reflection configuration.
type ReflectOption func(*ReflectConfiguration)

var (
	// DefaultReflectOptions represents the default reflection options used for reflection.
	DefaultReflectOptions = []ReflectOption{
		WithInferredTableName(DefaultTableNameStrategy, true),
		WithInferredTableAlias(DefaultTableAliasStrategy, DefaultTableAliasLength),
		WithInferredColumnNames(DefaultColumnNameStrategy),
		WithPrimaryKeyColumn("id"),
	}
)

// ReflectConfiguration is the configuration leveraged when constructing
// tables via reflection.
type ReflectConfiguration struct {
	TableName              *string
	TableAlias             *string
	IsInferredTableName    bool
	IsInferredTableAlias   bool
	IsInferredColumnNames  bool
	TableNameStrategy      *CaseStrategy
	TableAliasStrategy     *CaseStrategy
	ColumnNameStrategy     *CaseStrategy
	Tag                    *string
	MethodExclusions       []string
	FieldExclusions        []string
	MethodExclusionPattern *string
	FieldExclusionPattern  *string
	IsTableNamePlural      bool
	TableAliasLength       *int
	PrimaryKeyColumns      []string
}

// HasTableName indicates if the table name is set.
func (c *ReflectConfiguration) HasTableName() bool {
	return c.TableName != nil && strings.TrimSpace(*c.TableName) != ""
}

// HasTableNameStrategy indicates if the table name strategy is set.
func (c *ReflectConfiguration) HasTableNameStrategy() bool {
	return c.TableNameStrategy != nil
}

// HasTableAliasStrategy indicates if the table alias strategy is set.
func (c *ReflectConfiguration) HasTableAliasStrategy() bool {
	return c.TableAliasStrategy != nil
}

// HasColumnNameStrategy indicates if the column name strategy is set.
func (c *ReflectConfiguration) HasColumnNameStrategy() bool {
	return c.ColumnNameStrategy != nil
}

// SnakeCaseTableName indicates if the table name strategy is snake case.
func (c *ReflectConfiguration) SnakeCaseTableName() bool {
	return c.HasTableNameStrategy() && *c.TableNameStrategy == SnakeCaseStrategy
}

// ScreamingSnakeCaseTableName indicates if the table name strategy is screaming snake case.
func (c *ReflectConfiguration) ScreamingSnakeCaseTableName() bool {
	return c.HasTableNameStrategy() && *c.TableNameStrategy == ScreamingSnakeCaseStrategy
}

// CamelCaseTableName indicates if the table name strategy is camel case.
func (c *ReflectConfiguration) CamelCaseTableName() bool {
	return c.HasTableNameStrategy() && *c.TableNameStrategy == CamelCaseStrategy
}

// SnakeCaseColumnName indicates if the column name strategy is snake case.
func (c *ReflectConfiguration) SnakeCaseColumnName() bool {
	return c.HasColumnNameStrategy() && *c.ColumnNameStrategy == SnakeCaseStrategy
}

// ScreamingSnakeCaseColumnName indicates if the column name strategy is scream snake case.
func (c *ReflectConfiguration) ScreamingSnakeCaseColumnName() bool {
	return c.HasColumnNameStrategy() && *c.ColumnNameStrategy == ScreamingSnakeCaseStrategy
}

// CamelCaseColumnName indicates if the column name strategy is camel case.
func (c *ReflectConfiguration) CamelCaseColumnName() bool {
	return c.HasColumnNameStrategy() && *c.ColumnNameStrategy == CamelCaseStrategy
}

// UppercaseColumnName indicates if the column name strategy is uppercase.
func (c *ReflectConfiguration) UppercaseColumnName() bool {
	return c.HasColumnNameStrategy() && *c.ColumnNameStrategy == UpperCaseStrategy
}

// LowercaseColumnName indicates if the column name strategy is uppercase.
func (c *ReflectConfiguration) LowercaseColumnName() bool {
	return c.HasColumnNameStrategy() && *c.ColumnNameStrategy == LowerCaseStrategy
}

// LowercaseTableAlias indicates if the table alias strategy is lowercase.
func (c *ReflectConfiguration) LowercaseTableAlias() bool {
	return c.HasTableAliasStrategy() && *c.TableAliasStrategy == LowerCaseStrategy
}

// UppercaseTableAlias indicates if the table alias strategy is uppercase.
func (c *ReflectConfiguration) UppercaseTableAlias() bool {
	return c.HasTableAliasStrategy() && *c.TableAliasStrategy == UpperCaseStrategy
}

// HasTag indicates if the tag is set.
func (c *ReflectConfiguration) HasTag() bool {
	return c.Tag != nil && strings.TrimSpace(*c.Tag) != ""
}

// HasMethodExclusions indicates if method exclusions are set.
func (c *ReflectConfiguration) HasMethodExclusions() bool {
	return len(c.MethodExclusions) > 0
}

// HasFieldExclusions indicates if field exclusions are set.
func (c *ReflectConfiguration) HasFieldExclusions() bool {
	return len(c.FieldExclusions) > 0
}

// HasMethodExclusionPattern indicates if the method exclusion pattern is set.
func (c *ReflectConfiguration) HasMethodExclusionPattern() bool {
	return c.MethodExclusionPattern != nil && strings.TrimSpace(*c.MethodExclusionPattern) != ""
}

// HasFieldExclusionPattern indicates if the field exclusion pattern is set.
func (c *ReflectConfiguration) HasFieldExclusionPattern() bool {
	return c.FieldExclusionPattern != nil && strings.TrimSpace(*c.FieldExclusionPattern) != ""
}

// PluralTableName indicates if the table name is plural.
func (c *ReflectConfiguration) PluralTableName() bool {
	return c.IsTableNamePlural
}

// HasTableAliasLength indicates if the table alias length is set.
func (c *ReflectConfiguration) HasTableAliasLength() bool {
	return c.TableAliasLength != nil
}

var (
	// WithTableName specifies the table name, effectively overriding any
	// table name inference.
	WithTableName = func(name string) ReflectOption {
		return func(c *ReflectConfiguration) {
			c.TableName = &name
			c.IsInferredTableName = false
		}
	}

	// WithTableAlias specifies the table alias, effectively overriding any
	// table alias inference.
	WithTableAlias = func(alias string) ReflectOption {
		return func(c *ReflectConfiguration) {
			c.TableAlias = &alias
			c.IsInferredTableAlias = false
		}
	}

	// WithInferredTableName indicates tha the table name should be inferred
	// based on the provided strategy and plurality.
	WithInferredTableName = func(strategy CaseStrategy, plural bool) ReflectOption {
		return func(c *ReflectConfiguration) {
			c.IsInferredTableName = true
			c.TableNameStrategy = &strategy
			c.IsTableNamePlural = plural
		}
	}

	// WithInferredTableAlias indicates that the table alias should be inferred
	// based on the provided strategy and length.
	WithInferredTableAlias = func(strategy CaseStrategy, length int) ReflectOption {
		return func(c *ReflectConfiguration) {
			if length < 1 {
				length = DefaultTableAliasLength
			}
			c.IsInferredTableAlias = true
			c.TableAliasStrategy = &strategy
			c.TableAliasLength = &length
		}
	}

	// WithInferredColumnNames indicates that the column names should be inferred
	// based on the provided strategy.
	WithInferredColumnNames = func(strategy CaseStrategy) ReflectOption {
		return func(c *ReflectConfiguration) {
			c.IsInferredColumnNames = true
			c.ColumnNameStrategy = &strategy
		}
	}

	// WithTag specifies the struct tag to use for field reflection.
	WithTag = func(tag string) ReflectOption {
		return func(c *ReflectConfiguration) {
			c.Tag = &tag
		}
	}

	// WithoutMethods specifies the methods by name to exclude from reflection.
	WithoutMethods = func(methods ...string) ReflectOption {
		return func(c *ReflectConfiguration) {
			if c.MethodExclusions == nil {
				c.MethodExclusions = []string{}
			}
			c.MethodExclusions = methods
		}
	}

	// WithoutMatchingMethods specifies a regular expression used to match against
	// method names, where matches are excluded from reflection.
	WithoutMatchingMethods = func(pattern string) ReflectOption {
		return func(c *ReflectConfiguration) {
			c.MethodExclusionPattern = &pattern
		}
	}

	// WithoutFields specifies the struct fields by name to exclude from reflection.
	WithoutFields = func(fields ...string) ReflectOption {
		return func(c *ReflectConfiguration) {
			if c.FieldExclusions == nil {
				c.FieldExclusions = []string{}
			}
			c.FieldExclusions = fields
		}
	}

	// WithoutMatchingFields specifies a regular expression used to match against
	// struct field names, where matches are excluded from reflection.
	WithoutMatchingFields = func(pattern string) ReflectOption {
		return func(c *ReflectConfiguration) {
			c.FieldExclusionPattern = &pattern
		}
	}

	// WithPrimaryKeyColumn specifies the name of the column that acts as
	// the primary key.
	WithPrimaryKeyColumn = func(name string) ReflectOption {
		return func(c *ReflectConfiguration) {
			if c.PrimaryKeyColumns == nil {
				c.PrimaryKeyColumns = []string{}
			}
			c.PrimaryKeyColumns = []string{name}
		}
	}

	// WithPrimaryKeyColumns specifies the names of the column that acts together as
	// the primary key.
	WithPrimaryKeyColumns = func(names []string) ReflectOption {
		return func(c *ReflectConfiguration) {
			if c.PrimaryKeyColumns == nil {
				c.PrimaryKeyColumns = []string{}
			}
			pks := append([]string{}, names...)
			c.PrimaryKeyColumns = pks
		}
	}
)
