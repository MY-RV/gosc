package gosc

type ServiceFactoryFunction[T any] func(s *ServiceCollection) *T
