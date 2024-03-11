# Asynchronous Chat Client with Ollama

This project is an asynchronous chat client leveraging the Ollama API for answering user queries. It's built with Python and utilizes `prompt_toolkit` for an enhanced terminal interface, offering a more interactive and engaging user experience.

## Prerequisites

Before installing and running this project, you need to have Python and some package management tools installed on your system. This guide covers the setup for macOS users.

### Installing Python on macOS

1. **Homebrew**: The simplest way to install Python on macOS is via [Homebrew](https://brew.sh/), a popular package manager for macOS. If you haven't installed Homebrew, open Terminal and run:
   ```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```
2. **Install Python**: Once Homebrew is installed, you can install Python by running:
   ```bash
   brew install python
   ```
   This command installs the latest version of Python.

### Using pipx for Global Python Tools

`pipx` is a tool for installing and running Python applications in isolated environments. It's ideal for tools you want to use globally across your system but keep isolated from other Python projects.

1. **Install pipx**:
   ```bash
   brew install pipx
   ```
2. **Initialize pipx**:
   ```bash
   pipx ensurepath
   ```

### Managing Project Dependencies with Poetry

`Poetry` is a tool for dependency management and packaging in Python. It allows you to declare the libraries your project depends on and it will manage (install/update) them for you.

1. **Install Poetry**:
   ```bash
   pipx install poetry
   ```

### Install ollama for macOS

1. **Install ollama**: Download the latest version of ollama from the [official website](https://ollama.com/download) and follow the installation instructions.


## Installation

With Python and the necessary tools installed, follow these steps to get the project up and running.

1. **Clone the repository**:
   ```bash
   git clone https://github.com/rajasoun/ollama-poc.git
   ```
2. **Navigate to the project directory**:
   ```bash
   cd ollama-poc
   poetry shell
   ```
3. **Install dependencies with Poetry**:
   ```bash
   poetry install
   ```

## Usage

To start the chat client, activate the Poetry shell and run the main script:

1. **Activate the Poetry shell**:
   ```bash
   poetry shell
   ```
2. **Run the application**:
   ```bash
   python main.py
   ```

Follow the on-screen instructions. Type your questions and receive responses from the Ollama model. Type `exit` to terminate the session when you're finished.

## Contributing

Contributions to this project are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your features or fixes.
3. Commit your changes and push them to your fork.
4. Submit a pull request to the original repository.

Please adhere to the project's coding standards and guidelines.

## License

This project is open-sourced under the MIT License. See the [LICENSE](LICENSE) file for more details.
