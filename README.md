# README

## Usage

**List available profiles**
```
cutter profiles --json
[
  {
    "label": "John",
    "browser": "bunde.id"
    "path": "/Users/oliver/Library/Application Support/Arc/User Data/Default"
  },
  ...
}
```
Read cookies

```
cutter read --path "/Users/oliver/Library/Application Support/Arc/User Data/Default"
```


## Installation

**Via Github**

* installs to `$GOBIN/cutter`

```
git clone git@github.com:oschrenk/cutter.git
cd cutter
task install
```
