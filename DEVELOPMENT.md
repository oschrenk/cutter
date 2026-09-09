# Development

**Requirements**

The flake ships a devShell with go, gopls and golangci-lint. With
[direnv](https://direnv.net/) installed, `direnv allow` puts them on the path.
Without it, use `nix develop`.

Tools the devShell does not provide:

- [air](https://github.com/cosmtrek/air) `go install github.com/air-verse/air@latest`
- [staticcheck]() `go install honnef.co/go/tools/cmd/staticcheck@latest`

## Dependencies

kooky is pinned to the head commit of
[browserutils/kooky#108](https://github.com/browserutils/kooky/pull/108), which
adds `chromium.KeyringConfigArc`. Arc and Chrome hold their cookie keys under
their own keychain entries, and the released kooky hardcodes "Chrome Safe
Storage", so no tagged version can read Arc. Drop the pin once #108 ships in a
release.

After changing `go.mod` or `go.sum`, `vendorHash` in `flake.nix` no longer
matches. Set it to `nixpkgs.lib.fakeHash`, run `nix build`, and paste the
expected hash from the error.

## Tasks

- `task build` Build project
- `task run` Run example
- `task test` Run tests
- `task lint` Lint
- `task install` Install app in `$GOBIN/`
- `task uninstall` Removed app from `$GOBIN/`
- `task artifacts` Produces artifact in `./`
- `task tag` Pushes git tag from `VERSION`
- `task release` Creates GitHub release from artifacts
- `task sha` Prints hashes from artifacts
- `task clean` Removes build directory `.build`
- `task updates` Find dependency updates

## Release

1. Increase version number in `VERSION`
2. `task release` to tag and push
3. `task sha` to print hashes to `stdout`
4. Make changes in [homebrew-made](https://github.com/oschrenk/homebrew-made) and push
5. `brew update` to update taps
6. `brew upgrade` to upgrade formula
