package mapslice_test

import (
	"fmt"
	"github.com/pikanezi/mapslice"
)

type Tag struct {
	Name string
}

type Post struct {
	Title string
	Tags  []*Tag
}

func Example() {
	post := &Post{
		Title: "GOLANG",
		Tags: []*Tag{
			{"Go"}, {"Golang"}, {"Gopher"},
		},
	}
	s, _ := mapslice.MapSliceToString(post.Tags, "Name")
	fmt.Println(s)

	// Output:
	// [Go Golang Gopher]
}
