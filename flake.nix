{
  description = "Cutter - extract cookies from browsers";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  # Offer prebuilt binaries from the Cachix cache so `nix profile install`
  # downloads instead of compiling. Consumers are prompted to trust these.
  nixConfig = {
    extra-substituters = [ "https://oschrenk.cachix.org" ];
    extra-trusted-public-keys = [
      "oschrenk.cachix.org-1:3JOMfkq2vFiLw4UsCVwzu8kWFBkuS/3DD5AojcO9pks="
    ];
  };

  outputs =
    { self, nixpkgs }:
    let
      # Single source of truth for the version: ./VERSION holds a bare semver
      # (e.g. 0.2.0); the "v" prefix is added by the taskfile release flow.
      version = nixpkgs.lib.fileContents ./VERSION;

      # Darwin only: cutter reads Safari's container under
      # ~/Library/Containers/com.apple.Safari, so a linux build has nothing to
      # read. aarch64 only: it is what CI builds, so it is the only system the
      # binary cache is ever populated for.
      systems = [
        "aarch64-darwin"
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
    in
    {
      packages = forAllSystems (pkgs: rec {
        cutter = pkgs.buildGoModule {
          pname = "cutter";
          inherit version;
          src = self;

          # Regenerate after changing go.mod/go.sum: set to lib.fakeHash,
          # run `nix build`, then paste the expected hash from the error.
          vendorHash = "sha256-Xd4U0TbRrG+EYJ8+9JtZWX4h7HIYYiej9E4JHJsDRcc=";

          # cmd.version is what `cutter --version` prints; the taskfile stamps
          # the same "v"-prefixed value.
          ldflags = [
            "-s"
            "-w"
            "-X github.com/oschrenk/cutter/cmd.version=v${version}"
          ];

          # `completion` is hidden (cmd/root.go) but not disabled, and it never
          # reaches the config loader, so it is safe to run in the sandbox.
          nativeBuildInputs = [ pkgs.installShellFiles ];
          postInstall = ''
            installShellCompletion --cmd cutter \
              --bash <($out/bin/cutter completion bash) \
              --zsh <($out/bin/cutter completion zsh) \
              --fish <($out/bin/cutter completion fish)
          '';

          meta = {
            description = "Extract cookies from browsers";
            homepage = "https://github.com/oschrenk/cutter";
            mainProgram = "cutter";
            platforms = nixpkgs.lib.platforms.darwin;
          };
        };
        default = cutter;
      });

      apps = forAllSystems (pkgs: rec {
        cutter = {
          type = "app";
          program = "${self.packages.${pkgs.stdenv.hostPlatform.system}.cutter}/bin/cutter";
        };
        default = cutter;
      });

      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go # go, language
            golangci-lint # go, linter runner
            gopls # go, lsp
          ];
        };
      });
    };
}
