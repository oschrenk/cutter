# README

## Features

- prints json be default
- fetches cookies from Safari, Arc and Chrome

## Limitations

- macOS only

## Usage

- `cutter list [--browser <name>] [--profile <id>]` **List cookies**
- `cutter profiles [--browser <name>]` **List profiles**

`--browser` accepts `safari` (default), `arc` or `chrome`. `--profile` takes the
ID from `cutter profiles`, not the display name. Arc and Chrome IDs are profile
directory names such as `Default` or `Profile 1`, so they need quoting.

### Keychain

Arc and Chrome encrypt cookie values with a key held in the login keychain, so
macOS asks for permission the first time cutter reads them. **Always Allow** is
safe here: the grant is tied to that one cutter binary and that one keychain
entry, and no other program can use it. Upgrading cutter changes the binary, so
macOS asks once more. Safari needs no keychain access.

### Examples

List all domains
```
cutter list | jq -r .[].domain | sort | uniq
```

List all keys for a given domain
```
cutter list | jq -r '.[] | select(.domain==".acme.com") | .name' | sort
```
Get a specific cookie value

```
cutter list | jq -r '.[] | select(.domain==".acme.com") | select(.name=="foo") | .value'
```

List all profile names
```
cutter profiles | jq -r .[].name
```

List cookies for a specific Safari profile
```
cutter list --profile 59869AEE-3D5E-4F8C-87DF-8554461D5221
```

List Arc profiles, then read one of them
```
cutter profiles --browser arc
cutter list --browser arc --profile "Profile 1"
```

Read Chrome cookies once, then filter repeatedly without a new keychain prompt
```
cutter list --browser chrome > cookies.json
jq -r '.[] | select(.domain==".github.com") | .name' cookies.json
```

## Installation

### nix

```bash
nix profile install github:oschrenk/cutter
```

Prebuilt binaries come from the `oschrenk` Cachix cache, which the flake offers
as a substituter.

### homebrew

```bash
brew tap oschrenk/made git@github.com:oschrenk/homebrew-made
brew install oschrenk/made/cutter
```

### From source

* installs to `$GOBIN/cutter`

```bash
git clone git@github.com:oschrenk/cutter.git
cd cutter
task install
```

