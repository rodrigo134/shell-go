# Shell-Go

A lightweight command-line shell implementation written in Go. This project demonstrates fundamental shell architecture including command parsing, built-in command execution, external command invocation, and output redirection.

## Features

- **Interactive Shell Prompt**: Real-time command line interface with input handling
- **Built-in Commands**: Core shell commands implemented natively
  - `echo`: Output text arguments
  - `cd`: Change current directory
  - `pwd`: Print working directory
  - `exit`: Terminate the shell
  - `type`: Display command type information
  - `clear`: Clear the terminal screen
- **External Command Execution**: Execute system binaries with full argument support
- **Output Redirection**: Redirect command output to files using `>`
- **Command Parsing**: Advanced tokenization with quote handling and escape sequences
- **Standard I/O Handling**: Complete stdin/stdout/stderr integration

## Requirements

- Go 1.26.2 or later

## Installation

Clone the repository and build the project:

```bash
git clone https://github.com/rodrigo134/shell-go.git
cd shell-go
go build -o shell main.go
```

## Usage

Run the shell executable:

```bash
./shell
```

The shell will display a prompt (`$ `) and await user input. Enter commands as you would in any standard shell:

```bash
$ echo Hello, World!
Hello, World!

$ pwd
/home/user/projects

$ cd /tmp
$ pwd
/tmp

$ ls -la
total 24
drwxr-xr-x  3 user user 4096 May 11 10:00 .
drwxr-xr-x 10 user user 4096 May 10 15:30 ..
-rw-r--r--  1 user user  123 May 11 10:00 main.go

$ echo "Hello World" > output.txt
$ cat output.txt
Hello World

$ type echo
echo is a shell builtin

$ type ls
ls is /usr/bin/ls

$ exit
```

## Built-in Commands

| Command | Description | Example |
|---------|-------------|---------|
| `echo` | Print arguments to standard output | `echo Hello World` |
| `cd` | Change current directory | `cd /tmp` |
| `pwd` | Print the current working directory | `pwd` |
| `type` | Display whether a command is a builtin or external command path | `type ls` |
| `clear` | Clear the terminal screen | `clear` |
| `exit` | Terminate the shell | `exit` |

## Output Redirection

The shell supports output redirection using the `>` operator:

```bash
$ echo "Hello World" > output.txt
$ pwd > current_dir.txt
$ ls -la > directory_listing.txt
```

Note: Append redirection (`>>`) is not currently supported.

## Project Structure

```
.
├── main.go           # Main shell implementation
├── shell             # Compiled binary (generated)
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── README.md         # Project documentation
└── *.txt             # Test output files (generated)
```

## Architecture

The shell operates on a command execution loop with the following steps:

1. **Input Reading**: Commands are read from stdin using a buffered reader
2. **Command Parsing**: Input is tokenized with support for quotes, escape sequences, and redirection operators
3. **Redirection Detection**: Parse output redirection operators (`>`) and prepare file handles
4. **Command Resolution**: Commands are checked against built-in commands first, then system PATH
5. **Execution**: Built-in commands execute directly with output redirection support; external commands spawn new processes
6. **Cleanup**: Output files are properly closed after command execution
7. **Loop**: Returns to step 1 to await the next command

## Command Parsing Features

- **Quote Handling**: Support for both single (`'`) and double (`"`) quotes
- **Escape Sequences**: Backslash escaping for special characters
- **Output Redirection**: `>` operator for redirecting stdout to files
- **Argument Tokenization**: Proper splitting of commands and arguments

## License

This project is open source and available under the MIT License.

## Author

Rodrigo - [GitHub](https://github.com/rodrigo134)
