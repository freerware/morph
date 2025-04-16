package morph_benchmark

import (
	"testing"

	"github.com/freerware/morph"
)

type Starship struct {
	ID         int     `morph:"id"`
	Name       string  `morph:"ship_name"`
	Category   string  `morph:"category"`
	Class      string  `morph:"class"`
	Length     float64 `morph:"length"`
	LengthUnit string  `morph:"length_unit"`
	Crew       int     `morph:"crew_capacity"`
	Speed      float64 `morph:"max_speed"`
	SpeedUnit  string  `morph:"speed_unit"`
	Owners     []Owner
}

type Owner struct {
	ID         int    `morph:"id"`
	Name       string `morph:"name"`
	StarshipID int    `morph:"starship_id"`
}

var (
	millenniumFalcon = Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
		Owners: []Owner{
			{
				ID:         1,
				Name:       "Han Solo",
				StarshipID: 1,
			},
			{
				ID:         2,
				Name:       "Chewbacca",
				StarshipID: 1,
			},
			{
				ID:         3,
				Name:       "Lando Calrissian",
				StarshipID: 1,
			},
			{
				ID:         4,
				Name:       "Nien Nunb",
				StarshipID: 1,
			},
		},
	}
)

func BenchmarkReflect_DefaultOptions(b *testing.B) {
	for b.Loop() {
		_, err := morph.Reflect(millenniumFalcon)
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReflect_WithTag(b *testing.B) {
	for b.Loop() {
		_, err := morph.Reflect(millenniumFalcon, morph.WithTag("morph"))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReflect_WithoutFields(b *testing.B) {
	for b.Loop() {
		_, err := morph.Reflect(millenniumFalcon, morph.WithoutFields("Length", "LengthUnit", "Speed", "SpeedUnit"))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReflect_WithoutMatchingFields(b *testing.B) {
	for b.Loop() {
		_, err := morph.Reflect(millenniumFalcon, morph.WithoutMatchingFields(`^\w+Unit`))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReflect_WithPrimaryKeyColumn(b *testing.B) {
	for b.Loop() {
		_, err := morph.Reflect(millenniumFalcon, morph.WithPrimaryKeyColumn("ID"))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReflect_WithPrimaryKeyColumns(b *testing.B) {
	for b.Loop() {
		_, err := morph.Reflect(millenniumFalcon, morph.WithPrimaryKeyColumns("ID"))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReflect_WithTableName(b *testing.B) {
	for b.Loop() {
		_, err := morph.Reflect(millenniumFalcon, morph.WithTableName("RazorCrest"))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableInsertQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.InsertQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableInsertQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.InsertQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableInsertQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.InsertQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.UpdateQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.UpdateQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.UpdateQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQuery_WithoutEmptyValues(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.UpdateQuery(morph.WithoutEmptyValues(millenniumFalcon))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableDeleteQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.DeleteQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableDeleteQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.DeleteQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableDeleteQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.DeleteQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableSelectQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.SelectQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableSelectQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.SelectQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableSelectQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, err := table.SelectQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableMustInsertQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustInsertQuery()
	}
}

func BenchmarkTableMustInsertQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustInsertQuery(morph.WithPlaceholder("$", true))
	}
}

func BenchmarkTableMustInsertQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustInsertQuery(morph.WithNamedParameters())
	}
}

func BenchmarkTableMustUpdateQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustUpdateQuery()
	}
}

func BenchmarkTableMustUpdateQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustUpdateQuery(morph.WithPlaceholder("$", true))
	}
}

func BenchmarkTableMustUpdateQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustUpdateQuery(morph.WithNamedParameters())
	}
}

func BenchmarkTableMustUpdateQuery_WithoutEmptyValues(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustUpdateQuery(morph.WithoutEmptyValues(millenniumFalcon))
	}
}

func BenchmarkTableMustDeleteQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustDeleteQuery()
	}
}

func BenchmarkTableMustDeleteQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustDeleteQuery(morph.WithPlaceholder("$", true))
	}
}

func BenchmarkTableMustDeleteQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustDeleteQuery(morph.WithNamedParameters())
	}
}

func BenchmarkTableMustSelectQuery_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustSelectQuery()
	}
}

func BenchmarkTableMustSelectQuery_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustSelectQuery(morph.WithPlaceholder("$", true))
	}
}

func BenchmarkTableMustSelectQuery_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		table.MustSelectQuery(morph.WithNamedParameters())
	}
}

