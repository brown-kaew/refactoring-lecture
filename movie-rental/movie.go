package rental

const (
	_ = iota
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Movie struct {
	title     string
	priceCode int
	Pricer    Pricer
}

type Pricer interface {
	Charge(daysRented int) float64
	PriceCode() int
}

type Regular struct {
	priceCode int
}

func (r Regular) Charge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}

func (r Regular) PriceCode() int {
	return r.priceCode
}

func CreateRegular() Pricer {
	return Regular{
		priceCode: REGULAR,
	}
}

type NewRelease struct {
	priceCode int
}

func (r NewRelease) Charge(daysRented int) float64 {
	return float64(daysRented) * 3.0
}

func (r NewRelease) PriceCode() int {
	return r.priceCode
}

func CreateNewRelease() Pricer {
	return NewRelease{
		priceCode: NEW_RELEASE,
	}
}

type Childrens struct {
	priceCode int
}

func (r Childrens) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

func (r Childrens) PriceCode() int {
	return r.priceCode
}

func CreateChildrens() Pricer {
	return Childrens{
		priceCode: CHILDRENS,
	}
}

func New(title string, pricer Pricer) Movie {
	return Movie{
		title:     title,
		priceCode: pricer.PriceCode(),
		Pricer:    pricer,
	}
}

func NewMovie(title string, priceCode int) Movie {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}
func (m Movie) PriceCode() int {
	return m.priceCode
}
func (m Movie) Title() string {
	return m.title
}
func (m Movie) SetPriceCode(arg int) {
	m.priceCode = arg
}
