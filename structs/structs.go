package structs

type Task struct {
	ID          int
	Title       string
	Description string
	Deadline    string
}

type TempTask struct {
	Title       string
	Description string
	Deadline    string
}
