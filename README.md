# Shell-Go

A lightweight command-line shell implementation written in Go. This project demonstrates fundamental shell architecture including command parsing, built-in command execution, and external command invocation.

## Features

- **Interactive Shell Prompt**: Real-time command line interface with input handling
- **Built-in Commands**: Core shell commands implemented natively
  - `echo`: Output text arguments
  - `pwd`: Print working directory
  - `exit`: Terminate the shell
  - `type`: Display command type information
- **External Command Execution**: Execute system binaries with full argument support
- **Standard I/O Handling**: Complete stdin/stdout/stderr integration
- **Command Parsing**: Automatic tokenization and argument processing

## Requirements

- Go 1.26.2 or later

## Installation

Clone the repository and build the project:

```bash
git clone https://github.com/rodrigo134/shell-go.git
cd shell-go
go build -o shell-go main.go
```

## Usage

Run the shell executable:

```bash
./shell-go
```

The shell will display a prompt (`$ `) and await user input. Enter commands as you would in any standard shell:

```bash
$ echo Hello, World!
Hello, World!

$ pwd
/home/user/projects

$ ls -la
total 24
drwxr-xr-x  3 user user 4096 May 11 10:00 .
drwxr-xr-x 10 user user 4096 May 10 15:30 ..
-rw-r--r--  1 user user  123 May 11 10:00 main.go

$ type echo
echo is a shell builtin

$ exit
```

## Built-in Commands

| Command | Description | Example |
|---------|-------------|---------|
| `echo` | Print arguments to standard output | `echo Hello World` |
| `pwd` | Print the current working directory | `pwd` |
| `type` | Display whether a command is a builtin or not found | `type ls` |
| `exit` | Terminate the shell | `exit` |

## Project Structure

```
.
├── main.go      # Shell implementation
├── go.mod       # Go module definition
└── README.md    # Project documentation
```

## Architecture

The shell operates on a simple command execution loop:

1. **Input Reading**: Commands are read from stdin using a buffered reader
2. **Command Parsing**: Input is tokenized into command name and arguments
3. **Command Resolution**: Commands are checked against built-in commands first, then system PATH
4. **Execution**: Built-in commands execute directly; external commands spawn new processes
5. **Loop**: Returns to step 1 to await the next command

## License

This project is open source and available under the MIT License.

## Author

Rodrigo - [GitHub](https://github.com/rodrigo134)
