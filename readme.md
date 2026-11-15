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

### Releases

This repository is a Go library; releases are versioned Git tags and GitHub
releases, without binary artifacts. To request a new stable release, change
`.github/release-request.txt` to the next version in a pull request. Merging
that change into `master` runs Go module verification, race-enabled tests,
and `go vet` before publishing the tag and generated release notes.

The existing CI/CD workflow also supports manually dispatched patch, minor,
and major releases.

### License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.
