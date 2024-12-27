package definitions

type TypedFactory[T any] interface {
	Factory
	TypedCreate(DependencyCollection) T
}
