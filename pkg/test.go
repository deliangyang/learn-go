package pkg

type T interface {
	Less(than int) bool
}

// Int int
type Int int

func (int Int) Less(than int) bool {
	return false
}
