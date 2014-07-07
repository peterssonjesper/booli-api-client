package client

type Listing struct {
	BooliId          int
	ConstructionYear int
	EnableStreetView int
	Floor            int
	HasImages        int
	HasShowings      int
	ListPrice        int
	LivingArea       float32
	Location         Location
	MonthlyPayment   float32
	ObjectType       string
	OperatingCost    int
	Pageviews        int
	Published        string
	Rent             int
	Rooms            float32
	SecondOpinion    int
	ShortDesc        string
	Source           Source
	Thumb            Image
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
	Distance   Distance
	NamedAreas []string
	Position   Position
	Region     Region
}

type Source struct {
	Id   int
	Name string
	Url  string
	Type string
	Logo Image
}

type Distance struct {
	Ocean int
}

type ListingsEnvelope struct {
	Listings []Listing
}

type Image struct {
	Width  int
	Height int
	Id     int
}
