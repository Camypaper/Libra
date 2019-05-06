package libra

/*
Task !
*/
type Task interface {
	Runnable
	Name() string
}
