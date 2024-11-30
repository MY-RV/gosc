package pkg

func GetService[T any](serviceCollection *ServiceCollection) T {
	service, err := TryGetService[T](serviceCollection)

	if err != nil {
		panic(err)
	}

	return service
}
