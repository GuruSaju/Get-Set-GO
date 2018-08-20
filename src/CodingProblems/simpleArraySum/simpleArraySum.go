/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
    var sum int32 = 0
    for _,value := range ar {
        fmt.Println(value)
        sum =sum + value
    }
    return sum
}