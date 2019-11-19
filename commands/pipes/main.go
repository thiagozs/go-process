package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

// execute a stack of command
func execute(output_buffer *bytes.Buffer, stack ...*exec.Cmd) (err error) {
	var error_buffer bytes.Buffer
	pipe_stack := make([]*io.PipeWriter, len(stack)-1)
	i := 0
	for ; i < len(stack)-1; i++ {
		stdin_pipe, stdout_pipe := io.Pipe()
		stack[i].Stdout = stdout_pipe
		stack[i].Stderr = &error_buffer
		stack[i+1].Stdin = stdin_pipe
		pipe_stack[i] = stdout_pipe
	}
	stack[i].Stdout = output_buffer
	stack[i].Stderr = &error_buffer

	if err := call(stack, pipe_stack); err != nil {
		return err
	}

	return nil
}

// call to execute a pipe command
func call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = call(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}

func main() {
	var b bytes.Buffer

	// commands to execute
	cmds := []*exec.Cmd{
		exec.Command("cat", "/home/thiagozs/Downloads/oui.csv"),
		exec.Command("wc", "-l"),
		//exec.Command("sort", "-M"),
	}

	// execute pipe commands on shell
	if err := execute(&b, cmds...); err != nil {
		log.Printf("Erro on execute: %s", err)
		return
	}

	// show all returned pipes stdout to buffer
	_, _ = io.Copy(os.Stdout, &b)
}
