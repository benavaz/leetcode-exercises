func plusOne(digits []int) []int {
    digit := 0
    carriage := 0
    size := len(digits)
    sumar := 1
    for i := 1; i <= size ; i++ {
        digit = digits[size - i]
        digit = digit + sumar + carriage
        if digit < 10{
            digits[size - i] = digit
            return digits
        } 

        digits[size - i] = digit - 10
        carriage = 1
        
        if size == i {
            digits = append([]int{1}, digits...)
        }
        sumar = 0
    }
    return digits
}