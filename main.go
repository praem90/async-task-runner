package main

import (
"fmt"
"os"
"os/exec"
)

func main() {
	fmt.Println("Executing pre-commit hook")

	phpChan := make(chan int)
	jsChan := make(chan int)


	go php(phpChan)
	go js(jsChan)

	php := <-phpChan
	fmt.Println("Status of php", php)

	js := <-jsChan
	fmt.Println("Status of JS", js)
}


func executor(cmd string) int {
	fmt.Println("Executing", cmd)
	_, err := exec.Command("cmd", "/C", cmd).Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return 1
	}

	return 0
}

func php(c chan int) int {
	fmt.Println("Executing php")
	executor("composer cs-check")
	executor("composer test")


	c <- 1
	
	return 1
}

func js(c chan int) {
	fmt.Println("Executing js")
	executor("npm run lint")
	executor("npx stylelint 'src/**/*.scss'")
	executor("npm run test")
	
	c <- 1
}
