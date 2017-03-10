package query

const failure = `
  a{s}s... sss 123 aaaa aaa
`

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

const kitchenSink = `query queryName($foo: ComplexType, $site: Site = MOBILE) {
  whoever123is: node(id: [123]) {
    id ,
    ... on User @defer {
      field2 {
        id ,
        alias: field1(first:10, after:$foo,) @include(if: $foo) {
          id,
          ...frag
        }
      }
    }
    ... @skip(unless: $foo) {
      id
    }
    ... {
      id
    }
  }
}

mutation likeStory {
  like(story: 123) @defer {
    story {
      id
    }
  }
}

subscription StoryLikeSubscription($input: StoryLikeSubscribeInput) {
  storyLikeSubscribe(input: $input) {
    story {
      likers {
        count
      }
      likeSentence {
        text
      }
    }
  }
}

fragment frag on Friend {
  foo(size: $size, bar: $b, obj: {key: "value"})
}

{
  unnamed(truthy: true, falsey: false, nullish: null),
  query
}`
