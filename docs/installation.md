# Installation Guide

This document provides comprehensive instructions for installing `geminik8s`.

## Prerequisites

- Go 1.18+
- Docker (for building and testing)
- Access to two nodes that will form the cluster

## Installation from Source

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/turtacn/geminik8s.git
    cd geminik8s
    ```

2.  **Build the binary:**
    ```bash
    make build
    ```
    The binary `gemin_k8s` will be created in the `bin/` directory.

3.  **Install the binary:**
    You can copy the binary to a directory in your `PATH`, for example:
    ```bash
    sudo cp bin/gemin_k8s /usr/local/bin/
    ```

## Installation from Release

Pre-built binaries for Linux, macOS, and Windows are available on the [GitHub Releases](https://github.com/turtacn/geminik8s/releases) page.

1.  Download the appropriate archive for your operating system and architecture.
2.  Extract the archive.
3.  Move the `gemin_k8s` binary to a directory in your `PATH`.

## Next Steps

Once `geminik8s` is installed, you can proceed to the [Configuration Guide](./configuration.md) to set up your first cluster.
