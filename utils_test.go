package dsutils

import (
	"testing"
	"cloud.google.com/go/datastore"
)

var ResultKey *datastore.Key

func BenchmarkIncompleteKeyWithNamespace_Valid(b *testing.B) {
	b.ReportAllocs()
	var k *datastore.Key
	for i := 0; i < b.N; i++ {
		k = IncompleteKeyWithNamespace("MyKind", "MyNamespace", nil)
	}
	ResultKey = k
}

func BenchmarkIncompleteKeyWithNamespace_Invalid(b *testing.B) {
	b.ReportAllocs()
	var k *datastore.Key
	for i := 0; i < b.N; i++ {
		k = IncompleteKeyWithNamespace("", "MyNamespace", nil)
	}
	ResultKey = k
}

func TestIncompleteKeyWithNamespace_EmptyKind(t *testing.T) {
	key := IncompleteKeyWithNamespace("", "ns", nil)
	if key != nil {
		t.Errorf("Expected nil for empty kind, got %+v", key)
	}
}
