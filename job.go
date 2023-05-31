package main

import (
	"os/exec"
	"sync"
	"time"
)

type JobParams struct {
	Title string
	Cwd   string
	Cmd   string
	Args  []string
}

type Job struct {
	Params      *JobParams
	Status      bool
	Output      string
	isCompleted bool
	ExitCode    int
	emitter     *Emitter
}

func NewJob(params *JobParams, emitter *Emitter) *Job {
	job := &Job{}

	job.Params = params
	job.Status = false
	job.Output = ""
	job.ExitCode = 0
	job.emitter = emitter
	job.isCompleted = false

	return job
}

func (job *Job) Dispath(i int, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)

	// TODO: Execute the cmd and emit an event
	cmd := exec.Command(job.Params.Cmd, job.Params.Args...)
	cmd.Dir = job.Params.Cwd

	out, err := cmd.Output()

	job.ExitCode = cmd.ProcessState.ExitCode()

	if err != nil {
		job.Status = false
		job.Output = err.Error()

		job.emitter.Emit("failed", job, i)
	} else {

		job.Output = string(out)
		job.Status = true

		job.emitter.Emit("success", job, i)
	}

	job.isCompleted = true

	job.emitter.Emit("complete", job, i)

	wg.Done()
}
