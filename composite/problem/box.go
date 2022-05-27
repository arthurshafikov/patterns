package main

type Box struct {
	Elements []any
}

func NewBox(elements []any) *Box {
	return &Box{
		Elements: elements,
	}
}

func (b *Box) OpenBox() []any {
	var result []any

	for _, elem := range b.Elements {
		smallerBox, ok := elem.(*Box)
		if !ok {
			result = append(result, elem)
		} else {
			result = append(result, smallerBox.OpenBox()...)
		}
	}

	return result
}
