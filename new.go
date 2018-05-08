package main

//new allocates zeroed storage for a new item or type whatever and then returns a pointer to it.
//I don't think it really matters on if you use new vs short variable declaration := type{} it's mostly just preference

type SomeUser struct {
	Name string `json:"name" form:"name"`
}

func newTest() {
	//u := new(SomeUser)
}
