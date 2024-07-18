package pattern

// Строитель (Builder) - шаблон проектирования, который инкапсулирует создание объекта
// и позволяет разделить его на различные этапы.

// Шаблон проектирования строитель состоить из 4 компонентов:
// Product: представляет объект, который должен быть создан.
// В данном случае все части объекта заключены в списке parts.
// Builder: определяет интерфейс для создания различных частей объекта Product.
// ConcreteBuilder: конкретная реализация Buildera.
// Создает объект Product и определяет интерфейс для доступа к нему.
// Director: распорядитель - создает объект, используя объекты Builder.

// Паттер строитель применяется когда процесс создания нового объекта не должен зависеть от того,
// из каких частей этот объект состоит и как эти части связаны между собой.
// Строитель применяется когда необходимо обеспечить получение различных вариаций объекта в процессе его создания.

// Паттер строитель позволяет изменять внутреннее представление продукта.
// Объект Builder предоставляет распорядителю абстрактный интерфейс для конструирования продукта,
// за которым он может скрыть представление и внутреннюю структуру продукта, а также процесс его сборки.
// Поскольку продукт конструируется через абстрактный интерфейс, то для изменения внутреннего представления
// достаточно всего лишь определить новый вид строителя.

// Строитель изолирует код, реализующий конструирование и представление.
// Улучшается модульность, инкапсулируя способ конструирования и представления сложного объекта.

// Строитель предоставляет более точный контроль над процессом конструирования.
// В отличие от порождающих паттернов, которые сразу конструируют весь объект целиком,
// builder делает это шаг за шагом под управлением director. Когда продукт завершен, director забирает его у builder.

// Строители конструируют свои продукты шаг за шагом, поэтому интерфейс класса Builder должен быть достаточно общим,
// чтобы обеспечить конструирование при любом виде конкретного строителя.

type Computer struct {
	CPU string
	RAM int
	SSD int
}

type IComputerBuilder interface {
	SetCPU(cpu string) IComputerBuilder
	SetRAM(ram int) IComputerBuilder
	SetSSD(ssd int) IComputerBuilder
	Build() Computer
}

type ComputerBuilder struct {
	cpu string
	ram int
	ssd int
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

func (c *ComputerBuilder) SetCPU(cpu string) IComputerBuilder {
	c.cpu = cpu
	return c
}
func (c *ComputerBuilder) SetRAM(ram int) IComputerBuilder {
	c.ram = ram
	return c
}
func (c *ComputerBuilder) SetSSD(ssd int) IComputerBuilder {
	c.ssd = ssd
	return c
}

func (c *ComputerBuilder) Build() Computer {
	return Computer{
		c.cpu,
		c.ram,
		c.ssd,
	}
}

type Director struct {
	builder IComputerBuilder
}

func NewDirector(computerBuilder IComputerBuilder) *Director {
	return &Director{builder: computerBuilder}
}

func (d *Director) SetBuilder(computerBuilder IComputerBuilder) {
	d.builder = computerBuilder
}
