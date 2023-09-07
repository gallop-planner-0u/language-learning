package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("user_id"),
		field.String("word").Unique(),
		field.String("translation"),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return nil
}
