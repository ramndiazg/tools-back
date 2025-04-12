package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_uuid", uuid.UUID{}),
		field.UUID("tool_uuid", uuid.UUID{}),
		field.Int("rating"),
		field.String("comment"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return nil
}
