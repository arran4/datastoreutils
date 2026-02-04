package dsutils

import (
	"testing"
	"cloud.google.com/go/datastore"
)

func TestIncompleteKeyWithNamespace_PanicsOnIncompleteParent(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			t.Logf("Recovered from panic: %v", r)
		}
	}()

	// Create an incomplete parent key
	parent := datastore.IncompleteKey("ParentKind", nil)

	// Create a key with incomplete parent using the utility
	// This should panic
	IncompleteKeyWithNamespace("Kind", "ns", parent)
}

func BenchmarkIncompleteKeyWithNamespace(b *testing.B) {
	parent := datastore.NameKey("Parent", "name", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IncompleteKeyWithNamespace("Kind", "ns", parent)
	}
}
