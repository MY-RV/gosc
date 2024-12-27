package scope

type Enum int

const (
	SINGLETON Enum = iota
	TRANSIENT
	SCOPED
)
