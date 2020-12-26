package schema

import "github.com/facebookincubator/ent"

// Entity1 holds the schema definition for the Entity1 entity.
type Entity1 struct {
	ent.Schema
}

// Fields of the Entity1.
func (Entity1) Fields() []ent.Field {
	return nil
}

// Edges of the Entity1.
func (Entity1) Edges() []ent.Edge {
	return nil
}
