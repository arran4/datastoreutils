package dsutils

import (
	"testing"
	"cloud.google.com/go/datastore"
)

var result *datastore.Key

func TestNameKeyWithNamespace_EmptyKind(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, got none")
		} else {
			if str, ok := r.(string); ok {
				expected := "dsutils: kind cannot be empty"
				if str != expected {
					t.Errorf("Expected panic message %q, got %q", expected, str)
				}
			} else {
				t.Errorf("Expected panic string, got %v", r)
			}
		}
	}()

	_ = NameKeyWithNamespace("", "ns", "name", nil)
}

func BenchmarkNameKeyWithNamespace_Valid(b *testing.B) {
	kind := "MyKind"
	namespace := "MyNamespace"
	name := "MyName"
	var r *datastore.Key

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = NameKeyWithNamespace(kind, namespace, name, nil)
	}
	result = r
}
