# morph
A compact package for organizing and maintaining your entity database metadata.

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci]
[![Coverage Status][coverage-img]][coverage] [![Release][release-img]][release]
[![License][license-img]][license]

## What is it?

`morph` organizes and maintains the necessary metadata to map your entities to
and from relational database tables. This is accomplished using the
[Metadata Mapping][metadata-mapping] pattern popularized by Martin Fowler.
With these metadata mappings, your application is empowered to construct SQL
queries dynamically using the entities themselves.

## Why use it?

With `morph`, your application reaps several benefits:

- dynamic construction of queries using entities and their fields.
- metadata generation using files in several formats, including [YAML][yaml] and [JSON][json].
- decoupling of code responsible for manufacturing queries from code tasked with SQL generation.

## How to use it?

Using `morph` is super straightforward. You utilize [`Table`][table-doc] and
[`Column`][column-doc] to organize metadata for your entities and their
associated relational representations. Let's suppose you have a `User` entity
in your application (`user`):
```go
var usernameCol morph.Column
usernameCol.SetName("username")
usernameCol.SetField(Fields.Username)

var passwordCol morph.Column
passwordCol.SetName("password")
passwordCol.SetField(Fields.Password)

var userTable morph.Table
userTable.SetName("user")
userTable.SetAlias("U")
userTable.SetType(user)
userTable.AddColumns(usernameCol, passwordCol)
```

### Loading

Capturing the metadata mappings can be tedious, especially if your application
has many entities with corresponding relational representations. Instead
of constructing them manually, we recommend loading in a file that
specifies the metadata mapping configuration:
```json
{
  "tables": [
    {
      "typeName": "example.User",
      "name": "user",
      "alias": "U",
      "columns": [
        {
          "name": "username",
          "field": "Username"
        },
        {
          "name": "password",
          "field": "Password"
        }
      ]
    }
  ]
}
```
```go
configuration, err := morph.Load("./metadata.json")
if err != nil {
	panic(err)
}

tables := configuration.AsMetadata()
```

### Custom Loader

At this time, we currently support YAML (`.yaml`, `.yml`) and JSON (`.json`)
configuration files. However, if you would like to utilize a different file
format, you can construct a type that implements [`morph.Loader`][loader-doc]
and add the appropriate entries in [`morph.Loaders`][loaders-doc]. The
[`morph.Load`][load-doc] function will leverage `morph.Loaders` by extracting
the file extension using the path provided to it.

## Contribute

Want to lend us a hand? Check out our guidelines for
[contributing][contributing].

## License

We are rocking an [Apache 2.0 license][apache-license] for this project.

## Code of Conduct

Please check out our [code of conduct][code-of-conduct] to get up to speed
how we do things.

[contributing]: https://github.com/freerware/morph/blob/master/CONTRIBUTING.md
[apache-license]: https://github.com/freerware/morph/blob/master/LICENSE.txt
[code-of-conduct]: https://github.com/freerware/morph/blob/master/CODE_OF_CONDUCT.md
[doc-img]: https://godoc.org/github.com/freerware/morph?status.svg
[doc]: https://godoc.org/github.com/freerware/morph
[ci-img]: https://travis-ci.org/freerware/morph.svg?branch=master
[ci]: https://travis-ci.org/freerware/morph
[coverage-img]: https://coveralls.io/repos/github/freerware/morph/badge.svg?branch=master
[coverage]: https://coveralls.io/github/freerware/morph?branch=master
[license]: https://opensource.org/licenses/Apache-2.0
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[release]: https://github.com/freerware/morph/releases
[release-img]: https://img.shields.io/github/tag/freerware/morph.svg?label=version
[loaders-doc]: https://godoc.org/github.com/freerware/morph#Loaders
[loader-doc]: https://godoc.org/github.com/freerware/morph#Loader
[load-doc]: https://godoc.org/github.com/freerware/morph#Load
[metadata-mapping]: https://www.martinfowler.com/eaaCatalog/metadataMapping.html
[table-doc]: https://godoc.org/github.com/freerware/morph#Table
[column-doc]: https://godoc.org/github.com/freerware/morph#Column
[yaml]: https://yaml.org/
[json]: https://www.json.org/
