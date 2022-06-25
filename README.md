# generate-version

* Generates .go file containing binary version by running `git describe --abbrev=7 --dirty`.
* Meant to be used with `go generate` command. 
* Alternative to `go build -ldflags "-X=xxx.Version=${VERSION}"` command.
