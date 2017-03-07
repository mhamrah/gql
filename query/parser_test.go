package query

import "testing"
import "github.com/stretchr/testify/assert"

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
	doc, err := ParseString(q)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
}
