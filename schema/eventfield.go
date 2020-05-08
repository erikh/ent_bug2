package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// EventField holds the schema definition for the EventField entity.
type EventField struct {
	ent.Schema
}

// Fields of the EventField.
func (EventField) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Immutable(),
		field.String("value").Immutable(),
	}
}

// Edges of the EventField.
func (EventField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("field"),
	}
}
