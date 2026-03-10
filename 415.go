func addStrings(num1, num2 string) string {
   var (
    result = make([]byte, max(len(num1), len(num2))+ 1)
    digit, carry byte
    idx = len(result)
   )

   for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0; i, j = i-1, j-1 {
        digit = carry
        if i >= 0 {
            digit += num1[i] - '0'
        }
        if j >= 0 {
            digit += num2[j] - '0'
        }

        carry = digit / 10
        result[idx-1] = '0' + (digit % 10)
        idx--
   }

   if carry == 1 {
    result[idx-1] = '1'
    idx--
   }

   return string(result[idx:])
}