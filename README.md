# bubbletea-playground

# Bubbletea Playground

Welcome to the Bubbletea Playground! This project is a Go program designed to generate other Go programs based on a clean architecture template. It leverages the Bubbletea library, which employs the ELM architecture to render a CLI user interface.

## Features

- **Repository Naming**: Easily name your new repository.
- **Team Information**: Enter the name of your team.
- **Program Type Selection**: Choose the type of program you want to create:
    - CLI (Cobra)
    - Consumer (NATS)
    - Producer (NATS)
    - Service (Chi)
- **Database Libraries**: Select the database libraries to include:
    - SQLite
    - Oracle
    - MongoDB

## Getting Started

To get started with Bubbletea Playground, follow these steps:

1. **Clone the repository**:
        ```sh
        git clone https://github.com/jonsampson/bubbletea-playground.git
        cd bubbletea-playground
        ```

2. **Install dependencies**:
        ```sh
        go mod tidy
        ```

3. **Run the program**:
        ```sh
        go run cmd/app/main.go
        ```

## Usage

Upon running the program, you will be guided through a series of prompts to configure your new Go project. Follow the on-screen instructions to complete the setup.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Bubbletea](https://github.com/charmbracelet/bubbletea) for the CLI framework.
- [Cobra](https://github.com/spf13/cobra) for the CLI application library.
- [NATS](https://nats.io/) for the messaging system.
- [Chi](https://github.com/go-chi/chi) for the HTTP router.
- [SQLite](https://www.sqlite.org/index.html), [Oracle](https://www.oracle.com/database/), and [MongoDB](https://www.mongodb.com/) for the database libraries.
