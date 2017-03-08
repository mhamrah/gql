package query


const simpleQuery = `query FetchLukeQuery {
			human(id: "1000") {
				name
			}
		}`

const mutation = `mutation  {
						likeStory(storyID: 12345) {
							story {
								likeCount
							}
						}
					}`
