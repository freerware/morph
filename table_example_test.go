package morph_test

import (
	"fmt"

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
}

func ExampleTable_SelectQueryWithArgs_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.SelectQueryWithArgs(mf)
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// SELECT category, class, crew, id, length, length_unit, name, speed, speed_unit FROM starships AS S WHERE 1=1 AND S.id = ?;
	// [1]
}

func ExampleTable_SelectQueryWithArgs_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.SelectQueryWithArgs(mf, morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// SELECT category, class, crew, id, length, length_unit, name, speed, speed_unit FROM starships AS S WHERE 1=1 AND S.id = $1;
	// [1]
}

func ExampleTable_SelectQuery_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.SelectQuery()
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: SELECT category, class, crew, id, length, length_unit, name, speed, speed_unit FROM starships AS S WHERE 1=1 AND S.id = ?;
}

func ExampleTable_SelectQuery_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.SelectQuery(morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: SELECT category, class, crew, id, length, length_unit, name, speed, speed_unit FROM starships AS S WHERE 1=1 AND S.id = $1;
}

func ExampleTable_UpdateQueryWithArgs_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.UpdateQueryWithArgs(mf)
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// UPDATE starships AS S SET S.category = ?, S.class = ?, S.crew = ?, S.length = ?, S.length_unit = ?, S.name = ?, S.speed = ?, S.speed_unit = ? WHERE 1=1 AND S.id = ?;
	// [Light Freighter Yacht 4 34.75 m Millennium Falcon 1050 km/h 1]
}

func ExampleTable_UpdateQueryWithArgs_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.UpdateQueryWithArgs(mf, morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// UPDATE starships AS S SET S.category = $1, S.class = $2, S.crew = $3, S.length = $4, S.length_unit = $5, S.name = $6, S.speed = $7, S.speed_unit = $8 WHERE 1=1 AND S.id = $9;
	// [Light Freighter Yacht 4 34.75 m Millennium Falcon 1050 km/h 1]
}

func ExampleTable_UpdateQuery_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.UpdateQuery()
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: UPDATE starships AS S SET S.category = ?, S.class = ?, S.crew = ?, S.length = ?, S.length_unit = ?, S.name = ?, S.speed = ?, S.speed_unit = ? WHERE 1=1 AND S.id = ?;
}

func ExampleTable_UpdateQuery_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.UpdateQuery(morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: UPDATE starships AS S SET S.category = $1, S.class = $2, S.crew = $3, S.length = $4, S.length_unit = $5, S.name = $6, S.speed = $7, S.speed_unit = $8 WHERE 1=1 AND S.id = $9;
}

func ExampleTable_DeleteQueryWithArgs_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.DeleteQueryWithArgs(mf)
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// DELETE FROM starships WHERE 1=1 AND id = ?;
	// [1]
}

func ExampleTable_DeleteQueryWithArgs_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.DeleteQueryWithArgs(mf, morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// DELETE FROM starships WHERE 1=1 AND id = $1;
	// [1]
}

func ExampleTable_DeleteQuery_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.DeleteQuery()
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: DELETE FROM starships WHERE 1=1 AND id = ?;
}

func ExampleTable_DeleteQuery_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.DeleteQuery(morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: DELETE FROM starships WHERE 1=1 AND id = $1;
}

func ExampleTable_InsertQueryWithArgs_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.InsertQueryWithArgs(mf)
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// INSERT INTO starships (category, class, crew, id, length, length_unit, name, speed, speed_unit) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	// [Light Freighter Yacht 4 1 34.75 m Millennium Falcon 1050 km/h]
}

func ExampleTable_InsertQueryWithArgs_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, args, err := table.InsertQueryWithArgs(mf, morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	fmt.Println(args)
	// Output:
	// INSERT INTO starships (category, class, crew, id, length, length_unit, name, speed, speed_unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	// [Light Freighter Yacht 4 1 34.75 m Millennium Falcon 1050 km/h]
}

func ExampleTable_InsertQuery_noOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.InsertQuery()
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: INSERT INTO starships (category, class, crew, id, length, length_unit, name, speed, speed_unit) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
}

func ExampleTable_InsertQuery_withOptions() {
	mf := Starship{
		ID:         1,
		Name:       "Millennium Falcon",
		Category:   "Light Freighter",
		Class:      "Yacht",
		Length:     34.75,
		LengthUnit: "m",
		Crew:       4,
		Speed:      1050,
		SpeedUnit:  "km/h",
	}

	table := morph.Must(morph.Reflect(mf))

	query, err := table.InsertQuery(morph.WithPlaceholder("$", true))
	if err != nil {
		panic(err)
	}

	fmt.Println(query)
	// Output: INSERT INTO starships (category, class, crew, id, length, length_unit, name, speed, speed_unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
}
