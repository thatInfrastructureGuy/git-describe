# generate-version
Generates .go file containing binary version by running `git describe --abbrev=7 --dirty`. Meant to be used with `go generate` command. Replacement for `go build -X ldflag main.Version=${VERSION}`
