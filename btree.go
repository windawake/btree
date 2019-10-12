package main

import "fmt"
import "bufio"
import "os"
import "errors"
import "sort"
import "reflect"
import "strconv"

// var internal_node = []int{20, 50, 70, 100}
// var leaf_node = [][]int{{10,20},{33,40,50},{70,100}}

var internal_node = []int{}
var leaf_node = [][]int{}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
    exists = false
    index = -1

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }

    return
}

func getInternalPos(num int) (int, error) {
	last := 0
	for i,v := range internal_node {
		if num == v {
			return -1,errors.New("exist num in internal node")
		}

		if num < v {
			return i,nil
		}

		last = i + 1
	}

	return last, nil
}

func insertNode(num int) {
	pos,err := getInternalPos(num)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	leafArr := []int{}

	if len(leaf_node) > pos {
		leafArr = leaf_node[pos]
	}
	

	exists, _ := in_array(num, leafArr)
	if exists {
		fmt.Println("exist num in leaf node")
		os.Exit(1)
	}

	leafArr = append(leafArr, num)
	sort.Ints(leafArr)

	if len(leafArr) == 4 {
		leaf_left := leafArr[:2]
		leaf_right := leafArr[2:4]

		// 重新计算内部节点
		internalPretend := internal_node[:pos]
		internalNew := []int{}
		internalNew = append(internalNew,internalPretend...)
		internalNew = append(internalNew,leaf_left[1])
		if len(internal_node)>pos {
			internalAppend := internal_node[pos+1:]
			internalNew = append(internalNew,leaf_right[1])
			internalNew = append(internalNew,internalAppend...)
		}
		internal_node = internalNew

		// 重新计算叶子节点
		leafPretend := leaf_node[:pos]
		leafNew := [][]int{}
		leafNew = append(leafNew,leafPretend...)
		leafNew = append(leafNew,leaf_left)
		leafNew = append(leafNew,leaf_right)
		if len(leaf_node)>pos {
			leafAppend := leaf_node[pos+1:]
			leafNew = append(leafNew,leafAppend...)
		}
		leaf_node = leafNew
	}else{

		if len(leaf_node) > pos {
			leaf_node[pos] = leafArr
		}else{
			leaf_node = append(leaf_node, leafArr)
		}
	}
}

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输出primary key:")

	for scanner.Scan() {
		input := scanner.Text()

		if input == ".btree" {
			fmt.Println("btree内部节点:")
			fmt.Println(internal_node)

			fmt.Println("btree叶子节点:")
			fmt.Println(leaf_node)
			continue;
		}

		num, _ := strconv.Atoi(input)
		
		if num <= 0 {
			fmt.Println("必须输入大于0的整型数字")
			continue;
		}

		insertNode(num)
	}
}
