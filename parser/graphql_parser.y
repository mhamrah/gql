%{
package parser

import (
        "reflect"
        "github.com/mhamrah/gql/ast"
)
%}

%union {
    str string
    val reflect.Value
    doc *ast.Document

    definitions []*ast.Operation
    operationDefinition *ast.Operation
    opType ast.OpType
    selectionSet []*ast.Selection
    selection *ast.Selection
    field *ast.Field
    arguments []*ast.Argument
    argument *ast.Argument
}

%token <str> NAME
%token <val> VALUE
%token QUERY
%token MUTATION

%type <doc> document
%type <definitions> definition_list
%type <operationDefinition> operation_definition
%type <operationDefinition> definition
%type <opType> operation_type
%type <selectionSet> selection_set
%type <selectionSet> selection_set_opt
%type <selectionSet> selection_list
%type <selection> selection
%type <field> field
%type <arguments> arguments_opt
%type <arguments> arguments
%type <arguments> argument_list
%type <argument> argument
%type <str> name_opt

%%

document
        : definition_list
                {
                       $$ = &ast.Document{Definitions: $1}
                       yylex.(*lexer).doc = $$
                       return 0
                }
        ;

definition_list
    : /* nothing */ { $$ = nil }
    | definition_list definition { $$ = append($1, $2) }
    ;

definition
        : operation_definition { $$ = $1 }
        ;

operation_definition
        : operation_type name_opt selection_set
                {
                        $$ = &ast.Operation{OpType: $1, Name: $2, SelectionSet: $3}
                }
        ;

name_opt
        : /* nothing */ { $$ = "" }
        | NAME { $$ = $1 }
        ;

operation_type
        : QUERY { $$ = ast.Query }
        | MUTATION { $$ = ast.Mutation }
        ;

selection_set
        : '{' selection_list '}' { $$ = $2 }
        ;

selection_set_opt
         : /* nothing */ { $$ = nil }
         | selection_set
         ;

selection_list
        : /* nothing */ { $$ = nil }
        | selection_list selection { $$ = append($1, $2) }
        ;

selection
        : field { $$ = &ast.Selection{ Field: $1 } }
        ;

field
        : NAME arguments_opt selection_set_opt {
                $$ = &ast.Field{
                        Name: $1,
                        Arguments: $2,
                        SelectionSet: $3,
                        }
                }
        ;

arguments
        : '(' argument_list ')' { $$ = $2 }
        ;

arguments_opt
        : /* nothing */ { $$ = nil }
        | arguments { $$ = $1; }
        ;

argument_list
        : /* nothing */ { $$ = nil }
        | argument_list argument { $$ = append($1, $2) }
        ;

argument
        : NAME ':' VALUE { $$ = &ast.Argument{Name: $1, Value: $3} }
        ;

/* todo add test for query name is optional */
/* todo query shorthand */
/* todo fragmentSpread and inlineFragment for selection */
/* todo fix string regex and validate index bounds */
/* todo bail if "any" is reached */
/* todo subscriptions */