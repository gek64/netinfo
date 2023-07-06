package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"net"
	"time"
)

type Record struct {
	ent.Schema
}

type NetInterface struct {
	Name string   `json:"name"`
	IPs  []net.IP `json:"ips"`
}

// Fields 数据库属性
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),

		field.Time("created_at").
			Default(time.Now()).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.String("description"),

		field.JSON("net_interfaces", []NetInterface{}),
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
