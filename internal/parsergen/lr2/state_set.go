package lr2

type StateSet struct {
	states      []*ItemSet             // list of all states
	stateMap    map[string]int         // map of state key to state index
	transitions map[int]*TransitionMap // map of state to transitions
}

func NewStateSet() *StateSet {
	return &StateSet{
		stateMap:    make(map[string]int),
		transitions: make(map[int]*TransitionMap),
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

func (c *StateSet) Add(key string, s *ItemSet) {
	if _, ok := c.stateMap[key]; ok {
		panic("state already exists")
	}
	c.states = append(c.states, s)
	c.stateMap[key] = len(c.states) - 1
}

func (c *StateSet) AddTransition(from, symbol, to int) {
	transMap := c.transitions[from]
	if transMap == nil {
		transMap = new(TransitionMap)
		c.transitions[from] = transMap
	}
	transMap.Add(symbol, to)
}
