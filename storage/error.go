package storage

type Error struct {
	Message string
	Status  int
}

func (e *Error) Error() string {
	return e.Message
}
