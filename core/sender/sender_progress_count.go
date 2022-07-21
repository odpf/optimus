package sender

type ProgressCount interface {
	Add(count int) error
	Inc() error
}

func ProgressAdd(p ProgressCount, count int) {
	if p == nil {
		return
	}
	p.Add(count)
}

func ProgressInc(p ProgressCount) {
	ProgressAdd(p, 1)
}
