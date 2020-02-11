## Utils

Currently just has key creation with namespace as a parameter.

Happy to add more.

### Usage

```go
		key := dsutils.IncompleteKeyWithNamespace(model.OptionKind, *namespace, nil)
		key := dsutils.NameKeyWithNamespace(model.OptionKind, *namespace, "Name", nil)
		key := dsutils.IDKeyWithNamespace(model.OptionKind, *namespace, int64(Id), nil)

```