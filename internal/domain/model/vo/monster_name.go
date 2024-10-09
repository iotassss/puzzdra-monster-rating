package vo

const (
	maxMonsterNameLength = 256
	minMonsterNameLength = 1
)

type ErrMonsterNameValidation struct {
	Message string
}

func (e *ErrMonsterNameValidation) Error() string {
	return e.Message
}

type MonsterName struct {
	value string
}

func NewMonsterName(value string) (MonsterName, error) {
	if value == "" {
		return MonsterName{}, &ErrGame8MonsterPointValidation{Message: "monster name is empty"}
	}
	if len(value) < minMonsterNameLength {
		return MonsterName{}, &ErrGame8MonsterPointValidation{Message: "monster name is too short"}
	}
	if len(value) > maxMonsterNameLength {
		return MonsterName{}, &ErrGame8MonsterPointValidation{Message: "monster name is too long"}
	}

	return MonsterName{value: value}, nil
}

func (n MonsterName) Value() string {
	return n.value
}

func (n MonsterName) EqualsString(other string) bool {
	return n.value == other
}
