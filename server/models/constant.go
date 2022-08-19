package models

const (
	// user
	ROOT = iota
	Admin
	None
	// community
	ACTIVE
	BANNED
	WAITING
	END
	// appliactions type
	NEW
	ADD
	DEL
	INFO
	ALL

	//commit
	PASS
	CHANGE

	//
	ERR
)