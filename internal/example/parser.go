package main

import (
	"fmt"

	lockstep "github.com/iStreamPlanet/go-lockstep/lockstep"
)

type Foobar struct {
	lockstep *lockstep.LockStep
}

type Parser struct {
}

func (p *Parser) ReduceExpressionPlus(e1 any, e3 any) any {
	panic("not-implemented")
}

func main() {
	fmt.Println("hey")
}
