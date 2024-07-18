package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
)

var (
	errCd   = errors.New("cd must have 1 argument")
	errPwd  = errors.New("pwd must not have any arguments")
	errEcho = errors.New("echo must have 1+ argument")
	errKill = errors.New("kill must have 1+ argument")
	errExec = errors.New("exec must have 1+ argument")
	errPs   = errors.New("exec must not have any argument")
)

func main() {
	sh := NewShell(os.Stdout, os.Stdin)
	err := sh.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

type Shell struct {
	Out      io.Writer
	In       io.Reader
	Pipe     bool
	PipeBuff *bytes.Buffer
}

func NewShell(w io.Writer, r io.Reader) *Shell {
	return &Shell{Out: w, In: r}
}

func (s *Shell) Run() error {
	if err := s.GetLines(); err != nil {
		fmt.Fprintln(s.Out, err)
	}
	return nil
}

func (s *Shell) cd(arg string) error {
	return os.Chdir(arg)
}

func (s *Shell) pwd() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	out := s.Out
	if s.Pipe {
		out = s.PipeBuff
	}

	_, err = fmt.Fprintln(out, path)
	return err
}

func (s *Shell) echo(args []string, fullLine string) error {
	printer := s.Out
	if s.Pipe {
		printer = s.PipeBuff
	}

	line := strings.Join(args, " ")
	line = strings.Trim(line, "\"")
	_, err := fmt.Fprintln(printer, line)
	return err
}

func (s *Shell) kill(pid []string) []error {
	var errs []error
	for _, value := range pid {
		cmd := exec.Command("taskkill", "/F", "/PID", value)
		if err := cmd.Run(); err != nil {
			errs = append(errs, fmt.Errorf("failed to kill process id %s: %v", value, err))
		}
	}
	return errs
}

func (s *Shell) ps() error {
	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	out := s.Out
	if s.Pipe {
		out = s.PipeBuff
	}

	for _, process := range processList {
		_, err = fmt.Fprintf(out, "%v\t%v\t%v\n", process.Pid(), process.PPid(), process.Executable())
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Shell) GetLines() error {
	scanner := bufio.NewScanner(s.In)
	fmt.Fprint(s.Out, "$ ")
	for scanner.Scan() {
		line := scanner.Text()
		if line == `\quit` {
			break
		}
		if err := s.Fork(line); err != nil {
			return err
		}
		fmt.Fprint(s.Out, "$ ")
	}
	return scanner.Err()
}

func (s *Shell) Exec(line []string) error {
	cmd := exec.Command(line[0], line[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if s.Pipe {
		cmd.Stdout = s.PipeBuff
	}
	return cmd.Run()
}

func (s *Shell) CaseShell(line string) error {
	commandAndArgs := strings.Fields(line)
	if len(commandAndArgs) == 0 {
		return nil
	}

	switch commandAndArgs[0] {
	case "cd":
		if len(commandAndArgs) == 2 {
			return s.cd(commandAndArgs[1])
		} else {
			return errCd
		}
	case "pwd":
		if len(commandAndArgs) == 1 {
			return s.pwd()
		} else {
			return errPwd
		}
	case "echo":
		if len(commandAndArgs) > 1 {
			return s.echo(commandAndArgs[1:], strings.Join(commandAndArgs[1:], " "))
		} else {
			return errEcho
		}
	case "kill":
		if len(commandAndArgs) > 1 {
			errs := s.kill(commandAndArgs[1:])
			for _, err := range errs {
				fmt.Fprintln(s.Out, err)
			}
		} else {
			return errKill
		}
	case "ps":
		if len(commandAndArgs) == 1 {
			return s.ps()
		} else {
			return errPs
		}
	case "exec":
		if len(commandAndArgs) > 1 {
			return s.Exec(commandAndArgs[1:])
		} else {
			return errExec
		}
	default:
		fmt.Fprintf(s.Out, "unknown command '%v'\n", commandAndArgs[0])
	}
	return nil
}

func (s *Shell) Fork(str string) error {
	str = strings.TrimRight(str, " ")
	if strings.HasSuffix(str, "&") {
		str = strings.TrimSuffix(str, "&")
		cmd := exec.Command("cmd", "/C", str)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		if err := cmd.Start(); err != nil {
			return err
		}
		go func() {
			cmd.Wait()
		}()
	} else {
		return s.CheckPipes(str)
	}
	return nil
}

func (s *Shell) CheckPipes(line string) error {
	strCmd := strings.Split(line, "|")
	if len(strCmd) > 1 {
		s.Pipe = true
		s.PipeBuff = new(bytes.Buffer)
		for _, value := range strCmd {
			s.PipeBuff.Reset()
			if err := s.CaseShell(value); err != nil {
				fmt.Fprintln(s.Out, err)
			}
		}
		s.Pipe = false
	} else {
		if err := s.CaseShell(line); err != nil {
			return err
		}
	}
	return nil
}
