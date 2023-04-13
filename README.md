# gpt-cli

`gpt-cli` is a command line interface (CLI) application that leverages OpenAI's ChatGPT to suggest shell commands based on natural language queries. This tool aims to improve productivity and make it easier for users to interact with the shell, especially those who are not well-versed in shell scripting or command line utilities.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Features

- Use natural language queries to get shell command suggestions.
- Supports various popular shells, such as Bash, Zsh, and Sh.
- Works on multiple platforms, including Windows, macOS, and Linux.
- Interactive dialogue mode with ChatGPT


## Prerequisites

- Go 1.20 or higher
- API key for OpenAI's ChatGPT service

## Installation

```sh
go install github.com/alex-ello/gpt-cli/cmd/gpt-cli@latest
```

## Usage

To use gpt-cli, simply enter a natural language query describing the command you want to execute:

```sh
gpt-cli Find all text files in the current directory
```

The application will then display a list of suggested shell commands that match your query:

```sh
find . -name "*.txt"
```

To enter the interactive dialogue mode with ChatGPT, just run gpt-cli without parameters.


## For a Better Experience

To make using `gpt-cli` even more convenient, you can create an alias for the command. For example, you can set up an alias called `ai` to quickly access the tool.

For Linux/macOS, add the following line to your `.bashrc` or `.zshrc` file:

```sh
alias ai="gpt-cli"
```

## License

gpt-cli is released under the [MIT License](LICENSE).