package showimages

import (
	"fmt"
    "os/exec"
    "log"
	"strings"
	"regexp"
)

type DockerImage struct {
	Repository string 
	Tag        string
	ImageID    string
	Created    string
	Size       string
}

func  GetAll() ([]DockerImage) {

	dockerImages := make([]DockerImage, 0)
	out, err := exec.Command("docker", "images").Output()

	if err!= nil {
        log.Fatal(err)
    }

	splitedOutput := strings.Split(string(out),"\n")

	for i, v := range splitedOutput {
		if i == 0 {
			fmt.Println(v)
            continue
        }
		r, _ := regexp.Compile(`\s{2,}`)
		v := r.ReplaceAllString(v, "  ")

        splitedV := strings.Split(v, "  ")
		if len(splitedV) < 3 {
            continue
        }

        image := DockerImage{
            Repository: splitedV[0],
            Tag:        splitedV[1],
            ImageID:    splitedV[2],
            Created:    splitedV[3],
            Size:       splitedV[4],
        }
		dockerImages =  append(dockerImages, image)
	}
	return dockerImages
}
