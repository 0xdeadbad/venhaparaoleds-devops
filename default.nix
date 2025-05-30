{
  pkgs ? import <nixpkgs> { },
}:
let
  go = (
    import ./go.nix {
      stdenv = pkgs.stdenv;
      fetchzip = pkgs.fetchzip;
    }
  );
in
pkgs.stdenv.mkDerivation {
  name = "ledsproj";
  src = ./.;

  # unpackPhase = ''
  #   for srcFile in $src; do
  #       cp -r $srcFile $(stripHash $srcFile)
  #   done
  # '';

  nativeBuildInputs = [
    go
  ];

  preConfigure = ''
    export GOCACHE=$TMPDIR/go-cache
    export GOPATH="$TMPDIR/go"
    export CGO=0

    mkdir -p $GOCACHE
  '';

  # configurePhase = ''
  #   go mod tidy
  # '';

  preBuild = ''
    export HOME=$TMPDIR
  '';

  buildPhase = ''
    go build -mod=vendor -ldflags "-s -w" -o $TMPDIR/$name .
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp $TMPDIR/$name $out/bin
  '';
}
