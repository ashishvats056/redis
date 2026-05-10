package redis

type ValueType int

const (
	StringType ValueType = iota
	HashType
	ListType
)

type Entry struct {
	Type  ValueType
	Value any
}

type DB struct {
	data map[string]Entry
}

type Command struct {
	Name   string
	Args   []string
	Result chan Response
}

type Response struct {
	Data any
	Err  error
}

type CommandRequest struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type HandlerFunc func(*DB, Command)

var commands = map[string]HandlerFunc{
	"SET":  handleSet,
	"GET":  handleGet,
	"DEL":  handleDel,
	"MGET": handleMGet,
	// "HSET":   handleHSet,
	// "HGET":   handleHGet,
	// "LADD":   handleLAdd,
	// "LRANGE": handleLRange,
	// "INDEX":  handleIndex,
}
