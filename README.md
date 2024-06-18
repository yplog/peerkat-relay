# Peerkat Relay

<p align="center">
    <img src="https://yalinpala.dev/projects/peerkat-relay.png" alt="peerkat logo"  width="250" height="250">
</p>

Peerkat Relay is a simple and lightweight relay server for [Peerkat](https://github.com/yplog/peerkat). It is designed to be used with Peerkat, a peer-to-peer file sending application.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Contribution](#contribution)
4. [License](#license)

## Installation

Clone the repository:

```bash
git clone git@github.com:yplog/peerkat-relay.git
```

Install the dependencies:

```bash
go mod tidy
```

To build Peerkat Relay, you can use the following command:

```bash
go build -o peerkat-relay cmd/relay-server/main.go
```

## Usage

To start the relay server, simply run the following command:

```bash
peerkat-relay
```

## Contribution

If you want to contribute to the development of Peerkat, please follow these steps:

1. Fork this repository.
2. Add a new feature or fix a bug.
3. Create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.