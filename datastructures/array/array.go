package array

var Strings []string

func init() {
    Strings = []string{"a", "b", "c", "d"}

    // push
    Strings = append(Strings, "e")

    // pop
    Strings = Strings[:len(Strings)-1]

    // add first elem
    Strings = append([]string{"x"}, Strings...)

    // add to the middle
    Strings = append(Strings[:3], Strings[2:]...)
    Strings[2] = "alien"
}

func New() Array {
  return Array{
    Data: make(map[int]string),
  }
}

type Array struct {
  Len int
  Data map[int]string
}

func (a Array) Get(index int) string {
  return a.Data[index] 
}

func (a *Array) Push(item string) {
  a.Data[a.Len] = item
  a.Len += 1
}

func (a *Array) Pop() string {
  a.Len -= 1
  val := a.Data[a.Len]
  delete(a.Data, a.Len)
  return val
}

func (a *Array) Delete(index int) string {
  deleteItem := a.Data[index]
  for i := index; i < a.Len-1; i++ {
    a.Data[i] = a.Data[i+1]
  }
  delete(a.Data, a.Len-1)
  a.Len -= 1
  return deleteItem
}

func Reverse(s string) string {
  if len(s) == 1 {
    return s
  }
  return string(s[len(s)-1]) + Reverse(string(s[:len(s)-1]))
}

func MergeSorted(arr1 []int, arr2 []int) []int {
  merged := []int{}
  idx1 := 0
  idx2 := 0

  for len(merged) != len(arr1) + len(arr2) {
    if idx1 >= len(arr1) {
      merged = append(merged, arr2[idx2])
      idx2 += 1
      continue
    }

    if idx2 >= len(arr2) {
      merged = append(merged, arr1[idx1])
      idx1 += 1
      continue
    }

    if arr1[idx1] <= arr2[idx2] {
      merged = append(merged, arr1[idx1])
      idx1 += 1
    } else {
      merged = append(merged, arr2[idx2])
      idx2 += 1
    }
  }
  return merged
}
