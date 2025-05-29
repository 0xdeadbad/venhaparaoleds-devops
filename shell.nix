{
  pkgs ? import <nixpkgs> { },
  go ? (
    import ./go.nix {
      stdenv = pkgs.stdenv;
      fetchzip = pkgs.fetchzip;
    }
  ),
}:
let
  ledsproj = (
    import ./default.nix {
      inherit pkgs;
    }
  );
in
pkgs.mkShell {
  packages = [
    pkgs.yaml-language-server
    pkgs.gopls
    ledsproj
    go
  ];

  shellHook = ''
    go get github.com/nsf/gocode
    go get github.com/tpng/gopkgs
    go get github.com/ramya-rao-a/go-outline
    go get honnef.co/go/tools/staticcheck
    go get golang.org/x/tools/cmd/guru
    # go install golang.org/x/tools/gopls@latest
  '';
}
