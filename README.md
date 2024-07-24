# README

## Features

- prints json be default
- fetches cookies from Safari

## Limitations

- currently only works with Safari

## Usage

- `cutter list` **List cookies**

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

