package reqres

type Mrequest struct{
	Url string
	Method string
}

type Detailres struct{
	Name string
	Url string
	Description string
}

type Jsonres struct {
	Class string
	Atom []Jsonresslice

}

type Jsonresslice struct {
	Name string
	Url string
	Dec string
	Class string
}

type Bigjson struct {
	Res []Jsonres
}