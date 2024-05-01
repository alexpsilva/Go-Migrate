package cli

type Command func(parameters []string)

type CLI struct {
	commands map[string]Command
}

func New() *CLI {
	return &CLI{
		commands: make(map[string]Command),
	}
}

func (cli *CLI) AddCommand(name string, command Command) {
	cli.commands[name] = command
}

func (cli *CLI) RunCommand(name string, parameters []string) {
	cli.commands[name](parameters)
}
