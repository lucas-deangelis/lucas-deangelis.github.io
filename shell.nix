{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/614b4613980a522ba49f0d194531beddbb7220d3.tar.gz" ) { } }:

pkgs.mkShell {
  packages = [
    pkgs.janet
  ];
}
