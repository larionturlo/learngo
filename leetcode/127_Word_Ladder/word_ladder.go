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

type layer map[string]string

// Problem https://leetcode.com/problems/word-ladder/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	layerList := make([]layer, len(wordList), len(wordList))
	minPathLen := 0
	temp := make(map[string]int)
	for i := range wordList {
		if wordList[i] == endWord {
			wordList[i], wordList[len(wordList)-1] = wordList[len(wordList)-1], wordList[i]
		} else if wordList[i] == beginWord {
			wordList[i], wordList[0] = wordList[0], wordList[i]
		}
	}

	for i := 0; i < len(wordList); i++ {
		layer := stringsDiff(beginWord, wordList[i])
		beforLayerIndex := layer - 1
		beforWord := beginWord

		if beforLayerIndex > 1 {
			for word := range layerList[beforLayerIndex] {
				if stringsDiff(word, wordList[i]) == 1 {
					beforWord = word
					break
				}
			}
			if beforWord == beginWord {
				temp[wordList[i]] = layer
			}
		}

		if endWord == wordList[i] {
			minPathLen = layer
		}

		if len(layerList[layer]) == 0 {
			layerList[layer] = make(map[string]string)
		}

		layerList[layer][wordList[i]] = beforWord
	}

	return minPathLen
}
