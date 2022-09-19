# Run linter with configuration file

```bash
go install github.com/mgechev/revive@latest
revive -config revive.toml
```

# LINTING IN A CODE EDITOR

it is easy to change the linter to revive using the preferences ➤ extensions ➤ go ➤ lint tool configuration option. if
you want to use a custom configuration file, use the lint Flags configuration option to add a flag with the value
-config=./revive.toml, which will select the revive.toml file.

# Fix common issue

```bash
go vet main.go

go vet -json main.go
# suppose assign is analyzed
go vet -assign=false
go vet -assign
```
