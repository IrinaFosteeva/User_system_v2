package model

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonUpdate struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}
