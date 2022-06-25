# generate-version

* Generates .go file containing binary version by running `git describe --abbrev=7 --dirty`.
* Alternative to `go build -ldflags "-X=xxx.Version=${VERSION}"` command.

## How to use

```
go run github.com/thatInfrastructureGuy/generate-version@latest && go build <<your-program>>
```

Probably you should put it in your tools.go file.

### What happens on running this program?

```
go run github.com/thatInfrastructureGuy/generate-version@latest
```

A new file gets created `version/const.go` with following contents:

```
// THIS IS GENERATED CODE; DO NOT EDIT.
// Generated at 2022-06-25 14:58:06.865286 -0700 PDT m=+0.000388855
package version

const Version = "v0.0.1"
```

### FAQs

<detail>
<summary>
I get an ERROR and the file has UNKNOWN version:  `const Version = "UNKNOWN"`
</summary>
1. Please make sure you have `git` binary installed.
2. Check the output of `git describe --abbrev=7 --dirty` manually for your repository.
</detail>

<detail>
<summary>
How do I change generated filename/variable name/package name?
</summary>
```
go run github.com/thatInfrastructureGuy/generate-version@latest --filepath=version/version.go --package=version --variable=VERSION
```
</detail>

<detail>
<summary>
Why do this?
</summary>
* I always had to look up `go build -ldflags` command.
* I wanted the build to error out if I _forgot_ to generate version. Now, my program fails to build if I forget to generate version, because somewhere in there it is looking for a const `Version` which is nowhere defined. Just make sure you using the `Version` variable somewhere in your program for it to error out.
</detail>

<detail>
<summary>
Is this method better than using ldflags?
</summary>
I am not sure.
</detail>

<detail>
<summary>
I do not use `git`. Can I use this program?
</summary>
No. PRs welcome. :)
</detail>
