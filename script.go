package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	exit := 0

	output, err := exec.Command("whoami").Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	user := strings.TrimSpace(string(output))

	if user == "root" {
		fmt.Println("You are root")
	} else {
		fmt.Println("Please switch to root to run this program")
		exit++
	}

	if exit == 0 {
		output, err = exec.Command("sh", "-c", "lsb_release -a | grep 'Distributor ID:'").Output()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		distro := string(output)
		distroID := ""

		for _, line := range strings.Split(distro, "\n") {
			if strings.HasPrefix(line, "Distributor ID:") {
				distroID = strings.TrimSpace(strings.Split(line, ":")[1])
				break
			}
		}

		var pacman string

		if distroID == "Ubuntu" || distroID == "Debian" || distroID == "Mint" || distroID == "Pop!_OS" || distroID == "elementaryOS" || distroID == "Kali" || distroID == "Parrot" {
			fmt.Println("You are using a Debian-based distro")
			pacman = "apt"
		} else if distroID == "Fedora" || distroID == "CentOS" || distroID == "Red Hat Enterprise Linux" {
			fmt.Println("You are using a Red Hat-based distro")
			pacman = "yum"
		} else {
			fmt.Println("Unsupported distribution")
			exit++
		}

		if pacman == "yum" {
			commands := []string{
				"yum update -y",
				"yum remove -y docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-selinux docker-engine-selinux docker-engine",
				"yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin",
				"systemctl start docker",
			}

			for _, cmd := range commands {
				fmt.Println("Executing:", cmd)
				_, err := exec.Command("sh", "-c", cmd).CombinedOutput()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					exit++
					break
				}
			}
		}

		if pacman == "apt" {
			commands := []string{
				"apt-get update",
				"apt-get install -y ca-certificates curl",
				"install -m 0755 -d /etc/apt/keyrings",
				"curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc",
				"chmod a+r /etc/apt/keyrings/docker.asc",
				`echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null`,
				"apt-get update",
				"apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin",
				"service docker start",
			}

			for _, cmd := range commands {
				fmt.Println("Executing:", cmd)
				_, err := exec.Command("sh", "-c", cmd).CombinedOutput()
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					exit++
					break
				}
			}
		}
	}

	if exit == 0 {
		cmd := exec.Command("docker", "run", "hello-world")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Println("Installation has failed")
			return
		}

		if strings.Contains(string(output), "Hello from Docker!") {
			fmt.Println("Installation is successful and Docker is running fine!!")
		} else {
			fmt.Println("Installation has failed")
		}
	}

	if exit != 0 {
		fmt.Println("Exiting with errors.")
	}
}