func BenchmarkTableInsertQueryWithArgs_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.InsertQueryWithArgs(millenniumFalcon)
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableInsertQueryWithArgs_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.InsertQueryWithArgs(millenniumFalcon, morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableInsertQueryWithArgs_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.InsertQueryWithArgs(millenniumFalcon, morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQueryWithArgs_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.UpdateQueryWithArgs(millenniumFalcon)
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQueryWithArgs_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.UpdateQueryWithArgs(millenniumFalcon, morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQueryWithArgs_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.UpdateQueryWithArgs(millenniumFalcon, morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableUpdateQueryWithArgs_WithoutEmptyValues(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.UpdateQueryWithArgs(millenniumFalcon, morph.WithoutEmptyValues(millenniumFalcon))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableDeleteQueryWithArgs_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.DeleteQueryWithArgs(millenniumFalcon)
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableDeleteQueryWithArgs_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.DeleteQueryWithArgs(millenniumFalcon, morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableDeleteQueryWithArgs_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.DeleteQueryWithArgs(millenniumFalcon, morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableSelectQueryWithArgs_DefaultOptions(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.SelectQueryWithArgs(millenniumFalcon)
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableSelectQueryWithArgs_WithPlaceholder(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.SelectQueryWithArgs(millenniumFalcon, morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkTableSelectQueryWithArgs_WithNamedParameters(b *testing.B) {
	table, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	for b.Loop() {
		_, _, err := table.SelectQueryWithArgs(millenniumFalcon, morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

//

func BenchmarkReferenceInsertQueryWithArgs_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.InsertQueryWithArgs(millenniumFalcon.Owners[0])
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceInsertQueryWithArgs_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.InsertQueryWithArgs(millenniumFalcon.Owners[0], morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceInsertQueryWithArgs_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.InsertQueryWithArgs(millenniumFalcon.Owners[0], morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQueryWithArgs_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.UpdateQueryWithArgs(millenniumFalcon.Owners[0])
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQueryWithArgs_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.UpdateQueryWithArgs(millenniumFalcon.Owners[0], morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQueryWithArgs_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.UpdateQueryWithArgs(millenniumFalcon.Owners[0], morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQueryWithArgs_WithoutEmptyValues(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.UpdateQueryWithArgs(millenniumFalcon.Owners[0], morph.WithoutEmptyValues(millenniumFalcon.Owners[0]))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceDeleteQueryWithArgs_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.DeleteQueryWithArgs(millenniumFalcon.Owners[0])
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceDeleteQueryWithArgs_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.DeleteQueryWithArgs(millenniumFalcon.Owners[0], morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceDeleteQueryWithArgs_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.DeleteQueryWithArgs(millenniumFalcon.Owners[0], morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceSelectQueryWithArgs_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.SelectQueryWithArgs(millenniumFalcon.Owners[0])
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceSelectQueryWithArgs_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.SelectQueryWithArgs(millenniumFalcon.Owners[0], morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceSelectQueryWithArgs_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, _, err := ref.SelectQueryWithArgs(millenniumFalcon.Owners[0], morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

//

func BenchmarkReferenceInsertQuery_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.InsertQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceInsertQuery_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.InsertQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceInsertQuery_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.InsertQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQuery_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.UpdateQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQuery_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.UpdateQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQuery_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.UpdateQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceUpdateQuery_WithoutEmptyValues(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.UpdateQuery(morph.WithoutEmptyValues(millenniumFalcon.Owners[0]))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceDeleteQuery_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.DeleteQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceDeleteQuery_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.DeleteQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceDeleteQuery_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.DeleteQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceSelectQuery_DefaultOptions(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.SelectQuery()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceSelectQuery_WithPlaceholder(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.SelectQuery(morph.WithPlaceholder("$", true))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkReferenceSelectQuery_WithNamedParameters(b *testing.B) {
	parent, err := morph.Reflect(millenniumFalcon)
	if err != nil {
		b.FailNow()
	}

	child, err := morph.Reflect(millenniumFalcon.Owners[0])
	if err != nil {
		b.FailNow()
	}

	ref, err := child.References(&parent, child.FindColumns(func(c morph.Column) bool {
		return c.Name() == "starship_id"
	}))

	for b.Loop() {
		_, err := ref.SelectQuery(morph.WithNamedParameters())
		if err != nil {
			b.FailNow()
		}
	}
}
