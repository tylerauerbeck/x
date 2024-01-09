// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entx

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// AuditMixin defines an interface of a Mixin that provides the created_at and updated_at timestamp fields
type AuditMixin interface {
	ent.Mixin
	CreatedByAnnotations(...schema.Annotation) AuditMixin
	UpdatedByAnnotations(...schema.Annotation) AuditMixin
}

type auditMixin struct {
	mixin.Schema
	createdByAnnotations []schema.Annotation
	updatedByAnnotations []schema.Annotation
}

// NewAuditMixin returns a Mixin that provides the created_by and updated_by timestamp fields
func NewAuditMixin() AuditMixin {
	return auditMixin{
		createdByAnnotations: []schema.Annotation{
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			entgql.OrderField("CREATED_BY"),
		},
		updatedByAnnotations: []schema.Annotation{
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			entgql.OrderField("UPDATED_BY"),
		},
	}
}

func (m auditMixin) CreatedByAnnotations(ants ...schema.Annotation) AuditMixin {
	m.createdByAnnotations = append(m.createdByAnnotations, ants...)
	return m
}

func (m auditMixin) UpdatedByAnnotations(ants ...schema.Annotation) AuditMixin {
	m.updatedByAnnotations = append(m.updatedByAnnotations, ants...)
	return m
}

// Fields provides the created_at and updated_at fields
func (m auditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").
			Optional().
			Immutable(),
		field.String("updated_by").
			Optional(),
	}
}

// Indexes provides indexes on both created_at and updated_at fields
func (auditMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_by"),
		index.Fields("updated_by"),
	}
}
