package speller

func SpellDigits(a int64) string {
	switch a {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	}
	return ""
}

func SpellDozens(a int64) string {
	if a < 10 {
		return SpellDigits(a)
	}
	switch a {
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 15:
		return "fifteen"
	}
	if a > 13 && a < 20 {
		return SpellDigits(a%10) + "teen"
	}
	ans := ""
	switch a / 10 {
	case 2:
		ans = "twenty"
	case 3:
		ans = "thirty"
	case 4:
		ans = "forty"
	case 5:
		ans = "fifty"
	case 6:
		ans = "sixty"
	case 7:
		ans = "seventy"
	case 8:
		ans = "eighty"
	case 9:
		ans = "ninety"
	}
	if a%10 > 0 {
		ans += "-" + SpellDigits(a%10)
	}
	return ans
}

func SpellHundreds(a int64) string {
	ans := ""
	if a == 0 {
		return SpellDozens(a % 100)
	}

	if a/100 > 0 {
		ans = SpellDigits(a/100) + " hundred"
		if a%100 > 0 {
			ans += " "
		}
	}

	if a%100 > 0 {
		ans += SpellDozens(a % 100)
	}
	return ans
}

func SpellParts(n int64, postfix string, f bool) string {
	ans := ""
	if f {
		ans += " "
	}
	ans += SpellHundreds(n) + postfix

	return ans
}

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}

	ans := ""
	if n < 0 {
		n = -n
		ans += "minus "
	}

	dozens := n % 1000
	thousands := n % 1000000 / 1000
	millions := n % 1000000000 / 1000000
	billions := n % 1000000000000 / 1000000000

	f := false
	if billions > 0 {
		ans += SpellParts(billions, " billion", f)
		f = true
	}

	if millions > 0 {
		ans += SpellParts(millions, " million", f)
		f = true
	}

	if thousands > 0 {
		ans += SpellParts(thousands, " thousand", f)
		f = true
	}

	if dozens > 0 {
		ans += SpellParts(dozens, "", f)
		f = true
	}

	return ans
}
