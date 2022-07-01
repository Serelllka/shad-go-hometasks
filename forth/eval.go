package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Evaluator struct {
	stack    []int
	commands map[string]*Command
}

type Command struct {
	executable []func() error
}

func (cmd *Command) Execute() (err error) {
	if cmd == nil {
		return fmt.Errorf("command isn't exists")
	}
	for _, item := range cmd.executable {
		err := item()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Evaluator) Execute(name string) (err error) {
	err = e.commands[name].Execute()
	if err != nil {
		return
	}
	return err
}

func (e *Evaluator) DefineCommand(name string, innerCommands []string) {
	tmp := Command{
		executable: []func() error{},
	}
	for _, item := range innerCommands {
		tmp.executable = append(tmp.executable, e.commands[item].executable...)
	}
	e.commands[name] = &tmp
}

// NewEvaluator creates evaluator.
func NewEvaluator() (e *Evaluator) {
	e = &Evaluator{
		stack:    []int{},
		commands: make(map[string]*Command),
	}
	e.commands["+"] = &Command{
		[]func() error{
			func() error {
				tmp, err := e.Pop(2)
				if err != nil {
					return err
				}
				a, b := tmp[0], tmp[1]
				return e.Push(a + b)
			},
		},
	}
	e.commands["*"] = &Command{
		[]func() error{
			func() error {
				tmp, err := e.Pop(2)
				if err != nil {
					return err
				}
				a, b := tmp[0], tmp[1]
				return e.Push(a * b)
			},
		},
	}
	e.commands["-"] = &Command{
		[]func() error{
			func() error {
				tmp, err := e.Pop(2)
				if err != nil {
					return err
				}
				a, b := tmp[0], tmp[1]
				return e.Push(a - b)
			},
		},
	}
	e.commands["/"] = &Command{
		[]func() error{
			func() error {
				tmp, err := e.Pop(2)
				if err != nil {
					return err
				}
				a, b := tmp[0], tmp[1]
				if b == 0 {
					return fmt.Errorf("division by zero")
				}
				return e.Push(a / b)
			},
		},
	}
	e.commands["dup"] = &Command{
		[]func() error{
			func() error {
				a, err := e.Top()
				if err != nil {
					return err
				}
				err = e.Push(a)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	e.commands["drop"] = &Command{
		[]func() error{
			func() error {
				_, err := e.Pop(1)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	e.commands["over"] = &Command{
		[]func() error{
			func() error {
				a, err := e.Pop(1)
				if err != nil {
					return err
				}
				b, err := e.Top()
				if err != nil {
					return err
				}
				err = e.Push(a[0], b)
				return err
			},
		},
	}
	e.commands["swap"] = &Command{
		[]func() error{
			func() error {
				tmp, err := e.Pop(2)
				if err != nil {
					return err
				}
				a, b := tmp[0], tmp[1]
				return e.Push(b, a)
			},
		},
	}

	return
}

func (e *Evaluator) Push(vars ...int) (err error) {
	if e == nil {
		return fmt.Errorf("this evaluator isn't exists (nil pointer)")
	}
	e.stack = append(e.stack, vars...)
	return nil
}

func (e *Evaluator) Pop(n int) (a []int, err error) {
	if len(e.stack) < n {
		return nil, fmt.Errorf("can't pop %d elems stack is too small", n)
	}
	a = e.stack[len(e.stack)-n:]
	e.stack = e.stack[:len(e.stack)-n]
	return a, nil
}

func (e *Evaluator) Top() (int, error) {
	if e == nil || len(e.stack) == 0 {
		return 0, fmt.Errorf("this evaluation isn't exists (nil pointer)")
	}
	return e.stack[len(e.stack)-1], nil
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(row string) ([]int, error) {
	cmdRegex := regexp.MustCompile("^: [A-Za-z-]+ [A-Za-z ]+ ;")

	for {
		tmp := cmdRegex.FindStringSubmatch(row)
		if len(tmp) == 0 {
			break
		}
		match := tmp[0]
		row = strings.ReplaceAll(row, match, "")
		fmt.Println("current row: ", row)
		match = strings.ReplaceAll(match, ":", "")
		match = strings.ReplaceAll(match, ";", "")
		match = strings.TrimLeft(match, " ")
		match = strings.TrimRight(match, " ")
		expressions := strings.Split(match, " ")
		e.DefineCommand(expressions[0], expressions[1:])
	}

	commands := strings.Split(row, " ")

	for _, item := range commands {
		if intItem, err := strconv.Atoi(item); err == nil {
			err := e.Push(intItem)
			if err != nil {
				return nil, err
			}
		} else if item != "" {
			fmt.Println("command: ", item)
			err := e.Execute(item)
			if err != nil {
				return nil, err
			}
		}
	}

	return e.stack, nil
}
