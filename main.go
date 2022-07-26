package main

import (
	"fmt"
	// "main/datastructures/array"
	// "main/datastructures/hash_table"
	"main/datastructures/linked_list"
)

func main() {
	// ========== Array ===========
	// arr := array.New()
	//  arr.Push("a")
	//  arr.Push("b")
	//  arr.Push("c")
	//  arr.Delete(1)
	//  fmt.Printf("%+v\n", arr)

	//  fmt.Println(array.Reverse("abcdef"))

	//  arr1 := []int{0, 3, 4, 31}
	//  arr2 := []int{4, 6, 30}
	//  fmt.Println(array.MergeSorted(arr1, arr2))

	// ============== Hash Table ==============
	// hashTable := hash_table.New(50)
	// hashTable.Set("Mandag", 1)
	// hashTable.Set("Mandag", 3)
	// hashTable.Set("Tirsdag", 88)
	// _, err := hashTable.Get("B")
	// if err != nil {
	//   fmt.Println(err)
	// }
	// fmt.Println(hashTable)
	// fmt.Println(hashTable.Keys())
	// a1 := []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	// a2 := []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	// a3 := []int{2, 3, 4, 5}
	// fmt.Println(hash_table.FirstRecurringCharacter(a1))
	// fmt.Println(hash_table.FirstRecurringCharacter(a2))
	// fmt.Println(hash_table.FirstRecurringCharacter(a3))

	// =============== Linked list ==============
	linkedList := linkedlist.New()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Prepend(3)
	linkedList.Append(4)
	linkedList.Insert(3, 45)
	linkedList.Insert(1, 45)
	linkedList.Print()
	linkedList.Remove(2)
	linkedList.Print()
	err := linkedList.Remove(33)
	if err != nil {
		fmt.Println(err)
	}
}
