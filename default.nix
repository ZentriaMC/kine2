{ lib, buildGoModule, runCommandNoCC, git, rev ? null }:

let
  versionInfo = src: import (runCommandNoCC "eteu-api-version" { } ''
    v=$(${git}/bin/git -C "${src}" rev-parse HEAD || echo "0000000000000000000000000000000000000000")
    printf '{ version = "%s"; }' "$v" > $out
  '');

  # Need to keep .git around for version string
  srcCleaner = name: type: let baseName = baseNameOf (toString name); in (baseName == ".git" || lib.cleanSourceFilter name type);
in
buildGoModule rec {
  pname = "kine2";
  version = if (rev != null) then rev else (versionInfo src).version;

  src = lib.cleanSourceWith { filter = srcCleaner; src = ./.; };

  doCheck = true;

  vendorSha256 = lib.fakeSha256;
  subPackages = [ "cmd/kine2" ];
}
