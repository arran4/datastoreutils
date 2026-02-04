package dsutils

import (
	"cloud.google.com/go/datastore"
	"errors"
)

// IncompleteKeyWithNamespace creates a new incomplete key.
// The supplied kind cannot be empty.
// The namespace of the new key can be empty.
func IncompleteKeyWithNamespace(kind, namespace string, parent *datastore.Key) (*datastore.Key, error) {
	if kind == "" {
		return nil, errors.New("dsutils: kind cannot be empty")
	}
	if parent != nil && parent.Incomplete() {
		return nil, errors.New("dsutils: can't use an incomplete key as a parent")
	}
	key := &datastore.Key{
		Kind:      kind,
		Parent:    parent,
		Namespace: namespace,
	}
	return key, nil
}

// MustIncompleteKeyWithNamespace creates a new incomplete key and panics on error.
func MustIncompleteKeyWithNamespace(kind, namespace string, parent *datastore.Key) *datastore.Key {
	key, err := IncompleteKeyWithNamespace(kind, namespace, parent)
	if err != nil {
		panic(err)
	}
	return key
}

// NameKeyWithNamespace creates a new key with a name.
// The supplied kind cannot be empty.
// The supplied parent must either be a complete key or nil.
// The namespace of the new key can be empty.
func NameKeyWithNamespace(kind, namespace, name string, parent *datastore.Key) (*datastore.Key, error) {
	if kind == "" {
		return nil, errors.New("dsutils: kind cannot be empty")
	}
	if parent != nil && parent.Incomplete() {
		return nil, errors.New("dsutils: can't use an incomplete key as a parent")
	}
	key := &datastore.Key{
		Kind:      kind,
		Name:      name,
		Parent:    parent,
		Namespace: namespace,
	}
	return key, nil
}

// MustNameKeyWithNamespace creates a new key with a name and panics on error.
func MustNameKeyWithNamespace(kind, namespace, name string, parent *datastore.Key) *datastore.Key {
	key, err := NameKeyWithNamespace(kind, namespace, name, parent)
	if err != nil {
		panic(err)
	}
	return key
}

// IDKeyWithNamespace creates a new key with an ID.
// The supplied kind cannot be empty.
// The supplied parent must either be a complete key or nil.
// The namespace of the new key can be empty.
func IDKeyWithNamespace(kind, namespace string, id int64, parent *datastore.Key) (*datastore.Key, error) {
	if kind == "" {
		return nil, errors.New("dsutils: kind cannot be empty")
	}
	if parent != nil && parent.Incomplete() {
		return nil, errors.New("dsutils: can't use an incomplete key as a parent")
	}
	key := &datastore.Key{
		Kind:      kind,
		ID:        id,
		Parent:    parent,
		Namespace: namespace,
	}
	return key, nil
}

// MustIDKeyWithNamespace creates a new key with an ID and panics on error.
func MustIDKeyWithNamespace(kind, namespace string, id int64, parent *datastore.Key) *datastore.Key {
	key, err := IDKeyWithNamespace(kind, namespace, id, parent)
	if err != nil {
		panic(err)
	}
	return key
}
