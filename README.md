# Certyclick Verify

Certyclick-verify is a verification tool that allows you to verify the hashed proofs generated from certyclick.com

## Features

- Verify the hash of a file against a previously calculated hash.
- User-friendly graphical user interface (GUI) and command-line interface (CLI) options.
- Cross-platform support for macOS, Windows, and Linux.

## Getting Started

### Prerequisites

- Go programming language (version 1.16 or later) installed on your system.

### Installation

To install the Certyclick Verify application, follow these steps:

1. Clone the project repository:

```
git clone https://github.com/scalingparrots/certyclick_verify.git
```

2. Build the CLI and GUI binaries:

```
cd certyclick_verify/cli
make
```

3. The compiled binaries will be available in the `dist` directory.

### Usage

#### Command-Line Interface (CLI)

To verify the hash of a file against a previously calculated hash:

```
./dist/certyclick_verify_cli <file_path> <hash>
```

#### Graphical User Interface (GUI)

To run the Certyclick Verify GUI:

```
./dist/certyclick_verify_gui
```

1. Click on the "Select File" button to choose a file.
2. Enter the previously calculated hash in the input field.
3. Click on the "Verify Hash" button to compare the calculated hash with the provided hash.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Fyne](https://fyne.io) - GUI toolkit for Go.

## Contact

For any inquiries or support, please contact our team at support@certyclick.com.

