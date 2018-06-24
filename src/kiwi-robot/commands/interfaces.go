package commands

// ICommand interface allow to implement command tah will be invoked during
// execution. You must call Execute() method to execute command
type ICommand interface {
	// Execute command
	Execute() error
}

// IPlacer interface implements ICommand interface for special PLACE command, because
// it need also to invoke with parameters: x, y and direction. Also don't forget
// to call Execute() method
type IPlacer interface {
	SetX(x uint8)
	SetY(y uint8)
	SetDir(dir string)
	Execute() error
	Clone() IPlacer
}