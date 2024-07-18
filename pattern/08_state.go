package pattern

// Состояние (State) — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.

// Паттерн состояние применяется:
// Когда поведение объекта должно зависеть от его состояния и может изменяться динамически во время выполнения.
// Когда в коде методов объекта используются многочисленные условные конструкции, выбор которых зависит от
// текущего состояния объекта.

// Преимущества паттерна состояние:
// Избавляет от множества больших условных операторов машины состояний.
// Концентрирует в одном месте код, связанный с определённым состоянием.
// Упрощает код контекста.

// Паттерн состояние может неоправданно усложнить код, если состояний мало и они редко меняются.

// Интерфейс состояния
type State interface {
	MethodA()
	MethodB()
}

// Контекст
type ContextState struct {
	StateA       State
	StateB       State
	StateC       State
	CurrentState State
}

func (c *ContextState) SetState(s State) {
	c.CurrentState = s
}

func newContextState() *ContextState {
	contextState := &ContextState{}

	stateA := &StateA{contextState}
	stateB := &StateB{contextState}
	stateC := &StateC{contextState}

	contextState.StateA = stateA
	contextState.StateB = stateB
	contextState.StateC = stateC

	contextState.SetState(stateA)

	return contextState
}

// Конкретный интерфейс
type StateA struct {
	contextState *ContextState
}

func (s *StateA) MethodA() {
	// необходимая логика
}

func (s *StateA) MethodB() {
	// необходимая логика
}

// Конкретный интерфейс
type StateB struct {
	contextState *ContextState
}

func (s *StateB) MethodA() {
	// необходимая логика
}

func (s *StateB) MethodB() {
	// необходимая логика
}

// Конкретный интерфейс
type StateC struct {
	contextState *ContextState
}

func (s *StateC) MethodA() {
	// необходимая логика
}

func (s *StateC) MethodB() {
	// необходимая логика
}
