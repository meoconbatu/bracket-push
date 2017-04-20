package brackets

const testVersion = 4

type RuneHeap []rune

func (h RuneHeap) Len() int { return len(h) }

func (h *RuneHeap) Push(x interface{}) {
	*h = append(*h, x.(rune))
}

func (h *RuneHeap) Pop() interface{} {
	old := *h
	n := len(old)
	if n == 0 {
		return nil
	}
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var bracket = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

func Bracket(input string) (output bool, err error) {
	h := &RuneHeap{}
	for _, a := range input {
		if a == '{' || a == '[' || a == '(' {
			h.Push(a)
		} else {
			temp := h.Pop()
			if bracket[a] != temp {
				return false, nil
			}
		}
	}
	if h.Len() > 0 {
		return false, nil
	}
	return true, nil
}
