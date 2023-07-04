package lr1

type StateSet struct {
	stateMap map[string]*State
	states   []*State
	changed  bool
}

func NewStateSet() *StateSet {
	return &StateSet{
		stateMap: make(map[string]*State),
	}
}

func (c *StateSet) Changed() bool {
	return c.changed
}

func (c *StateSet) ResetChanged() {
	c.changed = false
}

func (c *StateSet) Add(s *State) *State {
	if existing, ok := c.stateMap[s.Key]; ok {
		return existing
	}
	c.changed = true
	s.Index = len(c.states)
	c.states = append(c.states, s)
	c.stateMap[s.Key] = s
	return s
}

func (c *StateSet) ForEach(fn func(s *State)) {
	for _, state := range c.states {
		fn(state)
	}
}
