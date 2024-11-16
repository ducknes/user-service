package domain

type Role uint

const (
	Undefined Role = iota
	UserRole
	AdminRole
)
