package commands

type Commands struct{}

func New() *Commands {
	return &Commands{}
}

func (cmd *Commands) Execute(str string) error {
	return nil
}
