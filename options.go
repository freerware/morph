package morph

type Option func(*ReflectConfiguration)

type ReflectConfiguration struct {
	TableName              *string
	TableNameStrategy      *TableNameStrategy
	TableAliasStrategy     *TableAliasStrategy
	Tag                    *string
	MethodExclusions       []string
	FieldExclusions        []string
	MethodExclusionPattern *string
	FieldExclusionPattern  *string
	IsTableNamePlural      *bool
	TableAliasLength       *int
}

type TableNameStrategy int
type TableAliasStrategy int

const (
	SnakeTableNameStrategy TableNameStrategy = iota
	CamelTableNameStrategy
)

const (
	LowercaseTableAliasStrategy TableAliasStrategy = iota
	UppercaseTableAliasStrategy
)

var (
	WithTableName = func(name string) Option {
		return func(c *ReflectConfiguration) {
			c.TableName = &name
		}
	}

	WithInferredTableName = func(strategy TableNameStrategy, plural bool) Option {
		return func(c *ReflectConfiguration) {
			c.TableNameStrategy = &strategy
			c.IsTableNamePlural = &plural
		}
	}

	WithInferredTableAlias = func(strategy TableAliasStrategy, length int) Option {
		return func(c *ReflectConfiguration) {
			c.TableAliasStrategy = &strategy
			c.TableAliasLength = &length
		}
	}

	WithTag = func(tag string) Option {
		return func(c *ReflectConfiguration) {
			c.Tag = &tag
		}
	}

	WithoutMethods = func(methods []string) Option {
		return func(c *ReflectConfiguration) {
			if c.MethodExclusions == nil {
				c.MethodExclusions = []string{}
			}
			c.MethodExclusions = methods
		}
	}

	WithoutMatchingMethods = func(pattern string) Option {
		return func(c *ReflectConfiguration) {
			c.MethodExclusionPattern = &pattern
		}
	}

	WithoutFields = func(fields []string) Option {
		return func(c *ReflectConfiguration) {
			if c.FieldExclusions == nil {
				c.FieldExclusions = []string{}
			}
			c.FieldExclusions = fields
		}
	}

	WithoutMatchingFields = func(pattern string) Option {
		return func(c *ReflectConfiguration) {
			c.FieldExclusionPattern = &pattern
		}
	}
)
