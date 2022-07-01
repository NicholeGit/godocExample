package godocExample

// SomeTypeA 这一行，会被视为 SomeTypeA 的 GoDoc，
// 因为它紧挨着 SomeTypeA 的定义。
type SomeTypeA struct{}

// 这一行与 SomeTypeB 的定义之间隔了一行，
// 所以并不会认为是 SomeTypeB 的 GoDoc。

type SomeTypeB struct{}

/*
SomeTypeC 使用这种注释符的注释也是同理，因为整个注释块紧挨着 SomeTypeC 的定义，
因此会被视为 SomeTypeC 的注释。
*/
type SomeTypeC struct{}

// SomeNewLine 只是用来展示如何在 GoDoc 中换行。
//
// 你看，这就是新的一行了，耶～✌️
func SomeNewLine() error {
	return nil
}

// IntsElem 在注释块中，如果部分注释行符合以下标准之一，则视为代码块:
//
// - 注释行以制表符 \t 开头
//
// - 注释行以以多于一个空格（包括制表符）开头
//
// 普通注释和代码块之间可以不用专门的空注释行，但个人建议还是加上比较好。
// 代码块的例子:
//
//    sli := []int{0, -1, -2, -3}
//    val, idx := IntsElem(sli, -2)
//
// 返回得 val = -2, actualIndex = 2
func IntsElem(ints []int, index int) (value, actualIndex int) {
	// ......
	return
}

// Deprecated: ElemAt 这个函数弃用，后续请迁移到 IntsElem 函数中.
func ElemAt(ints []int, index int) int {
	// ......
	return 0
}

type Set struct {
}

func (s *Set) At() {

}

type Opt struct {
}
