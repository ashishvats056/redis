package redis

type ValueType int

const (
	StringType ValueType = iota
	HashType
	ListType
)

type Entry struct {
	Type      ValueType
	Value     any
	ExpiresAt int64 // unix nano, 0 = never expires
}

type DB struct {
	data map[string]Entry
}

func NewDB() DB {
	return DB{
		data: make(map[string]Entry),
	}
}

func (db *DB) Set(key string, value any, expiresAt int64, entryType ValueType) {

	db.data[key] = Entry{
		Type:      entryType,
		Value:     value,
		ExpiresAt: expiresAt,
	}
}

func (db *DB) Get(key string) (Entry, bool) {
	entry, ok := db.data[key]
	if !ok {
		return Entry{}, false
	}

	if isExpired(entry) {
		delete(db.data, key)
		return Entry{}, false
	}

	return entry, true
}

func (db *DB) Delete(key string) {
	delete(db.data, key)
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

	"HSET": handleHSet,
	"HGET": handleHGet,

	"LADD":   handleLAdd,
	"LRANGE": handleLRange,
	"INDEX":  handleIndex,

	"EXPIRE":  handleExpire,
	"TTL":     handleTTL,
	"PERSIST": handlePersist,
}
