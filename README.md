# pg_dump_tree

**THIS README IS A PROJECT PREVIEW.**

Simple utility to write file for each relation in a structured directory tree.

## Description

This utility is intended to be used in conjunction with `pg_dump` tool. It
parses the plain format output of `pg_dump` and writes a file for each relation
in a structured directory tree. Fully written in Go, bring alive with
[Fragment](https://github.com/fljdin/fragment) library.

## Usage

```sh
pg_dump --format=plain --schema-only | pg_dump_tree --output-dir=directory
```

## Options

```sh
pg_dump_tree --help
```

## Directory Structure

```
directory
└── schema
    ├── FUNCTION
    │   └── function1.sql
    ├── SEQUENCE
    │   └── sequence1.sql
    ├── TABLE
    │   ├── table1.sql
    │   └── table2.sql
    └── VIEW
        └── view1.sql
```

## Installation

```sh
go install github.com/fljdin/pg_dump_tree@main
```

## References

The `pg_dump` utility comes with a `directory` option but only for compressed
data in a `COPY` compatible format, required to bulk inserting with multiple
processes. The database definition is dumped in a unique file named `toc.dat`.

People looking for a structured tree had various solutions in the past.

- https://www.postgresql.org/message-id/AANLkTikLHA2x6U=q-t0j0YS78txHFmdtyxJfsrsRcLqN@mail.gmail.com
- https://stackoverflow.com/questions/18330358/postgresql-dump-each-table-into-a-different-file
- https://dba.stackexchange.com/questions/287593/pg-dump-how-to-split-it-into-directories-and-files
- https://github.com/kruckenb/split_postgres_dump

## Contributing

Please, feel free to contribute by opening issues or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
