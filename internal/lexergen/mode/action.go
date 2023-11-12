package mode

import (
	gotoken "go/token"
)

type ActionType int

const (
	ActionNone ActionType = iota
	ActionEmit
	ActionPushMode
	ActionPopMode
	ActionSkip
)

type Action struct {
	Type     ActionType
	Terminal int
	Pos      gotoken.Pos
	Mode     string
}
