{
  description = "";

  inputs = {
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages."${system}";
      lib = pkgs.lib;
    in rec {
      # `nix build`
      packages.telegram = pkgs.buildGoModule rec {
        name = "telegram-bot";
        src = ./telegram-bot;

        vendorSha256 = "sha256-NBY2T0Jj5f9XMoc0hleituKHFf+9AFq056a/EpCsuXo=";
      };
      defaultPackage = packages.telegram;

      # `nix run`
      apps.telegram = utils.lib.mkApp {
        drv = packages.telegram;
      };
      defaultApp = apps.telegram;

      # `nix develop`
      devShell = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [ go ];
      };
    });
}
