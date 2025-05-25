{
  pkgs ? import <nixpkgs> { },
  go ? (
    import ./go.nix {
      stdenv = pkgs.stdenv;
      fetchzip = pkgs.fetchzip;
    }
  ),
}:
pkgs.mkShell {
  packages = [
    go
    pkgs.yaml-language-server
    pkgs.gopls
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
