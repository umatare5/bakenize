# bakenize

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/umatare5/everlasting-hey-yo-http/blob/master/LICENSE

A simple CLI to generate a token for the Authorization header from a username and password.

"bakenize" is short for "basic-auth-tokenizer".

## Installation

```bash
docker run ghcr.io/umatare5/bakenize
```

or use the binary that can be downloaded from [release page](https://github.com/umatare5/bakenize/releases).

## Usage

It simply sets the username and password.

```bash
$ bakenize --username postman --password "$secret"
cG9zdG1hbjpwYXNzd29yZA==
```

Use a generated token:

```bash
curl -H "Authorization: Basic cG9zdG1hbjpwYXNzd29yZA==" https://postman-echo.com/basic-auth
```

## Supported Environment Variables

The application supports the following environment variables:

| Name         | Description                |
| ------------ | -------------------------- |
| `B_USERNAME` | The alias of `--username`. |
| `B_PASSWORD` | The alias of `--password`. |

## Release

Set and push a tag.

```shell
git tag vX.Y.Z && git push --tags
```

Run the release workflow.

- [GitHub Actions: release workflow](https://github.com/umatare5/bakenize/actions/workflows/release.yaml)

## Contribution

1. Fork ([https://github.com/umatare5/bakenize/fork](https://github.com/umatare5/bakenize/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Create a new Pull Request

## Author

[umatare5](https://github.com/umatare5)
