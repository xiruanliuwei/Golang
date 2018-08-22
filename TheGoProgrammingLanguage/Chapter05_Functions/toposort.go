
package main

import "fmt"
import "sort"

// 注意 {} 在Go语言中的特殊作用，{} 表示 literal 的意思
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string    // a slice

	seen := make(map[string]bool)    // a map, key type is string, value type is bool
	var visitAll func(items []string)    // 函数变量  function variable

	visitAll = func(items []string) {
		for _, item := range items {
			fmt.Println(item)
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])  // 注意形成的嵌套
				order = append(order, item)
				fmt.Println(order)
			}
		}
	}

	var keys []string
	for key := range m{
		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Println(keys)

	visitAll(keys)
	
	return order
}

func main() {

	var bv bool
	fmt.Println(bv)

	fmt.Printf("type of prereqs is: %T\n", prereqs)  // map[string][]string

	test_var1 := prereqs["algorithms"]
	fmt.Printf("type of test_var1 is: %T\n", test_var1)  // []string
	fmt.Println(len(test_var1))  // 1
	fmt.Println(test_var1)

	test_var2 := prereqs["compilers"]
	fmt.Printf("type of test_var2 is: %T\n", test_var2)  // []string
	fmt.Println(len(test_var2))  // 3
	fmt.Println(test_var2)

	var x = []string{
		"aaa", 
		"bbb", 
		"ccc",
	}
	fmt.Printf("type of x is: %T\n", x)
	fmt.Println(x)

	x1_2 := x[1:2]
	fmt.Println(x1_2)

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s \n", i + 1, course)
	}
}