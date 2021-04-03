package picker

import (
	"strings"

	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultModifier = ModifierNone

type Modifier string

func (m Modifier) String() string {
	return string(m)
}

func (m *Modifier) toLower() Modifier {
	*m = Modifier(strings.ToLower(m.String()))
	return *m
}

func (m *Modifier) IsValid() bool {
	mod := m.toLower()
	for _, v := range modifierValues {
		if mod == v {
			return true
		}
	}
	return false
}

const (
	ModifierUnspecified Modifier = ""
	// Do not apply any multiplier to the field value
	ModifierNone Modifier = "none"

	// Take the common logarithm of the field value. Because this function will return a negative value and cause an error if used on values between 0 and 1, it is recommended to use log1p instead.
	ModifierLog Modifier = "log"

	//	Add 1 to the field value and take the common logarithm
	ModifierLog1P Modifier = "log1p"

	//Add 2 to the field value and take the common logarithm
	ModifierLog2P Modifier = "log2p"

	// Take the natural logarithm of the field value. Because this function will return a negative value and cause an error if used on values between 0 and 1, it is recommended to use ln1p instead.
	ModifierLn Modifier = "ln"

	// Add 1 to the field value and take the natural logarithm
	ModifierLn1P Modifier = "ln1p"

	// Add 2 to the field value and take the natural logarithm
	ModifierLn2P Modifier = "ln2p"

	// Square the field value (multiply it by itself)
	ModifierSquare Modifier = "square"

	// Take the square root of the field value
	ModifierSqrt Modifier = "sqrt"

	//Reciprocate the field value, same as 1/x where x is the field’s value
	ModifierReciprocal Modifier = "reciprocal"
)

var modifierValues = []Modifier{
	ModifierNone,
	ModifierLog,
	ModifierLog1P,
	ModifierLog2P,
	ModifierLn,
	ModifierLn1P,
	ModifierLn2P,
	ModifierSquare,
	ModifierSqrt,
	ModifierReciprocal,
}

type modifierParam struct {
	modifier Modifier
}

type WithModifier interface {
	SetModifier(m Modifier) error
	Modifier() Modifier
}

func (m *modifierParam) SetModifier(modifier Modifier) error {
	if !modifier.IsValid() {
		return ErrInvalidModifier
	}
	m.modifier = modifier
	return nil
}
func (m *modifierParam) Modifier() Modifier {
	if m.modifier == ModifierUnspecified {
		return DefaultModifier
	}
	return m.modifier
}

func marshalModifierParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithModifier); ok {
		if b.Modifier() != DefaultModifier {
			return json.Marshal(b.Modifier())
		}
	}
	return nil, nil
}
func unmarshalModifierParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithModifier); ok {
		var m Modifier
		err := json.Unmarshal(data, &m)
		if err != nil {
			return err
		}
		return a.SetModifier(m)
	}
	return nil
}
