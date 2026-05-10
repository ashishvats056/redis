package redis

import "time"

func isExpired(e Entry) bool {
	if e.ExpiresAt == 0 {
		return false
	}
	return time.Now().UnixNano() > e.ExpiresAt
}

func startCleaner(db *DB) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			for k, v := range db.data {
				if isExpired(v) {
					delete(db.data, k)
				}
			}
		}
	}()
}

func getEntry(db *DB, key string) (Entry, bool) {
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
