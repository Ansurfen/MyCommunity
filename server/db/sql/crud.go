package sql

func (db *DB) Insert(value any) *DB {
	db.Create(value)
	return db
}

func (db *DB) InsertMany(values ...any) *DB {
	for _, value := range values {
		db.Create(value)
	}
	return db
}

func (db *DB) Delete(value any, where ...any) *DB {
	db.DB.Delete(value, where...)
	return db
}

func (db *DB) Update(value any) *DB {
	db.Save(value)
	return db
}

func (db *DB) Query(value any, where ...any) *DB {
	db.First(value, where...)
	return db
}