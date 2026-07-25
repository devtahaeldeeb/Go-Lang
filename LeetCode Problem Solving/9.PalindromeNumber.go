func isPalindrome(x int) bool {
    if x < 0{
        return false
    }
    original := x
    reserved := 0

    for x > 0{ 
        digit := x % 10
        reserved = reserved * 10 + digit
        x /= 10 
    }
    return original == reserved
}
