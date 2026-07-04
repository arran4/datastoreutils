package dsutils

import (
	"cloud.google.com/go/datastore"
	"testing"
)

var result *datastore.Key

func TestMustNameKeyWithNamespace_EmptyKind(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, got none")
		} else {
			if err, ok := r.(error); ok {
				expected := "dsutils: kind cannot be empty"
				if err.Error() != expected {
					t.Errorf("Expected panic message %q, got %q", expected, err.Error())
				}
			} else {
				t.Errorf("Expected panic error, got %v", r)
			}
		}
	}()

	_ = MustNameKeyWithNamespace("", "ns", "name", nil)
}

func TestNameKeyWithNamespace_EmptyKind(t *testing.T) {
	key, err := NameKeyWithNamespace("", "ns", "name", nil)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if key != nil {
		t.Errorf("Expected nil key, got %v", key)
	}
	expected := "dsutils: kind cannot be empty"
	if err.Error() != expected {
		t.Errorf("Expected error message %q, got %q", expected, err.Error())
	}
}

func TestIncompleteKeyWithNamespace_EmptyKind(t *testing.T) {
	key, err := IncompleteKeyWithNamespace("", "ns", nil)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if key != nil {
		t.Errorf("Expected nil key, got %v", key)
	}
}

func TestIDKeyWithNamespace_EmptyKind(t *testing.T) {
	key, err := IDKeyWithNamespace("", "ns", 123, nil)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if key != nil {
		t.Errorf("Expected nil key, got %v", key)
	}
}

func TestNameKeyWithNamespace_IncompleteParent(t *testing.T) {
	parent := datastore.IncompleteKey("ParentKind", nil)
	key, err := NameKeyWithNamespace("Kind", "ns", "name", parent)
	if err == nil {
		t.Error("Expected error for incomplete parent, got nil")
	}
	if key != nil {
		t.Errorf("Expected nil key, got %v", key)
	}
	expected := "dsutils: can't use an incomplete key as a parent"
	if err.Error() != expected {
		t.Errorf("Expected error message %q, got %q", expected, err.Error())
	}
}

func TestIncompleteKeyWithNamespace_IncompleteParent(t *testing.T) {
	parent := datastore.IncompleteKey("ParentKind", nil)
	key, err := IncompleteKeyWithNamespace("Kind", "ns", parent)
	if err == nil {
		t.Error("Expected error for incomplete parent, got nil")
	}
	if key != nil {
		t.Errorf("Expected nil key, got %v", key)
	}
	expected := "dsutils: can't use an incomplete key as a parent"
	if err.Error() != expected {
		t.Errorf("Expected error message %q, got %q", expected, err.Error())
	}
}

func TestIDKeyWithNamespace_IncompleteParent(t *testing.T) {
	parent := datastore.IncompleteKey("ParentKind", nil)
	key, err := IDKeyWithNamespace("Kind", "ns", 123, parent)
	if err == nil {
		t.Error("Expected error for incomplete parent, got nil")
	}
	if key != nil {
		t.Errorf("Expected nil key, got %v", key)
	}
	expected := "dsutils: can't use an incomplete key as a parent"
	if err.Error() != expected {
		t.Errorf("Expected error message %q, got %q", expected, err.Error())
	}
}

func BenchmarkNameKeyWithNamespace_Valid(b *testing.B) {
	kind := "MyKind"
	namespace := "MyNamespace"
	name := "MyName"
	var r *datastore.Key

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, _ = NameKeyWithNamespace(kind, namespace, name, nil)
	}
	result = r
}

func BenchmarkMustNameKeyWithNamespace_Valid(b *testing.B) {
	kind := "MyKind"
	namespace := "MyNamespace"
	name := "MyName"
	var r *datastore.Key

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = MustNameKeyWithNamespace(kind, namespace, name, nil)
	}
	result = r
}
