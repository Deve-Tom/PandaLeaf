package utility

import "fmt"

func ExamplePair() {
	// 创建一个Pair实例
	pair := Pair[int, string]{First: 1, Second: "Hello World!"}

	// 输出pair的内容
	fmt.Println(pair)

	// 输出pair的类型
	fmt.Printf("%T\n", pair)

	// 输出pair的第一个元素
	fmt.Println(pair.First)

	// 输出pair的第二个元素
	fmt.Println(pair.Second)

	// Output:
	// {1 Hello World!}
	// utility.Pair[int,string]
	// 1
	// Hello World!
}

func ExampleTuple_InitTuple() {
	// 创建一个Tuple实例
	tuple := Tuple{}
	tuple.InitTuple(1, "Hello World!", 3.14)

	// 输出tuple的内容
	fmt.Println(tuple)

	// 输出tuple的类型
	fmt.Printf("%T\n", tuple)

	// 输出tuple的第一个元素
	fmt.Println(tuple.Items[0])

	// 输出tuple的第二个元素
	fmt.Println(tuple.Items[1])

	// 输出tuple的第三个元素
	fmt.Println(tuple.Items[2])

	// Output:
	// {[1 Hello World! 3.14]}
	// utility.Tuple
	// 1
	// Hello World!
	// 3.14
}

func ExampleTuple_Compare() {
	// 创建两个Tuple实例-tuple1和tuple2类型和值都相同
	tuple1 := Tuple{}
	tuple1.InitTuple(1, "Hello World!", 3.14)
	tuple2 := Tuple{}
	tuple2.InitTuple(1, "Hello World!", 3.14)

	// 比较两个Tuple实例-tuple1和tuple2
	result, message := tuple1.Compare(tuple2)
	fmt.Println(result, message)

	// 创建两个Tuple实例-tuple3和tuple4类型相同但值不同
	tuple3 := Tuple{}
	tuple3.InitTuple(1, "Hello World!", 3.14)
	tuple4 := Tuple{}
	tuple4.InitTuple(1, "Hello World!", 3.1415926)

	// 比较两个Tuple实例-tuple3和tuple4
	result, message = tuple3.Compare(tuple4)
	fmt.Println(result, message)

	// 创建两个Tuple实例-tuple5和tuple6类型不同，值也不同
	tuple5 := Tuple{}
	tuple5.InitTuple(1, "Hello World!", 3)
	tuple6 := Tuple{}
	tuple6.InitTuple(1, "Hello World!", 3.14)

	// 比较两个Tuple实例-tuple5和tuple6
	result, message = tuple5.Compare(tuple6)
	fmt.Println(result, message)

	// 创建两个Tuple实例-tuple7和tuple8类型相同但数量不同
	tuple7 := Tuple{}
	tuple7.InitTuple(1, "Hello World!", 3.14)
	tuple8 := Tuple{}
	tuple8.InitTuple(1, "Hello World!")

	// 比较两个Tuple实例-tuple7和tuple8
	result, message = tuple7.Compare(tuple8)
	fmt.Println(result, message)

	// Output:
	// true Two tuples are equal.
	// false Value of two tuples are not equal.
	// false Type of two tuples are not equal.
	// false Length of two tuples are not equal.
}
