#!/bin/bash

set -e

REPO="tuoju/whros-cli"
BINARY_NAME="whros"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
VERSION="latest"

parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -d|--dir)
                INSTALL_DIR="$2"
                shift 2
                ;;
            -v|--version)
                VERSION="$2"
                shift 2
                ;;
            -h|--help)
                echo "Usage: install.sh [OPTIONS]"
                echo ""
                echo "Options:"
                echo "  -d, --dir DIR     Install to directory (default: /usr/local/bin)"
                echo "  -v, --version VER Install specific version (default: latest)"
                echo "  -h, --help        Show this help message"
                exit 0
                ;;
            *)
                echo "Unknown option: $1"
                exit 1
                ;;
        esac
    done
}

get_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)

    case $arch in
        x86_64) arch="amd64" ;;
        aarch64|arm64) arch="arm64" ;;
        armv7l) arch="armv7" ;;
        *)
            echo "Unsupported architecture: $arch" >&2
            exit 1
            ;;
    esac

    case $os in
        linux|darwin|windows)
            ;;
        msys*|mingw*|cygwin*)
            os="windows"
            ;;
        *)
            echo "Unsupported OS: $os" >&2
            exit 1
            ;;
    esac

    echo "$os:$arch"
}

download() {
    local version=$1
    local platform=$2
    local os="${platform%:*}"
    local arch="${platform#*:}"
    local ext=""
    local url

    [ "$os" = "windows" ] && ext=".exe"

    if [ "$version" = "latest" ]; then
        local api_url="https://gitee.com/api/v5/repos/${REPO}/releases/latest"
        version=$(curl -s "$api_url" | grep -o '"tag_name":[^,]*' | sed 's/"tag_name": "//;s/",//' | sed 's/^v//')
    fi

    local filename="${BINARY_NAME}-${os}-${arch}${ext}"
    url="https://gitee.com/${REPO}/releases/download/v${version}/${filename}"

    echo "Downloading $filename..."
    echo "URL: $url"

    local tmpfile=$(mktemp)
    trap "rm -f $tmpfile" EXIT

    if ! curl -L -o "$tmpfile" "$url" 2>/dev/null; then
        echo "Failed to download. Please check if version v${version} exists."
        echo "Create release at: https://gitee.com/${REPO}/releases"
        exit 1
    fi

    local dest="${INSTALL_DIR}/${BINARY_NAME}${ext}"
    cp "$tmpfile" "$dest"

    if [ "$os" != "windows" ]; then
        chmod +x "$dest"
    fi

    echo "Installed to: $dest"
    echo "Done!"
}

main() {
    parse_args "$@"
    local platform=$(get_platform)

    echo "Installing whros-cli v${VERSION} for ${platform}..."

    if [ ! -d "$INSTALL_DIR" ]; then
        echo "Error: Install directory does not exist: $INSTALL_DIR" >&2
        exit 1
    fi

    if [ ! -w "$INSTALL_DIR" ]; then
        echo "Error: No write permission to: $INSTALL_DIR" >&2
        echo "Try running with sudo: sudo $0 --dir $INSTALL_DIR"
        exit 1
    fi

    download "$VERSION" "$platform"
}

main "$@"
