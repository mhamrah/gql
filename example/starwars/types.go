// Code generated by thriftrw
// @generated

package starwars

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type Character struct {
	ID        string       `json:"id"`
	Name      *string      `json:"name,omitempty"`
	Friends   []*Character `json:"friends"`
	AppearsIn []Episode    `json:"appearsIn"`
}

type _List_Character_ValueList []*Character

func (v _List_Character_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_Character_ValueList) Size() int {
	return len(v)
}

func (_List_Character_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_Character_ValueList) Close() {
}

type _List_Episode_ValueList []Episode

func (v _List_Episode_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_Episode_ValueList) Size() int {
	return len(v)
}

func (_List_Episode_ValueList) ValueType() wire.Type {
	return wire.TI32
}

func (_List_Episode_ValueList) Close() {
}

func (v *Character) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.ID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Name != nil {
		w, err = wire.NewValueString(*(v.Name)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.Friends != nil {
		w, err = wire.NewValueList(_List_Character_ValueList(v.Friends)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.AppearsIn != nil {
		w, err = wire.NewValueList(_List_Episode_ValueList(v.AppearsIn)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Character_Read(w wire.Value) (*Character, error) {
	var v Character
	err := v.FromWire(w)
	return &v, err
}

func _List_Character_Read(l wire.ValueList) ([]*Character, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*Character, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Character_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func _Episode_Read(w wire.Value) (Episode, error) {
	var v Episode
	err := v.FromWire(w)
	return v, err
}

func _List_Episode_Read(l wire.ValueList) ([]Episode, error) {
	if l.ValueType() != wire.TI32 {
		return nil, nil
	}
	o := make([]Episode, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Episode_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *Character) FromWire(w wire.Value) error {
	var err error
	idIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.ID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				idIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Name = &x
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TList {
				v.Friends, err = _List_Character_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TList {
				v.AppearsIn, err = _List_Episode_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	if !idIsSet {
		return errors.New("field ID of Character is required")
	}
	return nil
}

func (v *Character) String() string {
	var fields [4]string
	i := 0
	fields[i] = fmt.Sprintf("ID: %v", v.ID)
	i++
	if v.Name != nil {
		fields[i] = fmt.Sprintf("Name: %v", *(v.Name))
		i++
	}
	if v.Friends != nil {
		fields[i] = fmt.Sprintf("Friends: %v", v.Friends)
		i++
	}
	if v.AppearsIn != nil {
		fields[i] = fmt.Sprintf("AppearsIn: %v", v.AppearsIn)
		i++
	}
	return fmt.Sprintf("Character{%v}", strings.Join(fields[:i], ", "))
}

type Droid struct {
	ID              string       `json:"id"`
	Name            *string      `json:"name,omitempty"`
	Friends         []*Character `json:"friends"`
	AppearsIn       []Episode    `json:"appearsIn"`
	PrimaryFunction *string      `json:"primary_function,omitempty"`
}

func (v *Droid) ToWire() (wire.Value, error) {
	var (
		fields [5]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.ID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Name != nil {
		w, err = wire.NewValueString(*(v.Name)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.Friends != nil {
		w, err = wire.NewValueList(_List_Character_ValueList(v.Friends)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.AppearsIn != nil {
		w, err = wire.NewValueList(_List_Episode_ValueList(v.AppearsIn)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.PrimaryFunction != nil {
		w, err = wire.NewValueString(*(v.PrimaryFunction)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Droid) FromWire(w wire.Value) error {
	var err error
	idIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.ID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				idIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Name = &x
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TList {
				v.Friends, err = _List_Character_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TList {
				v.AppearsIn, err = _List_Episode_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.PrimaryFunction = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if !idIsSet {
		return errors.New("field ID of Droid is required")
	}
	return nil
}

func (v *Droid) String() string {
	var fields [5]string
	i := 0
	fields[i] = fmt.Sprintf("ID: %v", v.ID)
	i++
	if v.Name != nil {
		fields[i] = fmt.Sprintf("Name: %v", *(v.Name))
		i++
	}
	if v.Friends != nil {
		fields[i] = fmt.Sprintf("Friends: %v", v.Friends)
		i++
	}
	if v.AppearsIn != nil {
		fields[i] = fmt.Sprintf("AppearsIn: %v", v.AppearsIn)
		i++
	}
	if v.PrimaryFunction != nil {
		fields[i] = fmt.Sprintf("PrimaryFunction: %v", *(v.PrimaryFunction))
		i++
	}
	return fmt.Sprintf("Droid{%v}", strings.Join(fields[:i], ", "))
}

type Episode int32

const (
	EpisodeNewhope Episode = 0
	EpisodeEmpire  Episode = 1
	EpisodeJedi    Episode = 2
)

func (v Episode) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *Episode) FromWire(w wire.Value) error {
	*v = (Episode)(w.GetI32())
	return nil
}

func (v Episode) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "NEWHOPE"
	case 1:
		return "EMPIRE"
	case 2:
		return "JEDI"
	}
	return fmt.Sprintf("Episode(%d)", w)
}

type Human struct {
	ID         string       `json:"id"`
	Name       *string      `json:"name,omitempty"`
	Friends    []*Character `json:"friends"`
	AppearsIn  []Episode    `json:"appearsIn"`
	HomePlanet *string      `json:"home_planet,omitempty"`
}

func (v *Human) ToWire() (wire.Value, error) {
	var (
		fields [5]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.ID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Name != nil {
		w, err = wire.NewValueString(*(v.Name)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.Friends != nil {
		w, err = wire.NewValueList(_List_Character_ValueList(v.Friends)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.AppearsIn != nil {
		w, err = wire.NewValueList(_List_Episode_ValueList(v.AppearsIn)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.HomePlanet != nil {
		w, err = wire.NewValueString(*(v.HomePlanet)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Human) FromWire(w wire.Value) error {
	var err error
	idIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.ID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				idIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Name = &x
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TList {
				v.Friends, err = _List_Character_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TList {
				v.AppearsIn, err = _List_Episode_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.HomePlanet = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if !idIsSet {
		return errors.New("field ID of Human is required")
	}
	return nil
}

func (v *Human) String() string {
	var fields [5]string
	i := 0
	fields[i] = fmt.Sprintf("ID: %v", v.ID)
	i++
	if v.Name != nil {
		fields[i] = fmt.Sprintf("Name: %v", *(v.Name))
		i++
	}
	if v.Friends != nil {
		fields[i] = fmt.Sprintf("Friends: %v", v.Friends)
		i++
	}
	if v.AppearsIn != nil {
		fields[i] = fmt.Sprintf("AppearsIn: %v", v.AppearsIn)
		i++
	}
	if v.HomePlanet != nil {
		fields[i] = fmt.Sprintf("HomePlanet: %v", *(v.HomePlanet))
		i++
	}
	return fmt.Sprintf("Human{%v}", strings.Join(fields[:i], ", "))
}

type IdRequest struct {
	ID string `json:"id"`
}

func (v *IdRequest) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.ID), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *IdRequest) FromWire(w wire.Value) error {
	var err error
	idIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.ID, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				idIsSet = true
			}
		}
	}
	if !idIsSet {
		return errors.New("field ID of IdRequest is required")
	}
	return nil
}

func (v *IdRequest) String() string {
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("ID: %v", v.ID)
	i++
	return fmt.Sprintf("IdRequest{%v}", strings.Join(fields[:i], ", "))
}
