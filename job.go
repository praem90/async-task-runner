package main

type JobParams struct {
	Title string
	Type  string
	Cwd   string
	Cmd   string
}

type Job struct {
	Params   JobParams
	Status   bool
	Output   string
	ExitCode int
}

func (job *Job) New(params JobParams) Job {
	job.Params = params
	job.Status = false
	job.Output = ""
	job.ExitCode = 0

	return *job
}

func (job *Job) Dispath() {
	// TODO: Execute the cmd and emit an event
}
