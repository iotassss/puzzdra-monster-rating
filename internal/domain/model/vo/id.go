package vo

const (
	TemporaryID = 0
)

type ID struct {
	value int
}

func NewID(value int) (ID, error) {
	return ID{
		value: value,
	}, nil
}

func NewTemporaryID() ID {
	return ID{
		value: TemporaryID,
	}
}

func (id ID) Value() int {
	return id.value
}
