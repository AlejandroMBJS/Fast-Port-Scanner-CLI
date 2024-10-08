# GoScan

## Overview

The **GoScan** is a lightweight command-line application designed to efficiently scan for open ports on a specified IP address. Built with performance in mind, this tool provides a straightforward and effective way to identify open ports and the associated services, all while utilizing minimal system resources.

## Features

- **Concurrent Scanning**: Leverages multiple goroutines to perform scans quickly, significantly reducing the time required to identify open ports.
- **Service Detection**: Automatically detects and displays common services associated with open ports, such as HTTP, FTP, and SSH.
- **User-Friendly Interface**: Clear command-line interface that provides intuitive input options and output formatting.
- **Progress Feedback**: Displays a progress bar and loading animation during the scanning process to keep the user informed.
- **Configurable Timeout**: Allows users to set a custom timeout for each connection attempt, making it adaptable to different network conditions.
- **Export Results**: Option to export scan results, including a suggested `nmap` command for further exploration of the found open ports.

## Requirements

- **Go**: Version 1.16 or higher is required to build and run the application.
- **Operating System**: Compatible with Windows, macOS, and Linux.

## Installation

Follow these steps to install and run the GoScan:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/goscan.git
   cd goscan
   ```

2. **Build the application**:
   ```bash
   go build -o goscan main.go
   ```

3. **Run the application**:
   ```bash
   ./goscan -ip <IP address> -timeout <timeout in ms optional> -o <output filename>
   ```

## Usage

### Basic Command

```bash
./goscan -ip 127.0.0.1
```

### Options

- `-ip <IP address>`: Specify the target IP address to scan. This option is required.
- `-timeout <milliseconds>`: Set the timeout duration for each connection attempt. Default value is 500 ms.
- `-o <filename>`: Specify an output filename to export results, including a suggested `nmap` command.


### Example

To scan the local loopback address with a timeout of 1000 milliseconds and export the results, use:

```bash
sudo ./goscan -ip 127.0.0.1 -timeout 1000 -o results.txt
```

## Contribution

Contributions are welcome! If you encounter any errors or have suggestions for improvements, please feel free to fork the repository, make your changes, and submit a pull request. I would be glad to hear any feedback or corrections you might have.

Let me know if you need any further adjustments!
