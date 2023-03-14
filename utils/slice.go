package utils

import (
	"hackz-allo/database"
	"math"
)

func RemoveFromSlice(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

func SortPost(posts []database.Post, limit int) []database.Post {
	l := len(posts)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if posts[i].Time < posts[j].Time {
				posts[i], posts[j] = posts[j], posts[i]
			}
		}
	}
	return posts[0:int(math.Min(float64(l), float64(limit)))]
}
