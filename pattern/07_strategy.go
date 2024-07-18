package pattern

// Стратегия (Strategy) — поведенческий шаблон проектирования, предназначенный для определения семейства алгоритмов,
// инкапсуляции каждого из них и обеспечения их взаимозаменяемости. Это позволяет выбирать алгоритм путём определения
// соответствующего класса. Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов,
// которые его используют.

// Паттерн стратегия применяется:
// Когда есть несколько родственных классов, которые отличаются поведением. Можно задать один основной класс,
// а разные варианты поведения вынести в отдельные классы и при необходимости их применять.
// Когда необходимо обеспечить выбор из нескольких вариантов алгоритмов, которые можно легко менять
// в зависимости от условий.
// Когда необходимо менять поведение объектов на стадии выполнения программы.
// Когда класс, применяющий определенную функциональность, ничего не должен знать о ее реализации.

// Преимущества паттерна стратегия:
// Изолирует код и данные алгоритмов от остальных классов.
// Уход от наследования к делегированию.
// Горячая замена алгоритмов на лету.

// Недостатки паттерна стратегия:
// Усложняет программу за счёт дополнительных классов.
// Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

// Интерфейс стратегии.
type DBconnecter interface {
	Connect()
}

// Конкретная стратегия.
type PostgresConnection struct {
	connectionString string
}

func (p *PostgresConnection) Connect() {
	// Логика подключения к БД.
}

// Конкретная стратегия.
type MySQLConnection struct {
	connectionString string
}

func (m *MySQLConnection) Connect() {
	// Логика подключения к БД.
}

// Конкретная стратегия.
type SQLLiteConnection struct {
	connectionString string
}

func (s *SQLLiteConnection) Connect() {
	// Логика подключения к БД.
}

type DBConnection struct {
	db DBconnecter
}

// Контекст
func (con DBConnection) DBConnect() {
	con.db.Connect()
}

// Клиентский код
// {
// 	conn := new(DBConnection)
// 	postgresConnection := PostgresConnection{"postgresql://user:password@host:port/dbname"}
// 	conn.db = postgresConnection
// 	conn.DBConnect()
// 	mySQLConnection := MySQLConnection("mysql://user:password@host:port/dbname")
// 	conn.db = mySQLConnection
// 	conn.DBConnect()
// }
