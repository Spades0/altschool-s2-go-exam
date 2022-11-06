package main

import "fmt"

type Car interface {
	start()      //start the car
	accelerate() //speed up
	stop()       //stop the car
}

var ()

type Product interface {
	showProduct()  //show a product and it's quantity
	showQuantity() //show the quantity of a product
	showPrice()    //show the price of a product
}

type Store interface {
	addItem()              //add an product to the store
	sellItem()             //sell a product
	showInventoryOfStock() //show all products in the store that have been sold or still in stock
	showSales()            //show the sold products and total sales
}

type car struct {
	name  string
	model string
	speed uint16
}

type product struct {
	car
	name     string
	price    uint
	quantity uint
}

type store struct {
	productList []struct {
		product
	}
	products     []string
	soldProducts []string
	totalSales   int
}

// Car methods
func (c car) start() {
	fmt.Println("This car is starting!")
}

func (c car) accelerate() {
	fmt.Printf("This car is accelerating at a speed of %v KM/h!\n", c.speed)
}

func (c car) stop() {
	fmt.Println("This car is stopping!")
}

// Product methods
func (p product) showProduct() {
	fmt.Println("The name of this product is: ", p.name)
	p.showQuantity()
}

func (p product) showQuantity() {
	if p.quantity <= 0 {
		fmt.Println("There is no quantity of this product. Quantity: ", p.quantity)
	} else {
		fmt.Println("Quantity: ", p.quantity)
	}
}

func (p product) showPrice() {
	fmt.Println("The price of this product is: ", p.price)
}

// Store method
func (s *store) addItem(p product) {
	fmt.Println("The ", p.name, " item has been added to the store.")
	s.products = append(s.products, p.name)
	// s.productList = append(s.productList, p)

}

func (s *store) sellItem(p *product) {
	for _, value := range s.products {
		if p.name == value {
			if p.quantity > 0 {
				p.quantity -= 1
				fmt.Println("1 ", p.name, " has been sold!", p.quantity, "left in the store.")
				s.soldProducts = append(s.soldProducts, p.name)
				s.totalSales += int(p.price)
			} else {
				fmt.Println("This product is out of stock!")
			}
		}
	}

}

func (s *store) showInventoryOfStock() {
	fmt.Println("The current inventory of the store is:")
	// for _, value2 := range s.products {
	// 	// // fmt.Println(value2.quantity)
	// 	// return value2.quantity
	// }
	// fmt.Println(s.products, p.showQuantity())
}

func (s *store) showSales() {
	fmt.Println("The sold items are: ")
	for _, value := range s.soldProducts {
		fmt.Printf("%v\n", value)
	}
	fmt.Printf("The total sale is: $%v", s.totalSales)
}

func main() {
	// initializing the car variable
	Car1 := new(product)
	Car1.name = "Mercedes-AMG GT 53"
	Car1.model = "Mercedes"
	Car1.speed = 125
	Car1.quantity = 3
	Car1.price = 250000
	// Car1.showProduct()

	Car2 := new(product)
	Car2.name = "BMW Ash"
	Car2.model = "BMW"
	Car2.speed = 140
	Car2.price = 120000
	Car2.quantity = 4

	Car2.showProduct()
	// Car2.showQuantity()
	// Car2.showPrice()

	altSchoolStore := new(store)
	// altSchoolStore.addItem(*Car1)
	altSchoolStore.addItem(*Car2)
	// altSchoolStore.showInventoryOfStock()
	// altSchoolStore.sellItem(Car2)
	// altSchoolStore.sellItem(Car2)
	// altSchoolStore.sellItem(Car1)
	// altSchoolStore.showSales()
	// Car1.start()
	// Car1.accelerate()
	// Car1.stop()

}
