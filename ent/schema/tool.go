package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Tool holds the schema definition for the Tool entity.
type Tool struct {
	ent.Schema
}

// Fields of the Tool.
func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Default("unknown"),
		field.String("description"),
		field.String("category"),
		field.String("website"),
		field.String("image_url"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return nil
}
