{
  stdenv,
  fetchzip,
}:

stdenv.mkDerivation {
  pname = "Golang";
  version = "1.24";

  src = fetchzip {
    url = "https://go.dev/dl/go1.24.0.linux-amd64.tar.gz";
    hash = "sha256-eICnZSd/aYOmUJ8HJqzSoQN1EIuU80GOa47W/7tOysM=";
  };

  phases = [
    "unpackPhase"
    "installPhase"
  ];

  installPhase = ''
    mkdir -p $out

    cp go/.* $out
  '';
}
