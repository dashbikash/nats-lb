{ pkgs }: {
    deps = [
        pkgs.poetry
        pkgs.python39Full
        pkgs.wget
        pkgs.links
        pkgs.unzip
        pkgs.go_1_17
        pkgs.gopls
    ];
}