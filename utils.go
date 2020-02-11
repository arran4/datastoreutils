package dsutils

import (
	"cloud.google.com/go/datastore"
)

// IncompleteKey creates a new incomplete key.
// The supplied kind cannot be empty.
// The namespace of the new key can be empty.
func IncompleteKeyWithNamespace(kind, namespace string, parent *datastore.Key) *datastore.Key {
	key := &datastore.Key{
		Kind:   kind,
		Parent: parent,
		Namespace: namespace,
	}
	return key
}

// NameKey creates a new key with a name.
// The supplied kind cannot be empty.
// The supplied parent must either be a complete key or nil.
// The namespace of the new key can be empty.
func NameKeyWithNamespace(kind, namespace, name string, parent *datastore.Key) *datastore.Key {
	key := &datastore.Key{
		Kind:   kind,
		Name:   name,
		Parent: parent,
		Namespace: namespace,
	}
	return key
}

// IDKey creates a new key with an ID.
// The supplied kind cannot be empty.
// The supplied parent must either be a complete key or nil.
// The namespace of the new key can be empty.
func IDKeyWithNamespace(kind, namespace string, id int64, parent *datastore.Key) *datastore.Key {
	key := &datastore.Key{
		Kind:   kind,
		ID:     id,
		Parent: parent,
		Namespace: namespace,
	}
	return key
}
