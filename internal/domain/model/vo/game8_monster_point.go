package vo

import (
	"fmt"
	"regexp"
	"strconv"
)

type Game8MonsterPoint struct {
	value string
}

func NewGame8MonsterPoint(valueStr string) (Game8MonsterPoint, error) {
	if valueStr == "-" {
		return Game8MonsterPoint{
			value: valueStr,
		}, nil
	}

	re := regexp.MustCompile(`^([0-9](\.[0-9])?)$`)
	if !re.MatchString(valueStr) {
		return Game8MonsterPoint{}, &ErrGame8MonsterPointValidation{Message: "score must be a number between 0.0 and 10.0"}
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return Game8MonsterPoint{}, &ErrGame8MonsterPointValidation{Message: fmt.Sprintf("failed to parse score: %s", err.Error())}
	}

	if value < 0.0 || value > 10.0 {
		return Game8MonsterPoint{}, &ErrGame8MonsterPointValidation{Message: "score must be a number between 0.0 and 10.0"}
	}

	return Game8MonsterPoint{
		value: valueStr,
	}, nil
}

func (s Game8MonsterPoint) Value() string {
	return s.value
}

func (s Game8MonsterPoint) Equals(other Game8MonsterPoint) bool {
	return s.value == other.value
}

func (s Game8MonsterPoint) EqualsString(other string) bool {
	return s.value == other
}

type ErrGame8MonsterPointValidation struct {
	Message string
}

func (e *ErrGame8MonsterPointValidation) Error() string {
	return e.Message
}
