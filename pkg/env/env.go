package env

import (
	"log"
	"os"
	"path"
	"strings"

	execute "github.com/alexellis/go-execute/pkg/v1"
)

// GetClientArch returns a pair of arch and os
func GetClientArch() (string, string) {
	task := execute.ExecTask{Command: "uname", Args: []string{"-m"}, StreamStdio: false}
	res, err := task.Execute()
	if err != nil {
		log.Println(err)
	}

	arch := strings.TrimSpace(res.Stdout)

	taskOS := execute.ExecTask{Command: "uname", Args: []string{"-s"}, StreamStdio: false}
	resOS, errOS := taskOS.Execute()
	if errOS != nil {
		log.Println(errOS)
	}

	os := strings.TrimSpace(resOS.Stdout)

	return arch, os
}

func LocalBinary(name, subdir string) string {
	home := os.Getenv("HOME")
	val := path.Join(home, ".bazaar/bin/")
	if len(subdir) > 0 {
		val = path.Join(val, subdir)
	}

	return path.Join(val, name)
}
