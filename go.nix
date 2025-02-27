{
  stdenv,
  fetchzip,
}:

stdenv.mkDerivation {
  pname = "go";
  version = "1.24";

  src = fetchzip {
    url = "https://go.dev/dl/go1.24.0.linux-amd64.tar.gz";
    hash = "sha256-XVGiZCb1ugqNgnrtZBGrJeiCw5dvfluImcVpnnmZEVI=";
  };

  phases = [
    "unpackPhase"
    "installPhase"
  ];

  installPhase = ''
    mkdir -p $out

    for f in $(ls); do
        cp -r "$f" $out;
    done
  '';
}
