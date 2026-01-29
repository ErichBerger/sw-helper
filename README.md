# SW Helper
Author: Erich Berger<br/>
Latest version: [0.1.3](https://github.com/ErichBerger/sw-helper/releases/tag/v0.1.3)

## About
This project is meant to help in Shopware development by creating the boilerplate files

## Install
### Build from source
- Make sure [Go](https://go.dev/dl/) is installed
- Clone the repo
- `cd` into the repo's root
- Run `go build . ` to compile binary, or `go run .` to run once

### Download prebuilt binary
[Releases](https://github.com/ErichBerger/sw-helper/releases)

### Homebrew
- `brew tap erichberger/tap`
- `brew install --cask sw-helper`
- `xattr -dr com.apple.quarantine "$(which sw-helper)"`

## Future Plans
This is a work in progress.

Current features:
- Generate Storefront JS Plugin
- Generate CMS Element

Future features:
- Generate CMS Block BoilerPlate
- Admin module?

## Requirements
- `sw-helper.toml` needs to be in same location as binary to be picked up
    - May change to support more 'config' like behavior like .bashrc or .vimrc
- Run `sw-helper` command from shopware project root