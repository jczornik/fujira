{
  description = "A Nix-flake-based go development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self , nixpkgs ,... }: let
    system = "x86_64-linux";
  in {
    devShells."${system}".default = let
      pkgs = import nixpkgs {
        inherit system;
      };
    in pkgs.mkShell {
      hardeningDisable = [ "fortify" ];
      packages = with pkgs; [
        go
        gopls
      ];

      shellHook = ''
      	PATH=$PATH:~/go/bin
      '';
    };
  };
}
