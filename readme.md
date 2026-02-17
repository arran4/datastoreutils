## Utils

Currently just has key creation with namespace as a parameter.

Happy to add more.

### Usage

```go
// Using Must versions (panics on error)
key := dsutils.MustIncompleteKeyWithNamespace(model.OptionKind, "namespace", nil)
key := dsutils.MustNameKeyWithNamespace(model.OptionKind, "namespace", "Name", nil)
key := dsutils.MustIDKeyWithNamespace(model.OptionKind, "namespace", int64(Id), nil)

// Using error return versions
key, err := dsutils.IncompleteKeyWithNamespace(model.OptionKind, "namespace", nil)
if err != nil {
    // handle error
}

key, err = dsutils.NameKeyWithNamespace(model.OptionKind, "namespace", "Name", nil)
if err != nil {
    // handle error
}

key, err = dsutils.IDKeyWithNamespace(model.OptionKind, "namespace", int64(Id), nil)
if err != nil {
    // handle error
}
```

### License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.
