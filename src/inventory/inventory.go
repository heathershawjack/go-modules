package inventory

type Inventory struct {
	MaxCapacity int
	SpaceFree   int
	Item        []Item
}

type Item struct {
	Id   int
	Name string
}

func (i Inventory) addItem(id int, name string) {
	newItem := Item{Id: id, Name: name}
	i.Item = append(i.Item, newItem)
}
