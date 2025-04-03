package DATA

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

type DataManager struct {
	Items      map[int]Item
	Orders     map[int]Order
	ImageCache map[string][]byte
}

func (d *DataManager) DataInit() {
	d.ImageCache = make(map[string][]byte)
	d.Items = make(map[int]Item)
	d.Orders = make(map[int]Order)
}

func (d *DataManager) DropItems() {
	d.Items = make(map[int]Item)
}

func (d *DataManager) UpdateOrder(itemID int, newName string, newAmount int) bool {
	if order, exists := d.Orders[itemID]; exists {
		order.OrderItemName = newName
		order.OrderAmount = newAmount
		d.Orders[itemID] = order // Update the map with the modified order
		return true
	}
	return false
}

func (d *DataManager) AddOrder(itemID int, itemName string, amount int) {
	order := Order{OrderItemID: itemID, OrderItemName: itemName, OrderAmount: amount}
	d.Orders[itemID] = order
}

func (d *DataManager) GetOrder(itemID int) (*Order, bool) {
	order, exists := d.Orders[itemID]
	return &order, exists
}

func (d *DataManager) RemoveOrder(itemID int) bool {
	if _, exists := d.Orders[itemID]; exists {
		delete(d.Orders, itemID)
		return true
	}
	return false
}

func (d *DataManager) AddCart(itemID int, amount int) {
	// Check if the order already exists using the map
	if _, exists := d.Orders[itemID]; exists {
		if d.Items[itemID].ItemStock > d.Orders[itemID].OrderAmount && d.Items[itemID].ItemStock != 0 {
			// If it exists, update the order amount (increment it)
			d.UpdateOrder(itemID, d.Items[itemID].ItemName, d.Orders[itemID].OrderAmount+1)
		} else {
			println("Not Enough Stock")
		}

	} else {
		if d.Items[itemID].ItemStock != 0 {
			// If it doesn't exist, add the order with amount 1
			d.AddOrder(itemID, d.Items[itemID].ItemName, 1)
		}
	}
}

func (d *DataManager) SortItems(items map[int]Item, sortBy string) []Item {
	// Extract map values to a slice
	//("before sort --------")
	itemSlice := make([]Item, 0, len(items))
	for _, item := range items {
		itemSlice = append(itemSlice, item)
		//println(item.ItemName)
	}
	//println("after sort --------")
	// Sort based on the specified field
	switch sortBy {
	case "id":
		sort.Slice(itemSlice, func(i, j int) bool {
			return itemSlice[i].ItemID < itemSlice[j].ItemID
		})
	case "name":
		sort.Slice(itemSlice, func(i, j int) bool {
			return itemSlice[i].ItemName < itemSlice[j].ItemName
		})
	case "price":
		sort.Slice(itemSlice, func(i, j int) bool {
			return itemSlice[i].ItemPrice < itemSlice[j].ItemPrice
		})
	case "stock":
		sort.Slice(itemSlice, func(i, j int) bool {
			return itemSlice[i].ItemStock < itemSlice[j].ItemStock
		})
	default:
		fmt.Println("Invalid sort option")
	}

	// for item := range itemSlice {
	// 	println(itemSlice[item].ItemName)
	// }
	return itemSlice
}

func (d *DataManager) FetchAllProduct() bool {
	result, err := http.Get("http://127.0.0.1:8080/api/products")
	if err != nil {
		fmt.Println("Error fetching product list:", err)
		return false
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return false
	}

	var response ProductResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return false
	}

	for i, data := range response.Data {

		// Check if image is already cached.
		if imgData, exists := d.ImageCache[data.ImagePath]; exists {
			imgData = imgData
			//println("Loaded from cache:", len(imgData), "bytes")
		} else {
			// Fetch image from server.
			imgResult, err := http.Get("http://127.0.0.1:8080/" + data.ImagePath)
			if err != nil {
				fmt.Println("Error fetching image:", err)
				continue
			}
			defer imgResult.Body.Close()

			imgData, err := io.ReadAll(imgResult.Body)
			if err != nil {
				fmt.Println("Error reading image data:", err)
				continue
			}

			// Cache the image.
			d.ImageCache[data.ImagePath] = imgData
			//println("Image downloaded and cached:", len(imgData), "bytes")
		}

		d.Items[int(data.ID)] = Item{
			ItemID:        int(data.ID),
			ItemName:      data.ProductName,
			ItemPrice:     int(data.Price),
			ItemStock:     int(data.Stock.Quantity),
			ItemImagePath: data.ImagePath,
		}
		i = i
		//println(data.ProductName+" added to items", strconv.Itoa(i))
	}
	return true
}
