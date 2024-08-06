package minimumpushes

import (
	"fmt"
	"sort"
)


type letter struct {
    label string
    frequency  int
}

func (l *letter) incrementFrequency() {
  l.frequency++
}

func minimumPushes(word string) int {

    letters:=make(map[string]letter)
	amp:=0;
	count:= 0
    output:= 0
    
    for i:=0; i<len(word); i++{
        label:=string(word[i])

        _, ok := letters[label]

        if !ok {
            letters[label] = letter{label: label, frequency: 1}
        }else{
            l2:= letters[label]
			l2.incrementFrequency()
			letters[label] = l2
        }
    }

	values := make([]letter, 0, len(letters))
    for k := range letters {
        values = append(values, letters[k])
    }

	sort.Slice(values[:], func(i, j int) bool {
		return values[i].frequency > values[j].frequency
	})

	fmt.Println(values)

	for _, let := range values{
		button:= (count%8)+1
		fmt.Println("B: ",button)
		if button == 1{
			amp++
			fmt.Println("A: ",amp)
		}

		output = output + (amp*let.frequency)
		fmt.Println("O : ",output)
		count++
	}

	return output
}