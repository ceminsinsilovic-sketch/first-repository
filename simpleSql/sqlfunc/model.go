package sqlfunc

import "time"

type Book struct {
	ID       int
	Name     string
	Author   string
	Review   string
	Year     int
	IsRead   bool
	IsThere  time.Time
	FullRead *time.Time
}
