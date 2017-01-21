package query

import (
	"fmt"
	"testing"
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

	fmt.Println(ParseString(q))
}
