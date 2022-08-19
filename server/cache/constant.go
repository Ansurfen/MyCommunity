package cache

const (
	// cacheSize options
	_byte = 1
	b     = _byte * 8
	kb    = b * 1024
	mb    = kb * 1024
	gb    = mb * 1024
	// expire options
	Sec  = 1
	Min  = Sec * 60
	Hour = Min * 60
	Day  = Hour * 24
	Week = Day * 7
)
