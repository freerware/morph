package morph

import (
	"slices"
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

// NonForeignKeyColumns returns the non-foreign key columns of the child table.
func (r Reference) NonForeignKeyColumns() []Column {
	return r.child.FindColumns(func(c Column) bool {
		return !slices.ContainsFunc(r.foreignKey, func(fk Column) bool {
			return fk.equals(c)
		})
	})
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

// InsertQuery generates an INSERT query for the reference.
func (r *Reference) InsertQuery(options ...QueryOption) (string, error) {
	return r.child.InsertQuery(options...)
}

// InsertQueryWithArgs generates an INSERT query for the reference along with arguments
// derived from the provided object.
func (r *Reference) InsertQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	return r.child.InsertQueryWithArgs(obj, options...)
}

// UpdateQuery generates an UPDATE query for the reference.
func (r *Reference) UpdateQuery(options ...QueryOption) (string, error) {
	return r.child.UpdateQuery(options...)
}

// UpdateQueryWithArgs generates an UPDATE query for the reference along with arguments
// derived from the provided object.
func (r *Reference) UpdateQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	return r.child.UpdateQueryWithArgs(obj, options...)
}

// SelectQuery generates a SELECT query for the reference.
func (r *Reference) SelectQuery(options ...QueryOption) (string, error) {
	generator := newQueryGenerator(&r.child, r.ForeignKey(), r.NonForeignKeyColumns())
	return generator.SelectQuery(options...)
}

// SelectQueryWithArgs generates a SELECT query with arguments for the reference.
func (r *Reference) SelectQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	generator := newQueryGenerator(&r.child, r.ForeignKey(), r.NonForeignKeyColumns())
	return generator.SelectQueryWithArgs(obj, opts...)
}

// DeleteQuery generates a DELETE query for the reference.
func (r *Reference) DeleteQuery(options ...QueryOption) (string, error) {
	generator := newQueryGenerator(&r.child, r.ForeignKey(), r.NonForeignKeyColumns())
	return generator.DeleteQuery(options...)
}

// DeleteQueryWithArgs generates a DELETE query with arguments for the reference.
func (r *Reference) DeleteQueryWithArgs(obj any, options ...QueryOption) (string, []any, error) {
	opts := append(options, WithNamedParameters())
	generator := newQueryGenerator(&r.child, r.ForeignKey(), r.NonForeignKeyColumns())
	return generator.DeleteQueryWithArgs(obj, opts...)
}
