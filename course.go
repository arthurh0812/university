package university

type Course struct {
	minStudents int32
	maxStudents int32
}

func (c *Course) IsCancelled() bool {
	return false
}