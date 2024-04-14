package internal_test

import (
	"testing"

	"github.com/fljdin/pg_dump_tree/internal"
	"github.com/stretchr/testify/require"
)

var fixtures = []internal.Statement{
	{Type: "unknown", Raw: `SET default_tablespace = '';`},
	{Type: "DOMAIN", Raw: `CREATE DOMAIN public."bigint" AS bigint;`, Schema: "public", Name: "bigint"},
	{Type: "INDEX", Raw: `CREATE UNIQUE INDEX idx_staff_id ON public.staff`, Schema: "public", Name: "idx_staff_id"},
	{Type: "OWNER", Raw: `ALTER TABLE public.add_seq OWNER TO postgres;`, Schema: "public", Name: "add_seq"},
	{Type: "TABLE", Raw: `ALTER TABLE ONLY public.staff`, Schema: "public", Name: "staff"},
	{Type: "TABLE", Raw: `CREATE TABLE public.staff`, Schema: "public", Name: "staff"},
}

func TestParse(t *testing.T) {
	for _, fixture := range fixtures {
		stmt := internal.Parse(fixture.Raw)
		require.Equal(t, fixture, stmt)
	}
}
