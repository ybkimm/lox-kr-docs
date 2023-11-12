package mode

import (
	gotoken "go/token"
)

type ActionType int

const (
	ActionNone     ActionType = 0
	ActionPushMode ActionType = 1
	ActionPopMode  ActionType = 2
	ActionAccept   ActionType = 3
	ActionDiscard  ActionType = 4
	ActionAccum    ActionType = 5
)

type Action struct {
	Type     ActionType
	Terminal int
	Mode     string
}

type Actions struct {
	Actions []Action
	Pos     gotoken.Pos
}
