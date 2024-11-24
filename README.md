# Articola Tools' file path validator

[![image size](https://ghcr-badge.egpl.dev/articola-tools/file-path-validator/size?color=dodgerblue)](https://ghcr-badge.egpl.dev/articola-tools/file-path-validator/size?color=dodgerblue)

This repo contains Dockerfile with preconfigured file paths validator written in
Go. The validator allows validating a file, and all folders in the provided
path are stick to the same naming convention. Currently, validator supports two
naming conventions - `snake_case` and `PascalCase`.

This linter is used in Articola Tools organization's repositories to validate
new files in PRs.

## Usage

Use `ghcr.io/articola-tools/file-path-validator` Docker image with two flags:

- `--naming-convention` - `snake_case` or `PascalCase`
- `--path-to-validate` - a file path to validate

Example command to use this linter:

```bash
docker run --rm ghcr.io/articola-tools/file-path-validator \
--naming-convention snake_case \
--path-to-validate "./repo/main.go"
```
