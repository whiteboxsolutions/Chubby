package chubby

import (
	"fmt"
)

type Roll struct {
	Value uint
}

type Rolls map[string]Roll

func (r Roll) String() string {
	return fmt.Sprintf("Value: %d", r.Value)
}

func New() Rolls {
	return make(map[string]Roll)
}

func (r Rolls) NewRoll(name string) Roll {
	r[name] = Roll{Value: setBit(0, uint(len(r)))}
	return r[name]
}

func HasRoll(roll uint, requirement Roll) bool {
	return roll^requirement.Value <= roll
}

func (r Rolls) Get(name string) (Roll, error) {
	if val, ok := r[name]; ok {
		return val, nil
	}
	return Roll{}, fmt.Errorf(`no roll found with name "%s"`, name)
}

func (r Rolls) Combine(names ...string) uint {
	var n uint = 0
	for _, v := range names {
		if val, ok := r[v]; ok {
			n += val.Value
		}
	}
	return n
}

func setBit(n uint, pos uint) uint {
	n |= (1 << pos)
	return n
}
