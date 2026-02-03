# README

## Features

- prints json be default
- fetches cookies from Safari

## Limitations

- currently only works with Safari

## Usage

- `cutter list [--profile <id>]` **List cookies**
- `cutter profiles` **List Safari profiles**

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

List cookies for a specific profile
```
cutter list --profile 59869AEE-3D5E-4F8C-87DF-8554461D5221
```

## Installation

**Via Github**

* installs to `$GOBIN/cutter`

```
git clone git@github.com:oschrenk/cutter.git
cd cutter
task install
```

**Via homebrew**

```
brew tap oschrenk/made git@github.com:oschrenk/homebrew-made
brew install oschrenk/made/cutter
```

