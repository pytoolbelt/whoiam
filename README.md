[![Sketch fonts](https://see.fontimg.com/api/rf5/1jX4/ZjY1MjIzYzUxMTQ2NGUwNDllOWNjMjRhOGRkMjQzNTcudHRm/d2hvaWFt/sketchy.png?r=fs&h=186&w=2000&fg=A91DD7&bg=FFFFFF&tb=1&s=93)](https://www.fontspace.com/category/sketch)

`whoiam` is a CLI that allow you to run a command for a specific AWS account. This safeguards you from running a command in the wrong AWS account
and deploying or destroying resources where you didn't intend to.

## Features

- Run command for a specific AWS account
- Retrieve AWS IAM Role information
- Supports multiple AWS accounts

## Documentation
For more information and usage examples, please refer to the [documentation](https://pytoolbelt.github.io/whoiam/).

## Installation

### Using Homebrew

```sh
brew tap pytoolbelt/homebrew-awstools
brew install whoiam
```

### Download Binary

You can download the pre-compiled binaries from the [releases page](https://github.com/pytoolbelt/whoiam/releases).

To download the pre-compiled binaries from the releases page using `curl`, you can use the following commands. Replace `VERSION`, `OS`, and `ARCH` with the appropriate values for the version, operating system, and architecture you need.

For example, to download the Linux binary for version `v1.0.0`:

```sh
curl -L -o whoiam_linux_x86_64.tar.gz https://github.com/pytoolbelt/whoiam/releases/download/v1.0.0/whoiam_Linux_x86_64.tar.gz
```

For the Windows binary:

```sh
curl -L -o whoiam_windows_x86_64.zip https://github.com/pytoolbelt/whoiam/releases/download/v1.0.0/whoiam_Windows_x86_64.zip
```

For the macOS binary:

```sh
curl -L -o whoiam_darwin_x86_64.tar.gz https://github.com/pytoolbelt/whoiam/releases/download/v1.0.0/whoiam_Darwin_x86_64.tar.gz
```

Make sure to replace `v1.0.0` with the actual version number you want to download.

### Build from Source

```sh
git clone https://github.com/pytoolbelt/whoiam.git
cd whoiam
go build -o whoiam
```

## Usage

```sh
whoiam --help
```

## Initialisation
A config file can be generated at the default location `~/.whoiam/whoiam.yaml` by running the following command:

```sh
whoiam config init
```

## Configuration

`whoiam` uses the AWS SDK for Go, so it will look for credentials and configuration in the default locations used by the AWS CLI and SDKs.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.

## Author

Jesse Maitland - [jesse@pytoolbelt.com](mailto:jesse@pytoolbelt.com)
```