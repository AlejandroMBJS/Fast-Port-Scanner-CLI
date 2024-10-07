
# Fast Port Scanner CLI

## Overview

The **Fast Port Scanner CLI** is a lightweight command-line application designed to efficiently scan for open ports on a specified IP address. Built with performance in mind, this tool provides a straightforward and effective way to identify open ports and the associated services, all while utilizing minimal system resources.

## Features

- **Concurrent Scanning**: Leverages multiple goroutines to perform scans quickly, significantly reducing the time required to identify open ports.
- **Service Detection**: Automatically detects and displays common services associated with open ports, such as HTTP, FTP, and SSH.
- **User-Friendly Interface**: Clear command-line interface that provides intuitive input options and output formatting.
- **Progress Feedback**: Displays a progress bar and loading animation during the scanning process to keep the user informed.
- **Configurable Timeout**: Allows users to set a custom timeout for each connection attempt, making it adaptable to different network conditions.

## Requirements

- **Go**: Version 1.16 or higher is required to build and run the application.
- **Operating System**: Compatible with Windows, macOS, and Linux.

## Installation

Follow these steps to install and run the Fast Port Scanner CLI:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/fast-port-scanner.git
   cd fast-port-scanner
   ```

2. **Build the application**:
   ```bash
   go build -o port-scanner main.go
   ```

3. **Run the application**:
   ```bash
   ./port-scanner -ip <IP address> -timeout <timeout in ms>
   ```

## Usage

### Basic Command

```bash
./port-scanner -ip 127.0.0.1 -timeout 1000
```

### Options

- `-ip <IP address>`: Specify the target IP address to scan. This option is required.
- `-timeout <milliseconds>`: Set the timeout duration for each connection attempt. Default value is 500 ms.
- `-start <port>`: Specify the starting port for the scan. Default value is 1.
- `-end <port>`: Specify the ending port for the scan. Default value is 65535.

### Example

To scan the local loopback address with a timeout of 1000 milliseconds, use:

```bash
sudo ./port-scanner -ip 127.0.0.1 -timeout 1000
```

## Contribution

Contributions are welcome! If you encounter any errors or have suggestions for improvements, please feel free to fork the repository, make your changes, and submit a pull request. I would be glad to hear any feedback or corrections you might have.
Let me know if you need any further adjustments!
