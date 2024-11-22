![image](https://github.com/user-attachments/assets/3d5dabab-1f48-4954-a60d-889475f686a1)


# Govibes 🎹🔊
An unnecessary rewrite of [mechvibes.com](https://mechvibes.com) (mechanical keyboard sound simulator) disguised as a CLI tool.

## Prerequisites

### System Requirements
- Linux Operating System
- Go (version 1.16 or higher)
- SoX (Sound eXchange) - Audio playing utility

## Installation

### 1. Install Go
Ensure Go is installed on your system:
```bash
# Check Go installation
go version
```
If not installed, download from [official Go website](https://golang.org/dl/)

### 2. Install SoX
SoX is required for audio playback:
```bash
# On Ubuntu/Debian
sudo apt-get update
sudo apt-get install sox

# On Fedora
sudo dnf install sox

# On Arch Linux
sudo pacman -S sox
```

### 3. Clone the Repository
```bash
git clone https://github.com/manish-mehra/govibes.git
cd govibes
```

### 4. Install Dependencies
```bash
go mod tidy
```

### 5. Build the Project
```bash
go build
```

### 6. Run Govibes
```bash
# Run in interactive mode
./govibes
```

## Optional: Create an Alias
For quick access, you can add an alias to your `~/.bashrc`:
```bash
alias govibes="cd ~/path/to/govibes && ./govibes"
```

## Usage

### Commands
- `govibes sounds`: List available keyboard sound profiles
- `govibes <profile>`: Play a specific sound profile
- `govibes default`: Play the last used sound profile
- `govibes`: Enter interactive mode

## Keyboard Input
Govibes uses Linux-specific methods to listen to keyboard inputs:
- Reads from `/proc/bus/input/devices`
- Accesses keyboard event files in `/dev/input/`

## Limitations
- **Platform**: Linux only
- Requires access to keyboard event files (might need sudo/permissions)

## Troubleshooting
- Ensure you have read permissions for `/dev/input/eventX`
- Check SoX is correctly installed
- Verify Go environment is set up

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## Acknowledgments
Inspired by [mechvibes.com](https://mechvibes.com)
