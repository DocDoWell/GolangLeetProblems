package countSeniors

import "strconv"

func countSeniors(details []string) int {
    count := 0;
    for _, s := range details{
        ageString := string(s[len(s)-4]) + string(s[len(s)-3])
        age, _ := strconv.Atoi(ageString)
        if(age > 60){
            count++
        }
    }
    return count
}