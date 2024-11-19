package main

import (
	"fmt"
	"regexp"
	"strings"
)

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func editDistance(str1, str2 []string) (int, int, int, int) {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	i, j := m, n
	inserts, deletes, substitutions := 0, 0, 0
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			i, j = i-1, j-1
		} else {
			if dp[i][j] == dp[i-1][j-1]+1 {
				substitutions++
				i, j = i-1, j-1
			} else if dp[i][j] == dp[i-1][j]+1 {
				deletes++
				i--
			} else if dp[i][j] == dp[i][j-1]+1 {
				inserts++
				j--
			}
		}
	}

	inserts += j
	deletes += i

	return dp[m][n], inserts, deletes, substitutions
}

func addSpaces(text string) string {
	re := regexp.MustCompile(`([\p{Han}])`)
	text = re.ReplaceAllString(text, "$1 ")
	re = regexp.MustCompile(`([\p{Han}])([a-zA-Z0-9])`)
	text = re.ReplaceAllString(text, "$1 $2")
	re = regexp.MustCompile(`([a-zA-Z0-9])([\p{Han}])`)
	text = re.ReplaceAllString(text, "$1 $2")
	re = regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")
	return strings.TrimSpace(text)
}

func main3() {
	ref := "小思小思青年。"
	hyp := "不错"

	ref = strings.ToUpper(ref)
	hyp = strings.ToUpper(hyp)

	ref = addSpaces(ref)
	hyp = addSpaces(hyp)

	fmt.Println(ref)
	fmt.Println(hyp)

	refWords := strings.Split(ref, " ")
	hypWords := strings.Split(hyp, " ")

	distance, inserts, deletes, substitutions := editDistance(refWords, hypWords)

	fmt.Printf("Edit distance: %d\n", distance)
	fmt.Printf("Inserts: %d, Deletes: %d, substitutions: %d\n", inserts, deletes, substitutions)
}
