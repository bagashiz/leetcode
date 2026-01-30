package twopointers

import "strconv"

// https://leetcode.com/problems/palindrome-number/
func IsPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	num := strconv.Itoa(x)

	l, r := 0, len(num)-1

	for l < r {
		if num[l] != num[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func IsPalindromeNumber2(x int) bool {
	// Kasus khusus: negatif atau angka berakhiran 0 (kecuali 0)
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	// Kita terus memindahkan angka dari belakang x ke revertedNumber
	// Berhenti jika x sudah lebih kecil atau sama dengan revertedNumber
	// (Artinya kita sudah sampai di tengah-tengah angka)
	for x > revertedNumber {
		revertedNumber = (revertedNumber * 10) + (x % 10)
		x /= 10
	}

	// Jika jumlah digit genap: x == revertedNumber (misal 12 | 12)
	// Jika jumlah digit ganjil: x == revertedNumber / 10 (misal 12 | 121 -> 12 == 12)
	return x == revertedNumber || x == revertedNumber/10
}
