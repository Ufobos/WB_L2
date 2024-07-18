package pattern

import "fmt"

// Команда (Command) — поведенческий паттерн проектирования, который превращает запросы в объекты,
// позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
// а также поддерживать отмену операций.

// Паттерн команда применяется:
// Когда надо передавать в качестве параметров определенные действия, вызываемые в ответ на другие действия.
// Когда необходимо обеспечить выполнение очереди запросов, а также их возможную отмену.

// Преимущества паттерна команда:
// Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
// Позволяет реализовать простую отмену и повтор операций.
// Позволяет реализовать отложенный запуск операций.
// Позволяет собирать сложные команды из простых.

// Недостаток паттерна команда это усложнение кода программы из-за введения множества дополнительных классов.

// Интерфейс команды
type Command interface {
	execute()
}

// Отправитель
type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.execute()
}

// Интерфейс получателя
type Device interface {
	on()
	off()
}

// Конкретная команда
type CommandOn struct {
	device Device
}

func (c *CommandOn) execute() {
	c.device.on()
}

// Конкретная команда
type CommandOff struct {
	device Device
}

func (c *CommandOff) execute() {
	c.device.off()
}

// Конкретный получатель
type Player struct {
	power bool
}

func (p *Player) on() {
	p.power = true
	fmt.Println("Power on")
}

func (p *Player) off() {
	p.power = false
	fmt.Println("Power off")
}

// клиентский код
// {
// Создаем экземпляр получателя.
// player = &Player
// Создаем экземпляры команд.
// commandOn = &CommandOn{player}
// commandOff = &CommandOff{player}
// Создаем экземпляры отправителей.
// buttonOn = Button{commandOn}
// buttonOff = Button{commandOff}
// buttonOn.execute() //"Power on"
//	buttonOff.execute() // "Power off"
// }
