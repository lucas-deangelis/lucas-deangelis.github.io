{ pkgs ? import <nixpkgs> { } }:

let 
  customGopls = import "${builtins.getEnv "HOME"}/nixos/packages/gopls/default.nix" {
    inherit (pkgs) lib buildGo122Module fetchFromGitHub;
  };
in
pkgs.mkShell {
  buildInputs = [
    pkgs.janet
  ];
}
