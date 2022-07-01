package godocExample

import "fmt"

// 需要注意的是，示例代码文件也属于单元测试文件的内容，当执行 go test 的时候，示例文件也会纳入测试逻辑中。

//
//
// Example 这是示例代码的固有开头
//
// Set 表示这是类型 Set 的示例
//
// 第一个下划线 _  分隔符，在这个分隔符后面的，是 Set 类型的成员函数名
//
// At 表示这是函数 At() 的示例，搭配前面的内容，则表示这是类型 Set 的成员函数 At() 的示例
//
// 第二个下划线 _ 分隔符，在这个分隔符后面的内容，是示例代码的额外说明
//
// 1  这是示例代码的额外说明，也就是前面 “Example (1)” 括号里的部分

// ExampleSet_At_1  Set.At() 的示例
func ExampleSet_At_1() {
	fmt.Println("ExampleSet_At_1")
	//......
}

// ExampleOpt 相对应地，如果你想要给（不属于任何一个类型的）函数写示例的话，则去掉上文中关于 “类型” 的字段；
//
// 如果你不需要示例的额外说明符，则去掉 “额外说明” 字段。
func ExampleOpt() {
	fmt.Println("ExampleOpt")
	// Output: 单独起一行开头，剩下的每一行标准输出写一行注释。
	// {"null":null}
	// {}
}
