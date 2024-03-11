# Automated Dev Tools Bootstrap

This repository contains a collection of scripts and configuration files to automate the setup and teardown of development tools and applications on macOS.

---

## Setup and Teardown with Bootstrap Script

This guide explains how to set up or tear down your environment using our bootstrap script hosted on GitHub. The script automates the installation or uninstallation of applications and packages defined in `packages/brew.txt` and `packages/casks.txt` within the repository.

### Prerequisites

- Ensure you have `brew` installed on your machine. Visit [Homebrew's website](https://brew.sh/) for installation instructions.
- `gum` is required for the interactive menus. Install it via Homebrew with `brew install gum` if you haven't already.
- An active internet connection.

### Running the Script

You can directly source and execute the script from your terminal with the following command:

```bash
source <(curl -s https://raw.githubusercontent.com/rajasoun/dev-tools-bootstrap/main/assist.sh)
```

After executing the command above, follow the interactive prompts to either set up or tear down your environment.

### What the Script Does

- **Setup**: Installs the applications and packages listed in `packages/brew.txt` and `packages/casks.txt`.
- **Teardown**: Uninstalls the applications and packages listed in the same files.

### Notes

- Running the script requires executing commands from an external source. Always review scripts before executing them for security purposes.
- The script uses `gum` for improved user interaction. Ensure it is installed and up to date.
- The installation and uninstallation processes are dependent on the content of `packages/brew.txt` and `packages/casks.txt`. Modify these files according to your needs before running the script.

### Troubleshooting

If you encounter any issues:
- Check your internet connection to ensure the script can be fetched.
- Ensure `brew` and `gum` are correctly installed and accessible in your terminal.
- Verify the script's URL is correct and accessible from your browser.

---
