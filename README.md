> Archived in favor of golang/go#50603

# git-describe

* Runs `git describe --abbrev=7 --dirty`
* Generates .go file containing binary version.
* Alternative to `go build -ldflags "-X=xxx.Version=$(git describe --abbrev=7 --dirty)"` command.
* Ability to customize command via `--command` flag.

## Prerequisites
* git (if using default command)

## How to use

```
go run github.com/thatInfrastructureGuy/git-describe@latest \
   && go build <<your-program-with-flags>>
```

Please see [How can I track tool dependencies for a module?](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)

## What happens on running this program?

```
go run github.com/thatInfrastructureGuy/git-describe@latest
```

A new file gets created `version/const.go` with following contents:

```
// THIS IS GENERATED CODE; DO NOT EDIT.
// Generated at 2022-06-25 14:58:06.865286 -0700 PDT m=+0.000388855
package version

// Version is generated by running 'git describe --abbrev=7 --dirty'
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
     --filepath=version/version.go --package=version --variable=Version
  ```
</details>

<details>
<summary>
  Why?
</summary>

  * I wanted the build to error out if I _forgot_ to generate version. Just make sure you using the `Version` variable somewhere in your program for it to error out.

</details>

<details>
<summary>
 Go 1.18 already has build information.
</summary>
 
| Git-Describe Features | Go 1.18  | 
|-----------------------|----------|
|Git SHA| ✅ | 
|Git Tag| ❌ Works with `go install`.`go build` returns `(devel)`. https://github.com/golang/go/issues/50603 | 
|Number of commits since last tag| ❌ |
|Dirty Flag| ✅ |

</details>

<details>
<summary>
  Is this method better than using ldflags?
</summary>
  ¯\_(ツ)_/¯
</details>
