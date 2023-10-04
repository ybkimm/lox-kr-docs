package lr2

type StateSet struct {
	states      []*ItemSet             // list of all states
	stateMap    map[string]int         // map of state key to state index
	transitions map[int]*TransitionMap // map of state to transitions
	actions     map[int]*ActionMap
}

func NewStateSet() *StateSet {
	return &StateSet{
		stateMap:    make(map[string]int),
		transitions: make(map[int]*TransitionMap),
		actions:     make(map[int]*ActionMap),
	}
}

func (c *StateSet) States() []*ItemSet {
	return c.states
}

func (c *StateSet) GetStateByKey(key string) (*ItemSet, int) {
	stateIndex, ok := c.stateMap[key]
	if !ok {
		return nil, 0
	}
	return c.states[stateIndex], stateIndex
}

func (c *StateSet) GetStateByIndex(stateIndex int) *ItemSet {
	return c.states[stateIndex]
}

func (c *StateSet) Add(key string, s *ItemSet) int {
	if _, ok := c.stateMap[key]; ok {
		panic("state already exists")
	}
	c.states = append(c.states, s)
	c.stateMap[key] = len(c.states) - 1
	return len(c.states) - 1
}

func (c *StateSet) Transitions(from int) *TransitionMap {
	ts := c.transitions[from]
	if ts == nil {
		ts = new(TransitionMap)
		c.transitions[from] = ts
	}
	return ts
}

func (c *StateSet) Actions(state int) *ActionMap {
	am := c.actions[state]
	if am == nil {
		am = new(ActionMap)
		c.actions[state] = am
	}
	return am
}
