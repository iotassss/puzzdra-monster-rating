package vo

const (
	TemporaryID = 0
)

type ID struct {
	value int
}

func NewID(value int) (ID, error) {
	if value < 0 {
		return ID{}, ErrInvalidID{message: "ID must be greater than or equal to 0"}
	}
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

type ErrInvalidID struct {
	message string
}

func (e ErrInvalidID) Error() string {
	return e.message
}
