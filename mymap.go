package bootcamp

type MyMap struct {
	mp []struct {
		Key   string
		Value interface{}
	}
}

func (m *MyMap) Set(k string, v interface{}) {
	for i := range m.mp {
		if m.mp[i].Key == k {
			m.mp[i].Value = v
			return
		}
	}
	m.mp = append(m.mp, struct {
		Key   string
		Value interface{}
	}{k, v})
}

func (m *MyMap) Get(k string) interface{} {
	for i := range m.mp {
		if m.mp[i].Key == k {
			return m.mp[i].Value
		}
	}
	return nil
}

func (m *MyMap) Has(k string) bool {
	return m.Get(k) != nil
}

func (m *MyMap) Delete(k string) {
	for i := range m.mp {
		if m.mp[i].Key == k {
			m.mp = append(m.mp[:i], m.mp[i+1:]...)
			return
		}
	}
}

func (m *MyMap) Items() []struct {
	Key   string
	Value interface{}
} {
	return m.mp
}

// func main() {
// 	myMap := MyMap{}

// 	myMap.Set("key1", 42)
// 	myMap.Set("key2", "value2")

// 	fmt.Println(myMap.Get("key1")) // 42
// 	fmt.Println(myMap.Get("key2")) // value2
// 	fmt.Println(myMap.Has("key1")) // true
// 	fmt.Println(myMap.Has("key3")) // false

// 	myMap.Delete("key2")
// 	fmt.Println(myMap.Has("key2")) // false

// 	items := myMap.Items()
// 	for _, item := range items {
// 		fmt.Printf("Key: %s, Value: %v\n", item.Key, item.Value) // Key: key1, Value: 42
// 	}
// }
