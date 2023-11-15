# Teleporter CLI

This directory contains the source code for the Teleporter CLI. The CLI is a command line interface for interacting with the Teleporter contracts. It is written with [cobra](https://github.com/spf13/cobra) commands as a Go application.

## Building

To build the CLI, run `go build -o teleporter-cli` from this directory. This will create a binary called `teleporter-cli` in the current directory.

## Usage

The CLI has a number of subcommands. To see the list of subcommands, run `./teleporter-cli help`. To see the help for a specific subcommand, run `./teleporter-cli help <subcommand>`.

The supported subcommands include:

- `event`: given a log event's topics and data, attempts to decode into a Teleporter event in more readable format.
- `message`: given a Teleporter message bytes hex encoded, attempts to decode into a Teleporter message in more readable format.
- `transaction`: given a transaction hash, attempts to decode all relevant Teleporter and Warp log events in more readable format.