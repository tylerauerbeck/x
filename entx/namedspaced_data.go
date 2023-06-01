package entx

import (
	"encoding/json"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var (
	namespaceMinLength = 5
	namespaceMaxLength = 64

	// ErrUnmarshalJSON is returned when there is an error converting the provided
	// JSON to a json.RawMessage type
	ErrUnmarshalJSON = errors.New("an error occurred parsing json")
)

// NamespacedDataMixin defines an ent Mixin that captures raw json associated with a namespace.
type NamespacedDataMixin struct {
	mixin.Schema
}

// Fields provides the namespace and data fields used in this mixin.
func (m NamespacedDataMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Text("namespace").
			NotEmpty().
			MinLen(namespaceMinLength).
			MaxLen(namespaceMaxLength).
			Annotations(
				entgql.OrderField("NAMESPACE"),
			),
		field.JSON("data", json.RawMessage{}).
			Annotations(
				entgql.Type("JSON"),
				Annotation{IsNamespacedDataJSONField: true},
			),
	}
}

// MarshalRawMessage provides a graphql.Marshaler for json.RawMessage
// func MarshalRawMessage(t json.RawMessage) graphql.Marshaler {
// 	return graphql.WriterFunc(func(w io.Writer) {
// 		s, _ := t.MarshalJSON()
// 		_, _ = io.WriteString(w, string(s))
// 	})
// }

// // UnmarshalRawMessage provides a graphql.Unmarshaler for json.RawMessage
// func UnmarshalRawMessage(v interface{}) (json.RawMessage, error) {
// 	switch j := v.(type) {
// 	case string:
// 		return UnmarshalRawMessage([]byte(j))
// 	case []byte:
// 		return json.RawMessage(j), nil
// 	case map[string]interface{}:
// 		js, err := json.Marshal(v)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return json.RawMessage(js), nil
// 	default:
// 		// Attempt to cast it as a fall back but return an error if it fails
// 		js, err := json.Marshal(v)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return json.RawMessage(js), nil
// 	}
// }
