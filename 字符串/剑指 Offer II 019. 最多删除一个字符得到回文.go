package 字符串

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	isPalindrome := func(s string) bool {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
		}

		return true
	}
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left+1:right+1]) || isPalindrome(s[left:right])
		}

		left++
		right--
	}

	return true
}
