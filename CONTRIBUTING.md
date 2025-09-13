# Contributing to geminik8s

First off, thank you for considering contributing to `geminik8s`! It's people like you that make open source great.

We welcome contributions of all kinds, from bug reports and feature requests to code contributions and documentation improvements.

## Ways to Contribute

*   **Reporting Bugs**: If you find a bug, please [create an issue](https://github.com/turtacn/geminik8s/issues) and provide as much detail as possible.
*   **Suggesting Enhancements**: If you have an idea for a new feature or an improvement to an existing one, please [create an issue](https://github.com/turtacn/geminik8s/issues) to discuss it.
*   **Writing Documentation**: We are always looking to improve our documentation. If you see something that is unclear or missing, please let us know.
*   **Contributing Code**: If you would like to contribute code, please read the "Development Setup" section below and then submit a pull request.

## Development Setup

### Prerequisites

*   Go 1.18+
*   Docker
*   `make`

### Getting Started

1.  **Fork the repository**:
    Click the "Fork" button at the top right of the [repository page](https://github.com/turtacn/geminik8s).

2.  **Clone your fork**:
    ```bash
    git clone https://github.com/YOUR_USERNAME/geminik8s.git
    cd geminik8s
    ```

3.  **Add the upstream remote**:
    ```bash
    git remote add upstream https://github.com/turtacn/geminik8s.git
    ```

4.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

### Building and Testing

*   **Build the binary**:
    ```bash
    make build
    ```
    This will create the `gemin_k8s` binary in the `bin/` directory.

*   **Run unit tests**:
    ```bash
    make test-unit
    ```

*   **Run integration tests**:
    ```bash
    make test-integration
    ```

*   **Run E2E tests**:
    ```bash
    make test-e2e
    ```

*   **Run all tests**:
    ```bash
    make test
    ```

### Code Style

We use `golangci-lint` for linting and `gofumpt` for formatting. Before submitting a pull request, please make sure your code is linted and formatted correctly.

*   **Format the code**:
    ```bash
    make format
    ```

*   **Lint the code**:
    ```bash
    make lint
    ```

### Submitting a Pull Request

1.  Create a new branch for your changes:
    ```bash
    git checkout -b my-feature-branch
    ```

2.  Make your changes.

3.  Commit your changes with a descriptive commit message:
    ```bash
    git commit -m "feat: add new feature"
    ```
    We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification.

4.  Push your changes to your fork:
    ```bash
    git push origin my-feature-branch
    ```

5.  Create a pull request from your fork to the `main` branch of the `turtacn/geminik8s` repository.

## Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms. (We will add a CODE_OF_CONDUCT.md file later).
