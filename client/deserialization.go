package client

type Listing struct {
	BooliId          int
	ConstructionYear int
	ListPrice        int
	LivingArea       float32
	Location         Location
	ObjectType       string
	Pageviews        int
	Published        string
	Rent             int
	Rooms            int
	Source           Source
	Url              string
}

type Address struct {
	StreetAddress string
}

type Region struct {
	MunicipalityName string
	CountyName       string
}

type Position struct {
	Latitude  float32
	Longitude float32
}

type Location struct {
	Address    Address
	NamedAreas []string
	Position   Position
	Region     Region
}

type Source struct {
	Name string
	Url  string
	Type string
}

type ListingsEnvelope struct {
	Listings []Listing
}
