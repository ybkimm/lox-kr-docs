package lr1

type StateSet struct {
	stateMap map[string]*ItemSet
	states   []*ItemSet
}

func NewStateSet() *StateSet {
	return &StateSet{
		stateMap: make(map[string]*ItemSet),
	}
}

func (c *StateSet) States() []*ItemSet {
	return c.states
}

func (c *StateSet) Get(key string) *ItemSet {
	return c.stateMap[key]
}

func (c *StateSet) Add(key string, s *ItemSet) {
	if _, ok := c.stateMap[key]; ok {
		panic("state already exists")
	}
	s.Index = len(c.states)
	c.states = append(c.states, s)
	c.stateMap[key] = s
}

func (c *StateSet) ForEach(fn func(s *ItemSet)) {
	for _, state := range c.states {
		fn(state)
	}
}
