package main

import (
	"fmt"
	"path/filepath"
	"sync"
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

	wg.Wait()

	PrintSummary(jobs)

}

func PrintSummary(jobs []Job) {
	fmt.Println("Please find the summary below")
	fmt.Print("\n")
	success := 0
	for i := 0; i < len(jobs); i++ {
		var status string
		if jobs[i].Status {
			success++
			status = "Passed 🍕 ✔"
		} else {
			status = "Failed 🗙"
		}
		fmt.Printf("%d) Job %s was %s \n", i+1, jobs[i].Params.Title, status)
	}
	fmt.Print("\n")
	fmt.Printf("%d Success %d total \n", success, len(jobs))
}
