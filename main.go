package main

import (
"fmt"
"os"
"os/exec"
)

func main() {

	php()
	js()

}


func executor(cmd string) string {
	out, err := exec.Command("cmd", "/C", cmd).Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(out)
}

func php() {
	executor("composer cs-check")
	executor("composer test")
}

func js() {
	executor("npm run lint")
	executor("npx stylelint 'src/**/*.scss'")
	executor("npm run test")
}
