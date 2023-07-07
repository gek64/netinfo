package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"net/netip"
	"time"
)

type Record struct {
	ent.Schema
}

type NetInterface struct {
	Name string       `json:"name"`
	IPs  []netip.Addr `json:"ips"`
	Mac  string       `json:"mac,omitempty"`
}

// Fields 数据库属性
func (Record) Fields() []ent.Field {

	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),

		field.Time("createdAt").
			Default(time.Now()).
			Immutable(),

		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.String("description").
			Optional(),

		field.JSON("netInterfaces", []NetInterface{}),
	}
}

// Edges 数据库关系
func (Record) Edges() []ent.Edge {
	return nil
}

// Indexes 数据库相关的配置
func (Record) Indexes() []ent.Index {
	return nil
}
