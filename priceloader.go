// priceloader project priceloader.go
package priceloader

//"fmt"
//"logger"

type Item struct {
	Id          int32
	SupplierId  int32
	ReferenceId int32
	Name        string
	Code        string
	Model       string
	Brand       string
	Price       float64
	PriceRur    float64
	PriceUsd    float64
	Store       string
	StoreNsk    string
	StoreMsk    string
}

type Category struct {
	Id         int32
	Name       string
	URL        string
	Categories map[string]*Category
	Items      map[int]*Item
}

type Price struct {
	SupplierId    int32
	supplierCode  string
	Categories    map[string]*Category
	CategoryStack []*Category
	curLevel      int
}

var PriceList = new(Price)
var pCurrentCategory *Category

func (price *Price) PriceList(supplierCode string) {
	price.supplierCode = supplierCode
	price.Categories = make(map[string]*Category)
	price.CategoryStack = make([]*Category, 8)
}

func (price *Price) SetCurrentCategory(name string, level int) *Category {
	if level < 0 {
		level = 0
	}

	if level == 0 {
		pCurrentCategory = price.createAndAddCategory(name)
	} else {
		var pCurCategory *Category
		if level > price.curLevel {
			pCurCategory = price.CategoryStack[price.curLevel]
		} else {
			pCurCategory = price.CategoryStack[level-1]
		}

		category := Category{Id: 0, Name: name, URL: ""}
		category.Categories = make(map[string]*Category)
		category.Items = make(map[int]*Item)
		pCurrentCategory := &category
		pCurCategory.Categories[name] = pCurrentCategory
		PriceList.CategoryStack[level] = pCurrentCategory

	}
	price.curLevel = level
	return pCurrentCategory
}

func (price *Price) createAndAddCategory(name string) *Category {
	category := Category{Id: 0, Name: name, URL: ""}
	category.Categories = make(map[string]*Category)
	category.Items = make(map[int]*Item)
	pCategory := &category
	price.CategoryStack[0] = pCategory
	price.Categories[name] = pCategory
	return pCategory
}
