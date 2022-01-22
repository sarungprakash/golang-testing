package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/word-count", func(c *gin.Context) {
		var f GolangTestingForm
		if err := c.ShouldBind(&f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request",
			})
			return
		}
		s := []string{f.Description}
		r := getWordCount(s, 10)
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Successfully fetched words and its count",
			"words":   r,
		})
	})
	r.Run()
}

func getWordCount(s []string, c int) map[string]int {
	w := strings.Split(strings.Join(s, ""), " ")

	m := make(map[string]int)
	for _, word := range w {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	ct := make(map[string]int)
	for key, value := range m {
		if value > 0 {
			ct[key] = value
		}
	}

	keys := make([]string, 0, len(ct))
	for key := range ct {
		keys = append(keys, key)
	}
	fmt.Println("kkeys: ", keys)
	sort.Slice(keys, func(i, j int) bool {
		return ct[keys[i]] > ct[keys[j]]
	})
	fmt.Println("kkeys: ", keys)

	res := make(map[string]int)
	for _, key := range keys {
		res[key] = ct[key]
		c--
		if c == 0 {
			break
		}

	}
	return res
}

type GolangTestingForm struct {
	Description string `form:"description" binding:"required"`
}
