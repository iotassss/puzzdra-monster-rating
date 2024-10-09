package vo

type No struct {
	value int
}

func NewNo(value int) (No, error) {
	return No{
		value: value,
	}, nil
}

func (no No) Value() int {
	return no.value
}
