{
  description = "NixOS flake for the Hestia project.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable"; # or a specific version
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
        overlays = [
          # Add any overlays if necessary
        ];
      };
    in {
      packages.default = pkgs.mkShell {
        buildInputs = [
          pkgs.flutter
          pkgs.dart
          pkgs.go
          pkgs.vim
        ];

        shellHook = ''
          echo "Welcome to the Hestia project environment!"
        '';
      };

      # NixOS configuration
      nixosConfigurations.hestia = pkgs.nixosSystem {
        system = "x86_64-linux"; # Change if needed
        modules = [
          {
            environment.systemPackages = with pkgs; [
              flutter
              dart
              go
              vim
            ];

            # Custom Vim configuration
            programs.vim = {
              enable = true;
              extraConfig = ''
                set number
                syntax on
                set tabstop=4
                set shiftwidth=4
                set expandtab
              '';
            };
          }
        ];
      };
    });
}
