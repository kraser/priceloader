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
	Price       int64
	PriceRur    int64
	PriceUsd    int64
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

func (pCat *Category) AddItem(pItem *Item) {
	pCat.Items[len(pCat.Items)+1] = pItem
}

type Price struct {
	SupplierId    int32
	supplierCode  string
	Categories    map[string]*Category
	CategoryStack []*Category
	curLevel      int
}

type LoadTask struct {
	Pointer *Category
	Handler func()
	Message string
}

var PriceList = new(Price)
var pCurrentCategory *Category

func (price *Price) PriceList(supplierCode string) {
	price.supplierCode = supplierCode
	price.Categories = make(map[string]*Category)
	price.CategoryStack = make([]*Category, 8)
}

func (price *Price) SetCurrentCategory(name string, url string, level int) *Category {
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

		category := Category{Id: 0, Name: name, URL: url}
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

func (price *Price) AddItem(pCategory *Category, pItem *Item) {
	pCategory.AddItem(pItem)
}
