package cypher

import (
	"testing"
)

func TestQueryToString(t *testing.T) {
	strQuery := `MATCH (user :User {name: "Adam"}) RETURN user`

	q := Query{}
	user := "user"
	node := NodePattern{
		Variable: &user,
		Labels:   []string{"User"},
		Properties: map[string]Expr{
			"name": Literal("Adam"),
		},
	}
	m := MatchPattern{Elements: []PatternElement{node}}
	q.Stmt = SingleQuery{
		Reading: []ReadingClause{
			{Pattern: []MatchPattern{m}},
		},
		ReturnItems: []Expr{
			Symbol(user),
		},
	}

	r := q.String()
	if r != strQuery {
		t.Errorf("Did not generate correct query: \nExpected:\n\t%s\nGot:\n\t%s", strQuery, r)
	}
}
