package internal

var (
	// SQL identifiers and key words must begin with a letter (a-z, but also
	// letters with diacritical marks and non-Latin letters) or an underscore
	// (_). Subsequent characters in an identifier or key word can be letters,
	// underscores, digits (0-9), or dollar signs ($).
	// https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
	identifierExpr = `[a-zA-Z_\p{L}][a-zA-Z0-9_\p{L}]`
	schemaExpr     = `"?(?P<schema>` + identifierExpr + `+)"?`
	nameExpr       = `"?(?P<name>` + identifierExpr + `+)"?`

	expressions = []string{
		`(?:CREATE|ALTER) (?P<type>AGGREGATE) ` + schemaExpr + `\.` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>DOMAIN) ` + schemaExpr + `\.` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>FUNCTION) ` + schemaExpr + `\.` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>INDEX) ` + nameExpr + ` ON ` + schemaExpr + `\.`,
		`(?:CREATE|ALTER) (?P<type>MATERIALIZED) VIEW ` + schemaExpr + `\.` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>SCHEMA) ` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>SEQUENCE) ` + schemaExpr + `\.` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>TABLE) ` + schemaExpr + `\.` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>TYPE) ` + schemaExpr + `\.` + nameExpr,
		`(?:CREATE|ALTER) (?P<type>VIEW) ` + schemaExpr + `\.` + nameExpr,
		`ALTER (?P<type>TABLE) ONLY ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>TRIGGER) ` + nameExpr + ` .* ON ` + schemaExpr + `\.`,
		`CREATE UNIQUE (?P<type>INDEX) ` + nameExpr + ` ON ` + schemaExpr + `\.`,
	}
)
