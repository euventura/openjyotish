package domain

type Varga struct {
	Name     string
	Division int
	Grahas   []Graha
	Lagna    Bhava
}

type IVargaCalc interface {
	Name() string
	Division() int
	Calculate(info *VargaCalc) (Varga, error)
}

type VargaCalc struct {
	Kundli Kundli
}
