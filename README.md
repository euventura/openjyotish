# OpenJyotish

OpenJyotish is a free and open-source Vedic Astrology (Jyotish) software.

It is designed to be a comprehensive tool for astrologers, students, and enthusiasts of Vedic Astrology.

## Features

*   **Astrological Calculations:** Accurate calculations for planetary positions, divisional charts, dashas, and more.
*   **API Server:** Expose Jyotish calculation capabilities through a web API.
*   **Cross-Platform Desktop Application:** Run OpenJyotish as a standalone desktop application on Windows, macOS, and Linux.

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

*   Go programming language (check `go.mod` for version)

### Build

1.  Clone the repository.
2.  Navigate to the project directory.
3.  Build the application:
    ```sh
    go build -o openjyotish ./main.go
    ```

## Usage

The application can be run in two different modes.

### API Mode

To start the server, run the following command:

```sh
./openjyotish --mode api
```
*(Note: This is a placeholder for the actual command-line flag if implemented).*

### Desktop Mode

To launch the graphical user interface, run:

```sh
./openjyotish --mode desktop
```
*(Note: This is a placeholder for the actual command-line flag if implemented).*

## License

Distributed under the MIT License. See `LICENSE` for more information.
