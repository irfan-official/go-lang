package main


type Myne struct {
	Name string
	Age  int
}

func Check(theme Myne) Myne {
	return Myne{
		Name: theme.Name,
		Age:  theme.Age,
	}
}
