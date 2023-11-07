package checkdocker

import "os/exec"

func CheckDocker() bool {
	_, err := exec.Command("docker", "--version").Output()

    if err!= nil {
        return false
    }

    return true
}