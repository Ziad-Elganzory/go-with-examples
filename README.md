# GO With Examples

Welcome to the **GO With Examples** project repository!

## Overview

This repository contains the source code and resources for the GO principles. It is designed to [briefly describe the purpose or functionality of GO LANG principles]. [Official Examples](https://gobyexample.com/)

## Table of Content

- 1.[Hello World](https://github.com/Ziad-Elganzory/go-with-examples/blob/main/1.Hello%20World/hello_world.go)
- 2.[Values](https://github.com/Ziad-Elganzory/go-with-examples/blob/main/2.Values/values.go)
- 3.[Variables](https://github.com/Ziad-Elganzory/go-with-examples/blob/main/3.Variables/variables.go)
- 4.[Constants](https://github.com/Ziad-Elganzory/go-with-examples/blob/main/4.Constants/constants.go)
- 5.[For Loop](https://github.com/Ziad-Elganzory/go-with-examples/blob/main/5.For%20Loop/for.go)
- 6.[If-Else](https://github.com/Ziad-Elganzory/go-with-examples/blob/main/6.If-Else/if-else.go)
- 7.[Switch](https://github.com/Ziad-Elganzory/go-with-examples/blob/main/7.Switch/switch.go)

## Getting Started

1. **Clone the repository:**
    ```bash
    git clone https://github.com/Ziad-Elganzory/go-with-examples.git
    ```
2. **Navigate to the project directory:**
    ```bash
    cd go-with-examples
    ```

## Installation

To install Go, follow the steps for your operating system:

### Windows

1. Download the Windows installer from the [official Go downloads page](https://go.dev/dl/).
2. Run the installer and follow the prompts.
3. After installation, open Command Prompt and verify the installation:
    ```bash
    go version
    ```

### Linux

1. Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:
    ```bash
    $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz
    ```
    (You may need to run each command separately with the necessary permissions, as root or through sudo.)

    **Do not** untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

2. Add /usr/local/go/bin to the PATH environment variable.

    You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```
    **Note**: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

3. Verify that you've installed Go by opening a command prompt and typing the following command:

    ```bash
    $ go version
    ```
4. Confirm that the command prints the installed version of Go.

    ```bash
    go version
    ```

## Usage

```bash
# Example command to run the project
go run main.go
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements.

## Socials:
- [Linkedin](https://www.linkedin.com/in/ziadmohamed770/)