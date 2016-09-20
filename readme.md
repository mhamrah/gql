# gql: GraphQL support for YARPC-GO

This is an experimental POC to explore what
a GraphQL + YARPC combination would look like,
with only a small subset of functionality
implemented. Specifically query operations and
partial selection set support.
This is missing a lot of
major functionality, specifically type
inspection, which I believe can be handled
by `reflect` with auto-generation filling
in the missing pieces.

Packages

* ast
* * Types for representing a GraphQL query
* encoding
* * A YARPC encoding for parsing a GraphQL query,
invoking a GraphQL handler wrapping an implemented
service, and constructing a response.
* example
* * An example service using the library. The code
in `example/starwars/gql` is handwritten, but would
be autogen'd via thriftrw-plugin.
* parser
* * A YACC + Ragel parser/lexer for GraphQL queries,
using https://github.com/graphql/libgraphqlparser/blob/master/parser.ypp
as a foundation.
* transport
* * A GraphQLExtractor for YARPC's http transport. Requires
https://github.com/yarpc/yarpc-go/pull/338



