package inventory

type Inventory struct {
	MaxCapacity int    `json:"maxCapacity"`
	Item        []Item `json:"item"`
}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (i Inventory) addItem(id int, name string) {
	i.Item = append(i.Item, Item{Id: id, Name: name})
}

func (i Inventory) getFreeSpace() int {
	return (len(i.Item) - i.MaxCapacity)
}
