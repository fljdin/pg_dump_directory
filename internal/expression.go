package internal

var (
	// SQL identifiers and key words must begin with a letter (a-z, but also
	// letters with diacritical marks and non-Latin letters) or an underscore
	// (_). Subsequent characters in an identifier or key word can be letters,
	// underscores, digits (0-9), or dollar signs ($).
	// https://www.postgresql.org/docs/current/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS
	identifierExpr = `[a-zA-Z_\p{L}][a-zA-Z0-9_\p{L}]+`
	schemaExpr     = `"?(?P<schema>` + identifierExpr + `)"?`
	nameExpr       = `"?(?P<name>` + identifierExpr + `)"?`

	expressions = []string{
		`(?P<type>(?:GRANT|REVOKE)) .* ON ` + schemaExpr + `\.` + nameExpr,
		`(?P<type>(?:GRANT|REVOKE)) .* ON SCHEMA` + schemaExpr,
		`ALTER .* ` + schemaExpr + `\.` + nameExpr + `.* (?P<type>OWNER) TO`,
		`ALTER (?P<type>TABLE) ONLY ` + schemaExpr + `\.` + nameExpr,
		`ALTER SCHEMA ` + schemaExpr + ` (?P<type>OWNER) TO`,
		`CREATE (?P<type>AGGREGATE) ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>DOMAIN) ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>FUNCTION) ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>INDEX) ` + nameExpr + ` ON ` + schemaExpr + `\.`,
		`CREATE (?P<type>MATERIALIZED) VIEW ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>SCHEMA) ` + nameExpr,
		`CREATE (?P<type>SEQUENCE) ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>TABLE) ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>TRIGGER) ` + nameExpr + ` .* ON ` + schemaExpr + `\.`,
		`CREATE (?P<type>TYPE) ` + schemaExpr + `\.` + nameExpr,
		`CREATE (?P<type>VIEW) ` + schemaExpr + `\.` + nameExpr,
		`CREATE UNIQUE (?P<type>INDEX) ` + nameExpr + ` ON ` + schemaExpr + `\.`,
	}
)
