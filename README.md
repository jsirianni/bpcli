# bpcli

Lightweight command line utility for interacting with the BindPlane public API

* [bindplane.bluemedora.com](https://bindplane.bluemedora.com)
* [Bindplane Getting Started](https://docs.bindplane.bluemedora.com/docs/getting-started)
* [Bindplane API Documentation](https://docs.bindplane.bluemedora.com/reference#introduction)

[![Build Status](https://travis-ci.com/BlueMedoraPublic/bpcli.svg?branch=master)](https://travis-ci.com/BlueMedoraPublic/bpcli)
[![Go Report Card](https://goreportcard.com/badge/github.com/BlueMedoraPublic/bpcli)](https://goreportcard.com/report/github.com/BlueMedoraPublic/bpcli)


## Cloudshell Tutorial

[![Intro to bpcli in Cloud Shell](https://gstatic.com/cloudssh/images/open-btn.png)](https://console.cloud.google.com/cloudshell/open?git_repo=https://github.com/BlueMedoraPublic/bpcli&shellonly=true&tutorial=doc/google-cloud-tutorial/bpcli-intro.md)

## Install
Download the release zip for your platform, unzip, and place the
binary in your path:
```
unzip bpcli_darwin_amd64.zip
chmod +x bpcli
mv bpcli /usr/local/bin
```

## Usage
bpcli uses [cobra](https://github.com/spf13/cobra) for managing
commands and flags.

See `/doc` for usage examples.

#### Help
All commands have a built in help flag. `--help` can be passed at any time.
```
bpcli --help
```

#### Sources
```
bpcli source
bpcli source create
bpcli source delete
bpcli source get
bpcli source list
bpcli source type
bpcli source type get
bpcli source type list
bpcli source type template
```
#### Credentials
```
bpcli credential
bpcli credential create
bpcli credential delete
bpcli credential get
bpcli credential list
bpcli credential type
bpcli credential type get
bpcli credential type template
```
#### Collectors
```
bpcli collector get
bpcli collector list
bpcli collector delete
bpcli collector group get
bpcli collector group list
```
#### Jobs
````
bpcli job list
bpcli job get
````
#### Accounts
```
bpcli account list
bpcli account add
bpcli account set
bpcli account remove
```

## Credentials File
bpcli allows the user to change the account they are manipulating by utilizing
a credentials file.

*Warning*
The **BINDPLANE_API_KEY** environment variable will always take precedence over the
credentials file, if it is present.

#### Example Usage
Adding and setting an account to be used by bpcli
The `set` command makes the given account the **Active** account and is
required even when only one entry is present in the credentials file
```
bpcli account add --name=<ACCOUNT NAME> --id=<API_KEY>
bpcli account set --name=<ACCOUNT_NAME>
```

List all accounts that have been added to the configuration file
```
bpcli account list
```

Remove an account from the configuration file
```
bpcli account remove --name=<ACCOUNT_NAME>
```

## Shell Completion

#### Bash

bash-completion v2 requires bash version 4+
On MacOS, the default version is below 4 and will need to be updated!
Follow these instructions on [Upgrading Bash on MacOS](https://itnext.io/upgrading-bash-on-macos-7138bd1066ba).

To setup bash completion for bpcli on MacOS:
1. Install *bash-completion* by running `brew install bash-completion@2`&nbsp;
2. Include the following lines in `~/.bash_profile`&nbsp;
```
export BASH_COMPLETION_COMPAT_DIR="/usr/local/etc/bash_completion.d"
[[ -r "/usr/local/etc/profile.d/bash_completion.sh" ]] && . "/usr/local/etc/profile.d/bash_completion.sh"
```
3. Run the following command to include the bash-completion script in `/usr/local/etc/bash_completion.d/`\
`bpcli completion >/usr/local/etc/bash_completion.d/bpcli`
4. Restart the shell and bpcli tab completions will be available

#### ZSH

To setup zsh completion for bpcli on MacOS:
1. Include the following lines in `~/.zshrc`&nbsp;
```
autoload -Uz compinit
compinit
```
2. Locate `fpath` by running `echo $fpath`
3. Run the following command to generate the zsh tab completion script.\
`bpcli completion --zsh ><YOUR FPATH HERE>/_bpcli`
4. Restart zsh and the bpcli tab completions will be available.

## Developing

#### Lint

When making changes please write tests and lint your code with `golint`:
```
make lint
```

#### Test

Tests require an api key to be set, even if it is invalid
```
export BINDPLANE_API_KEY='somefakekey'
make test
```

Tests can be run against the live Bindplane api, however, a valid
Bindplane environment and api key must be present
```
export BINDPLANE_API_KEY='somerealkey'
export BINDPLANE_LIVE_TEST=1
make test
```

#### Alternative API

If you wish to target an API other than the default ("https://public-api.bindplane.bluemedora.com/v1"),
set the following environment variables:
```
export BINDPLANE_API_ENDPOINT="https://public-api.gcloud.bindplane.bluemedora.com"
export BINDPLANE_API_VERSION="/v1"
```

`BINDPLANE_API_ENDPOINT` and `BINDPLANE_API_VERSION` can be set independently

#### Build

The `Makefile` in this repository will use Docker to build
`bpcli`. This is to ensure a consistent build environment.
Compiling with Docker is optional.

Build with Docker, and check the artifacts directory when finished
- linux zip archive
- darwin zip archive
- windows zip archive
- SHA256 sum file
```
make
```

To build on your own, without Docker, clone this repo *outside* of your GOPATH, as
bpcli uses go modules:
```
env CGO_ENABLED=0 go build -a
```

To cross compile on your own system, without Docker, set
`GOOS`, and `GOARCH`:
```
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get ./...
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

## Third party packages

bpcli relies on:
- [cobra](https://github.com/spf13/cobra) for command line flags
