package inventory

type Inventory struct {
	MaxCapacity int    `json:maxCapacity`
	SpaceFree   int    `json:spaceFree`
	Item        []Item `json:item`
}

type Item struct {
	Id   int    `json:id`
	Name string `json:name`
}

func (i Inventory) addItem(id int, name string) {
	newItem := Item{Id: id, Name: name}
	i.Item = append(i.Item, newItem)
}
