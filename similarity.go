package animagi

const (
	dFactor  = 3 // additive constant for missing depth
	mlFactor = 5 // additive constant for missing letters, when len(str1) != len(str2)
	lFactor  = 1 // letter change constant, str1[i] != str[i]
)

/*
SimilarityRank computes the similarity between two strings
Some presumptions of the strings are to be considered:
 a '.' denotes a depth increase
 a letter is considered missing if one string is longer than the other
*/
func SimilarityRank(str1, str2 string) (rank int) {
	str1Len := len(str1)
	str2Len := len(str2)
	shorterLen := str1Len

	if str1Len == 0 {
		return mlFactor * str2Len
	} else if str2Len == 0 {
		return mlFactor * str1Len
	}

	shorterLen = str1Len

	if str1Len < str2Len {
		rank = mlFactor * (str2Len - str1Len)
	} else if str2Len < str1Len {
		shorterLen = str2Len
		rank = mlFactor * (str1Len - str2Len)
	}

	for i := 0; i < shorterLen; i++ {
		if str1[i] != str2[i] {
			rank += lFactor
		}
	}
	return rank
}
