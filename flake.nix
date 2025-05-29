{
  description = "Go dev env";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      gomod2nix,
    }:
    let
      forAllSystems = nixpkgs.lib.genAttrs nixpkgs.lib.platforms.unix;

      nixpkgsFor = forAllSystems (
        system:
        import nixpkgs {
          inherit system;
          config = { };
          overlays = [ ];
        }
      );
    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor."${system}";
          # pandoc = pkgs.pandoc;
          buildGoApplication = gomod2nix.legacyPackages.${system}.buildGoApplication;
        in
        {
          default = buildGoApplication {
            # Required args.
            name = "ledsproj";
            src = ./.;

            # Override default Go with Go 1.21.
            #
            # In the latest versions of Go, the go.mod can contain 1.21.5
            # In that case, if the toolchain doesn't match, the go build operation will
            # try and download the correct toolchain.
            #
            # To prevent this, update the go.mod file to contain `go 1.21` instead of `go 1.21.5`.
            go = (
              import ./go.nix {
                stdenv = pkgs.stdenv;
                fetchzip = pkgs.fetchzip;
              }
            );

            # Must be added due to bug https://github.com/nix-community/gomod2nix/issues/120
            pwd = ./.;

            # Optional flags.
            CGO_ENABLED = 0;
            flags = [ "-trimpath" ];
            ldflags = [
              "-s"
              "-w"
              "-extldflags -static"
            ];
          };
          # inherit go;
          # inherit pandoc;
        }
      );

      devShells = forAllSystems (
        system:
        let
          pkgs = nixpkgsFor."${system}";
          go = (
            import ./go.nix {
              stdenv = pkgs.stdenv;
              fetchzip = pkgs.fetchzip;
            }
          );
        in
        {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              gotools
              gomod2nix.packages.${system}.default # gomod2nix CLI
              sbomnix.packages.${system}.default # sbomnix CLI
            ];
          };
        }
      );
    };
}
