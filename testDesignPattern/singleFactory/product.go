package singleFactory

// 抽象的产品

type Product interface {
	SetName(name string)
	GetName() string
}

// Product1
type Product1 struct {
	name string
}

func (p *Product1) SetName(name string) {
	p.name = name
}

func (p *Product1) GetName() string {
	return "产品1 的名字是" + p.name
}

// Product2
type Product2 struct {
	name string
}

func (p *Product2) SetName(name string) {
	p.name = name
}

func (p *Product2) GetName() string {
	return "产品2 的名字是" + p.name
}

type ProductType int

const (
	p1 = iota
	p2
)

type ProductFactory struct {
}

func (p ProductFactory) Create(productType ProductType) Product {
	if productType == 0 {
		return new(Product1)
	}
	if productType == 1 {
		return new(Product1)
	}

	return nil
}
