package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	emitter := NewEmitter()

	var jobs []Job

	rootDir, err := filepath.Abs("$HOME/Documents/api-tools-skeleton")

	if err != nil {
		fmt.Println("Invalid Path")
	}

	phpcsParams := &JobParams{
		Title: "PHPCS",
		Cwd:   rootDir,
		Cmd:   "composer",
		Args:  []string{"run", "cs-check"},
	}

	jobs = append(jobs, *NewJob(phpcsParams, emitter))

	var wg sync.WaitGroup

	for i := 0; i < len(jobs); i++ {
		wg.Add(1)
		go jobs[i].Dispath(i, &wg)
	}

	go PrintSummary(jobs)

	wg.Wait()
	PrintSummary(jobs)
}

func PrintSummary(jobs []Job) {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
	fmt.Println("Please find the summary below")
	fmt.Print("\n")

	success := 0
	completed := true

	for i := 0; i < len(jobs); i++ {
		var status = "Processing "
		if jobs[i].isCompleted {
			if jobs[i].Status {
				success++
				status = "\033[32mPassed ✔\033[0m"
			} else {
				status = "\033[31mFailed 🗙\033[0m"
			}
		} else {
			completed = false
		}

		fmt.Printf("%d) Job %s was %s \n", i+1, jobs[i].Params.Title, status)
	}

	fmt.Print("\n")
	fmt.Printf("%d Success %d total \n", success, len(jobs))

	if !completed {
		time.Sleep(time.Second)
		go PrintSummary(jobs)
	}
}
