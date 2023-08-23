# gitcloak

`gitcloak` **BETA release** : Cloak (security) for selected files on Git. 

![gitcloak](images/gitcloak0.png "gitcloak")

## Introduction

`gitcloak` will help to securely place secretive or confidential information in 
git. `gitcloak` is a `git` plugin created using `golang` and `git`.

## Features

`gitcloak` allows to get most features of git like:-
* Line changes/updates
* Addition or deletion of Lines (when no line-random)
* this will provide enough idea about configuration change with respect to code changes.
* i.e. this tool is binding code with configuration without using any third party tool.
* Can be used as secret store across devices, synced via github/gitlab/bitbucket
* Can be used as secret Markdown document data store, e.g. Obsidian
* Encryption Algorithms can be different for different files based on rules
* Every line can be randomized with `--line-random` flag, 
  * more secure when several lines in files are matching, like
    * empty lines
    * `{`, `}` curly braces in C/C++/Java/JSON files


## ToDo

The following things are remaining to be built, most of these will be available in `BETA` release.
* `git commit hook` to be injected, so that `gitcloak encrypt` will be always run before the actual git commit. This will ensure that all the required files are encrypted before getting committed. **Done**
* `git reset` should also rollback `.gitcloak/config.yaml` so the decryption `git decrypt` should work well [*Time Travel*].
* Make Installable on MacOSX, *BSD, Windows OS

## Installation

* Currently, testing on Linux only, hence, installation on Linux can be done as given below:-
```
make build
make install
```
* Assuming `go` is installed.

## Examples

* [Initialization and Single File Encryption](examples/single_file.md)
* [Initialization and Multiple Files (Regex) Encryption](examples/multi_files.md)

## Acknowledgement

I would like to dedicate this package to my Friends Venkatesh Pitta 
and Puneet Vyas, without their inspiration and support, this would not be possible. 
Also, I would like to thank God, my Teachers, my Parents, my Wife and Daughter 
to stand with me all the times.

## Contributing


Whether reporting bugs, discussing improvements and new ideas or writing
extensions: Contributions to `gitcloak` are welcome! Here\'s how to get
started:

1.  Check for open issues or open a fresh issue to start a discussion
    around a feature idea or a bug
2.  Fork [the repository](https://github.com/miteshbsjat/gitcloak/) on
    Github, create a new branch off the `master` branch and start making
    your changes (known as [GitHub
    Flow](https://guides.github.com/introduction/flow/index.html))
3.  Write a test which shows that the bug was fixed or that the feature
    works as expected
4.  Send a pull request and bug the maintainer until it gets merged and
    published
