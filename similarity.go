package animagi

import (
	"errors"
	"fmt"
	"strings"
)

const (
	invalidString = "Invalid string"
	// additive constant for missing depth
	dFactor = 3
	// additive constant for missing letters, when len(str1) != len(str2)
	mlFactor = 5
	// letter change constant, str1[i] != str[i]
	lFactor = 1
	// MaxRank returned for invalid string or dissimalarity beyond size of uint
	MaxRank = ^uint(0)
)

/*
SimilarityRank computes the similarity between two strings
Some presumptions of the strings are to be considered:
 - a '.' denotes a depth increase
 - a string consistenting of only '.' will have maximum RANK
 - a letter is considered missing if one string is longer than the other
*/
func SimilarityRank(str1, str2 string) (rank uint) {

	if err := validateString(str1); err != nil {
		return MaxRank
	}

	if err := validateString(str2); err != nil {
		return MaxRank
	}

	str1Len := len(str1)
	str2Len := len(str2)
	shorterLen := str1Len

	str1Depths := strings.Split(str1, ".")
	str2Depths := strings.Split(str2, ".")

	fmt.Println("str1 split", len(str1Depths), "-", str1Depths)
	fmt.Println("str2 split", len(str2Depths), "-", str2Depths)

	if str1Len == 0 {
		return uint(mlFactor * str2Len)
	} else if str2Len == 0 {
		return uint(mlFactor * str1Len)
	}

	shorterLen = str1Len

	if str1Len < str2Len {
		rank = uint(mlFactor * (str2Len - str1Len))
	} else if str2Len < str1Len {
		shorterLen = str2Len
		rank = uint(mlFactor * (str1Len - str2Len))
	}

	for i := 0; i < shorterLen; i++ {
		if str1[i] != str2[i] {
			rank += lFactor
		}
	}
	return rank
}

func validateString(str string) (err error) {
	if str == "." {
		err = errors.New(invalidString)
	}
	strLen := len(str)
	if strLen > 0 {
		prevChar := str[0]
		for i := 1; i < len(str); i++ {
			if prevChar == '.' && prevChar == str[i] || str[i] == ' ' {
				return errors.New(invalidString)
			}
			prevChar = str[i]
		}
	}

	return err
}
