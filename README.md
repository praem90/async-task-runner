# Async Task Runner
Executes some commands asynchronously. Written in Go. It executes the jobs in parellel and provides summary. It is helpful to run set of commands something like git pre-commit hook 

## TODO
 - [ ] Read jobs from a config file
 - [ ] Ability to specify config file location via cmd line args
 - [ ] Use io buffers to pipe Cmd's `stdout` and `stderr`
