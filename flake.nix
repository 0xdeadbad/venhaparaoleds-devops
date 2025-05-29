{
  description = "Go dev env";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs =
    {
      self,
      nixpkgs,
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
          go = (
            import ./go.nix {
              stdenv = pkgs.stdenv;
              fetchzip = pkgs.fetchzip;
            }
          );
          pandoc = pkgs.pandoc;
          default = (
            import ./default.nix {
              inherit pkgs;
              inherit go;
            }
          );
        in
        {
          inherit default;
          inherit go;
          inherit pandoc;
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
          default = import ./shell.nix {
            inherit pkgs;
            inherit go;
          };
        }
      );
    };
}
