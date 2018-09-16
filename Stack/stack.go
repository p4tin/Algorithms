package Stack

type Stack struct {
	Data []interface{}
}

func CreateStack() *Stack {
	stack := &Stack{
		Data: make([]interface{}, 0),
	}
	return stack
}

func (stack *Stack) Push(i interface{}) {
	stack.Data = append(stack.Data, i)
}

func (stack *Stack) Pop() interface{} {
	if len(stack.Data) == 0 {
		return nil
	}
	retval := stack.Data[len(stack.Data)-1]
	stack.Data = append(stack.Data[:len(stack.Data)-1])
	return retval
}

func (stack *Stack) IsEmpty() bool {
	if len(stack.Data) == 0 {
		return true
	}
	return false
}
