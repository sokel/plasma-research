package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type levelDB struct {
	instance *leveldb.DB   // LevelDB instance
	path string            // Path to database
}

// newLDBDatabase returns a LevelDB wrapped object.
func NewLDBDatabase(file string, cache int, handles int, blockSize int) (*levelDB, error) {
	// Ensure we have some minimal caching and file guarantees
	if cache < 16 {
		cache = 16
	}
	if handles < 16 {
		handles = 16
	}

	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB,
		Filter:                 filter.NewBloomFilter(10),
		BlockSize: blockSize * opt.KiB,
	})

	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	// (Re)check for errors and abort if opening of the db failed
	if err != nil {
		return nil, err
	}

	return &levelDB{
		instance: db,
		path: file,
	}, nil
}

// Path returns the path to the database directory.
func (db *levelDB) Path() string {
	return db.path
}

// Put puts the given key / value to the queue
func (db *levelDB) Put(key []byte, value []byte) error {
	return db.instance.Put(key, value, nil)
}

// Has check for existence of key
func (db *levelDB) Has(key []byte) (bool, error) {
	return db.instance.Has(key, nil)
}

// Get returns the value of given key if it presents
func (db *levelDB) Get(key []byte) ([]byte, error) {
	dat, err := db.instance.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// Delete deletes the key from the queue and database
func (db *levelDB) Delete(key []byte) error {
	return db.instance.Delete(key, nil)
}

type table struct {
	db     levelDB
	prefix string
}

// NewTable returns a Database object that prefixes all keys with a given string
func NewTable(db levelDB, prefix string) *table {
	return &table{
		db:  db,
		prefix: prefix,
	}
}

func (dt *table) Put(key []byte, value []byte) error {
	return dt.db.Put(append([]byte(dt.prefix), key...), value)
}

func (dt *table) Has(key []byte) (bool, error) {
	return dt.db.Has(append([]byte(dt.prefix), key...))
}

func (dt *table) Get(key []byte) ([]byte, error) {
	return dt.db.Get(append([]byte(dt.prefix), key...))
}

func (dt *table) Delete(key []byte) error {
	return dt.db.Delete(append([]byte(dt.prefix), key...))
}

// NewBatch new data aggregator which will insert into leveldb in future
func (db *levelDB) NewBatch() *ldbBatch {
	return &ldbBatch{instance: db.instance, b: new(leveldb.Batch)}
}

type ldbBatch struct {
	instance   *leveldb.DB
	b    *leveldb.Batch
	size int
}

func (b *ldbBatch) Put(key, value []byte) error {
	b.b.Put(key, value)
	b.size += len(value)
	return nil
}

func (b *ldbBatch) Delete(key []byte) error {
	b.b.Delete(key)
	b.size += 1
	return nil
}

func (b *ldbBatch) Write() error {
	return b.instance.Write(b.b, nil)
}

func (b *ldbBatch) ValueSize() int {
	return b.size
}

func (b *ldbBatch) Reset() {
	b.b.Reset()
	b.size = 0
}

