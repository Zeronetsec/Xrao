package console

type Command interface {
    Execute(args []string)
}