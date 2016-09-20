package parser

import (
	"testing"

	"github.com/kr/pretty"
)

const q = `query FetchLukeQuery {
			human(id: "1000") {
				name
			}
		}`

const m = `mutation  {
						likeStory(storyID: 12345) {
							story {
								likeCount
							}
						}
					}`

func TestGraphQLParse(t *testing.T) {
	parse(q)
	pretty.Println(parse(m))
}
