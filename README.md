# go-log-analyser

## Description

go-log-analyser is a simple Linux log analyzer written in Go. This project is primarily for learning Go and exploring log analysis techniques. It analyzes SSH failed password attempts from a log file and provides a summary of the top IPs and usernames involved.

## Features

-   Parses a log file line by line.
-   Identifies lines indicating failed SSH password attempts.
-   Extracts usernames and IP addresses from those lines using regular expressions.
-   Counts the number of failed attempts per IP address and username.
-   Prints a summary of the top IPs and usernames with the number of failed attempts.
-   Accepts a command-line flag to specify the log file.

## Getting Started

### Prerequisites

-   Go installed on your system.

### Installation

1.  Clone the repository:

    ```bash
    git clone <repository-url>
    ```

2.  Navigate to the project directory:

    ```bash
    cd go-log-analyser
    ```

### Usage

1.  Build the application:

    ```bash
    go build main.go
    ```

2.  Run the application:

    ```bash
    ./main -file <path/to/log-file>
    ```

    For example, to analyze the `example.auth.log` file:

    ```bash
    ./main -file /var/log/auth.log
    ```

    If no file is specified, the program will default to reading `auth.log`.

### Configuration

-   The log file path can be specified using the `-file` command-line flag.
-   If no flag is provided, the program defaults to `auth.log`.

## Contributing

Feel free to contribute to this project by submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

-   This project was created as a learning exercise.