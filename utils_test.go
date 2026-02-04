package dsutils

import (
	"testing"
	"cloud.google.com/go/datastore"
)

func BenchmarkIDKeyWithNamespace(b *testing.B) {
	parent := &datastore.Key{Kind: "Parent", ID: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IDKeyWithNamespace("Kind", "Namespace", 123, parent)
	}
}

func TestIDKeyWithNamespace_EmptyKind(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = IDKeyWithNamespace("", "ns", 123, nil)
}
