# git-describe

* Runs `git describe --abbrev=7 --dirty`
* Generates .go file containing binary version.
* Alternative to `go build -ldflags "-X=xxx.Version=$(git describe --abbrev=7 --dirty)"` command.

## Prerequisites
* git

## How to use

```
go run github.com/thatInfrastructureGuy/git-describe@latest && go build <<your-program>>
```

Probably you should put it in your tools.go file.

### What happens on running this program?

```
go run github.com/thatInfrastructureGuy/git-describe@latest
```

A new file gets created `version/const.go` with following contents:

```
// THIS IS GENERATED CODE; DO NOT EDIT.
// Generated at 2022-06-25 14:58:06.865286 -0700 PDT m=+0.000388855
package version

const Version = "v0.0.1"
```

### FAQs

<details>
<summary>
  I get an ERROR and the file has UNKNOWN version? 
</summary>
  
  `const Version = "UNKNOWN"`
  1. Please make sure you have `git` binary installed.
  2. Check the output of `git describe --abbrev=7 --dirty` manually for your repository.
</details>

<details>
<summary>
  How do I change generated filename / variable name /package name ?
</summary> 
  
  ```
  go run github.com/thatInfrastructureGuy/git-describe@latest \
     --filepath=version/version.go --package=version --variable=VERSION
  ```
</details>

<details>
<summary>
  Why do this?
</summary>
  
  * I always had to look up `go build -ldflags` command.
  * I wanted the build to error out if I _forgot_ to generate version. Just make sure you using the `Version` variable somewhere in your program for it to error out.
</details>

<details>
<summary>
  Is this method better than using ldflags?
</summary>
  ¯\_(ツ)_/¯
</details>

<details>
<summary>
  I do not use `git`. Can I use this program?
</summary>
  No. PRs welcome. :)
</details>
