package main

import "fmt"

func main() {
myMap := make(map[string]int)

var numPairs int
fmt.Print("Enter the number of key-value pairs: ")
fmt.Scan(&numPairs)

for i := 0; i < numPairs; i++ {
var key string
var value int
fmt.Printf("Enter key %d: ", i+1)
fmt.Scan(&key)
fmt.Printf("Enter value %d: ", i+1)
fmt.Scan(&value)
myMap[key] = value
}

var key string
fmt.Print("Enter a key to retrieve its value: ")
fmt.Scan(&key)

value, exists := myMap[key]
if exists {
fmt.Printf("Value of %s: %d\n", key, value)
} else {
fmt.Printf("%s not found in the map\n", key)
}

var updatedKey string
var updatedValue int
fmt.Print("Enter a key to update its value: ")
fmt.Scan(&updatedKey)
fmt.Print("Enter the updated value: ")
fmt.Scan(&updatedValue)
myMap[updatedKey] = updatedValue
fmt.Println("Updated value of ", updatedKey, "is ", updatedValue)

var deleteKey string
fmt.Print("Enter a key to delete from the map: ")
fmt.Scan(&deleteKey)

delete(myMap, deleteKey)
fmt.Println("After deleting", deleteKey, ":", myMap)

for key, value := range myMap {
fmt.Println("Key:", key, "Value:", value)
}

fmt.Println("Length of the map:", len(myMap))
}


// Output:
// Enter the number of key-value pairs: 3
// Enter key 1: apple
// Enter value 1: 1
// Enter key 2: banana
// Enter value 2: 2
// Enter key 3: orange
// Enter value 3: 3
// Enter a key to retrieve its value: banana
// Value of banana: 2
// Enter a key to update its value: apple
// Enter the updated value: 5
// Updated value of  apple is  5
// Enter a key to delete from the map: orange
// After deleting orange : map[apple:5 banana:2]
// Key: banana Value: 2
// Key: apple Value: 5
// Length of the map: 2
