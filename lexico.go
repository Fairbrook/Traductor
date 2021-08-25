package main

func processString(str string) (res []Segment, err error) {
	index := 0
	var segment Segment
	input := str + "$"
	for index < len(input) {
		segment, err = evaluate(input[index:])
		index += segment.Index
		if err != nil {
			return
		}
		res = append(res, segment)
	}
	return
}
