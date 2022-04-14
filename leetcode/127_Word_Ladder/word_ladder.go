package main

func stringsDiff(a, b string) int {
	diffLetterCount := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffLetterCount++
		}
	}

	return diffLetterCount
}

// Problem https://leetcode.com/problems/word-ladder/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	q := make([]string, 1)
	d := make(map[string]int)
	m := make(map[string]struct{})

	q[0] = beginWord
	d[beginWord] = 1
	m[beginWord] = struct{}{}

	for len(q) > 0 {
		w := q[0]
		q = q[1:]

		for i := 0; i < len(wordList); i++ {
			diff := stringsDiff(w, wordList[i])

			_, marked := m[wordList[i]]
			if diff < 2 && !marked {
				d[wordList[i]] = d[w] + 1
				m[wordList[i]] = struct{}{}
				q = append(q, wordList[i])
			}
		}
	}

	minPathLen := d[endWord]

	return minPathLen
}
