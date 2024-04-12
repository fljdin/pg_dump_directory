package internal

import (
	"fmt"
	"regexp"
)

type Statement struct {
	Raw    string
	Name   string
	Schema string
	Type   string
}

func (s Statement) String() string {
	return fmt.Sprintf("type=%s schema=%s relname=%s", s.Type, s.Schema, s.Name)
}

func Parse(raw string) Statement {
	for _, expr := range expressions {
		re := regexp.MustCompile(expr)

		groupIndex := func(name string) int {
			for i, g := range re.SubexpNames() {
				if g == name {
					return i
				}
			}
			return 0
		}

		match := re.FindStringSubmatch(raw)
		if match != nil {
			return Statement{
				Raw:    raw,
				Name:   match[groupIndex("name")],
				Schema: match[groupIndex("schema")],
				Type:   match[groupIndex("type")],
			}
		}
	}

	return Statement{
		Type: "unknown",
		Raw:  raw,
	}
}
