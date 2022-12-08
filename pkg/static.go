package pkg

type Static interface {
	Apply(ctrl Controller)
}

type StaticWithDescription interface {
	DisplayName() string
	Description() string
}
