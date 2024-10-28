# Docker Installation Script

This script installs Docker on either Debian-based or Red Hat-based Linux distributions. It checks for root privileges, identifies the Linux distribution, installs Docker using the appropriate package manager, and verifies the installation by running the `hello-world` Docker container.

## Features

- Verifies if the user has root privileges.
- Detects the Linux distribution to determine the appropriate package manager (`apt` for Debian-based and `yum` for Red Hat-based).
- Installs Docker by executing a series of commands specific to each distribution.
- Verifies Docker installation by running a test container.

## Requirements

- **Root Access:** This script must be run as `root` to ensure proper installation of Docker and related dependencies.
- **Internet Connection:** An active internet connection is required to download and install Docker packages.

## Supported Distributions

- **Debian-based:** Ubuntu, Debian, Mint, Pop!_OS, elementaryOS, Kali, Parrot.
- **Red Hat-based:** Fedora, CentOS, Red Hat Enterprise Linux.

## Usage

1. Clone this repository or copy the script file.
2. Make sure you have Go installed on your machine and run the script as `root`:
    ```bash
    sudo go run script.go
    ```

## Explanation of Code

1. **Root Check:** The script checks if the user is running it as `root`. If not, it asks to switch to `root`.
2. **Distribution Detection:** It identifies the Linux distribution using the `lsb_release` command.
3. **Docker Installation:** Based on the detected distribution, it installs Docker with either `apt` or `yum` and starts the Docker service.
4. **Verification:** After installation, it runs the `hello-world` Docker container to confirm successful setup.

## Error Handling

If any command in the installation sequence fails, the script will stop and print an error message.

## Output

- **Successful Installation:** Outputs "Installation is successful and Docker is running fine!!".
- **Failed Installation:** Prints "Installation has failed" if there is an issue during setup.

---