#!/usr/bin/env bash

# Color code declarations for easy understanding and modification
COLOR_GREEN="70"
COLOR_YELLOW="178"
COLOR_RED="124"
COLOR_BLUE="35"

# Enhanced pretty_print function with color parameter
function pretty_print() {
    local message=$1
    local color=$2 # Use color variable here
    gum style --border normal --margin "1" --padding "1" --border-foreground $color "$message"
}

# Check if a command exists
function command_exists() {
    command -v "$1" &> /dev/null
}
# Check if an item is installed
function is_installed() {
    local item=$1
    local type=$2 # This should be either "--cask" for casks or left empty for formulae

    # Check if the type is specifically for casks or formulae
    if [[ -z "$type" ]]; then
        # When type is empty, explicitly list formulae to ignore casks
        brew list --formula | grep -q "^$item$"
    else
        # When type is "--cask", list casks
        brew list $type | grep -q "^$item$"
    fi
}

# Function to install Homebrew
function install_homebrew() {
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
}

# Function to check and install Homebrew if not installed
function check_install_homebrew() {
    if ! command_exists brew; then
        pretty_print "Homebrew is not installed. Installing..." red
        install_homebrew
        pretty_print "Homebrew installation completed." green
    fi
}

# Function to check and install GUM if not installed
function check_install_gum() {
    if ! command_exists gum; then
        pretty_print "GUM is not installed. Installing..." red
        brew install gum
        pretty_print "GUM installation completed." green
    fi
}

# Install a single item (package or cask), ignoring if already installed
function install_item() {
    local item=$1
    local type=$2 # Empty for packages, --cask for casks
    if command_exists "$item" || is_installed "$item" "$type"; then
        pretty_print "$item is already installed." $COLOR_YELLOW
        return 0
    fi

    if brew install $type "$item"; then
        pretty_print "Installed $item successfully." $COLOR_GREEN
    else
        pretty_print "Failed to install $item." $COLOR_RED
        return 1
    fi
}

# Install multiple items (packages or casks)
function install_items() {
    local items=("${@:2}") # Skip the first argument
    local type=$1 # Empty for packages, --cask for casks
    local item_type="package(s)"
    if [[ "$type" == "--cask" ]]; then
        item_type="cask(s)"
    fi

    if [ ${#items[@]} -ne 0 ]; then
        pretty_print "Installing $item_type..." $COLOR_BLUE
        for item in "${items[@]}"; do
            install_item "$item" "$type"
        done
    else
        pretty_print "No $item_type to install." $COLOR_YELLOW
    fi
}

# Main function to install apps
function install_apps() {
    install_items "" "${PACKAGES[@]}"
    install_items "--cask" "${CASKS[@]}"
}

# Uninstall a single item
function uninstall_item() {
    local item=$1
    local type=$2 # --formula for packages, --cask for casks
    if is_installed "$item" "$type"; then
        brew uninstall $type "$item"
        pretty_print "Uninstalled $item." $COLOR_GREEN
    else
        pretty_print "$item is not installed." $COLOR_YELLOW
    fi
}

# Uninstall multiple items
function uninstall_items() {
    local items=("${@:2}") # Skip the first argument
    local type=$1 # --formula for packages, --cask for casks
    local item_type="package(s)"
    if [[ "$type" == "--cask" ]]; then
        item_type="cask(s)"
    fi

    if [ ${#items[@]} -ne 0 ]; then
        pretty_print "Checking and uninstalling $item_type..." $COLOR_BLUE
        for item in "${items[@]}"; do
            uninstall_item "$item" "$type"
        done
    else
        pretty_print "No $item_type to uninstall." $COLOR_YELLOW
    fi
}

# Main function to uninstall apps
function uninstall_apps() {
    uninstall_items "--formula" "${PACKAGES[@]}"
    uninstall_items "--cask" "${CASKS[@]}"
}

# Execute the main function
function bootstrap_main() {
    echo "BASEDIR: $BASEDIR"
    # Get the app directory name
    APPDIR=$(basename "$BASEDIR")
    echo "$APPDIR Env Setup/Teardown"

    # Read packages and casks from files
    PACKAGES=($(cat "$BASEDIR/scripts/packages/brew.txt"))
    CASKS=($(cat "$BASEDIR/scripts/packages/casks.txt"))

    check_install_homebrew # Check and install Homebrew
    check_install_gum # Check and install GUM

    # Main menu using gum
    ACTION=$(gum choose "Setup" "Teardown")
    echo "Selected action: $ACTION"
    case "$ACTION" in
        Setup)
            install_apps
            ;;
        Teardown)
            uninstall_apps
            ;;
        *)
            pretty_print "Invalid option, exiting..." $COLOR_RED
            exit 1
            ;;
    esac
}
