%{
package query

import (
        "reflect"
)
%}

%union {
    str string
    val reflect.Value
    vals []reflect.Value
    doc Document

    definitions []Definition
    definition Definition

    directives []Directive
    directive Directive

    arguments []Argument
    argument Argument

    variables []Variable
    variable Variable

    selectionSet []Selection
    selection Selection

    operation Operation
    opType OpType
    field Field
    fragment Fragment
    fragmentSpread FragmentSpread
}

%token <str> NAME
%token <val> VALUE
%token QUERY MUTATION SUBSCRIPTION
%token FRAGMENT SPREAD ON
%token <str> VARIABLE
%token <str> FRAGMENT_NAME

%type <doc> document
%type <definitions> definition_list
%type <operation> operation_definition
%type <definition> definition
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
%type <fragment> fragment_definition
%type <directive> directive
%type <variables> variable_definitions
%type <variables> variable_definition_list
%type <variable> variable_definition
%type <str> type
%type <val> default_value_opt
%type <val> default_value
%type <val> value_const
%type <vals> value_list
%type <vals> list_value
%type <fragmentSpread> fragment_spread
%type <fragment> inline_fragment
%type <directives> directives_opt
%type <directives> directives
%type <str> type_condition


%%

/* 2.2 Document */

document
        : definition_list
                {
                       $$ = Document{Definitions: $1}
                       yylex.(*lexer).doc = $$
                       return 0
                }
        ;

name_opt
        : /* nothing */ { $$ = "" }
        | NAME { $$ = $1 }
        ;

definition_list
        : /* nothing */ { $$ = nil }
        | definition_list definition { $$ = append($1, $2) }
        ;

definition
        : operation_definition { $$ = Definition{ operation: $1 } }
        | fragment_definition { $$ = Definition{ fragment: $1 } }
        ;

/* 2.2.1 Operations */
operation_definition
        : selection_set
                {
                        $$ = Operation{OpType: Query, SelectionSet: $1 }
                }
        | operation_type name_opt selection_set
                {
                        $$ = Operation{OpType: $1, Name: $2, SelectionSet: $3}
                }
        | operation_type name_opt variable_definitions selection_set
                {
                        $$ = Operation{OpType: $1, Name: $2, Variables: $3, SelectionSet: $4 }
                }
        | operation_type name_opt directives selection_set
                {
                        $$ = Operation{OpType: $1, Name: $2, Directives: $3, SelectionSet: $4 }
                }
        | operation_type name_opt variable_definitions directives selection_set
                {
                        $$ = Operation{OpType: $1, Name: $2, Variables: $3, Directives: $4, SelectionSet: $5 }
                }
        ;

operation_type
        : QUERY { $$ = Query }
        | MUTATION { $$ = Mutation }
        | SUBSCRIPTION { $$ = Subscription }
        ;

variable_definitions
        : '(' variable_definition_list ')' { $$ = $2 }
        ;

variable_definition_list
        : /* nothing */ { $$ = nil }
        | variable_definition_list variable_definition { $$ = append($1, $2) }
        ;

variable_definition
        : VARIABLE ':' type default_value_opt
                {
                        $$ = Variable{Name: $1, Type: $3, Default: $4}
                }
        ;

default_value_opt
        : /* nothing */ { $$ = reflect.ValueOf(nil) }
        | default_value { $$ = $1 }
        ;

default_value
        : '=' value_const { $$ = $2; }

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
        : field { $$ = Selection{ Field: $1 } }
        | fragment_spread { $$ = Selection{ FragmentSpread: $1 } }
        | inline_fragment { $$ = Selection{ InlineFragment: $1 } }
        ;

field
        : NAME arguments_opt directives_opt selection_set_opt {
                $$ = Field{
                        Name: $1,
                        Arguments: $2,
                        Directives: $3,
                        SelectionSet: $4,
                        }
                }
        | NAME ':' NAME arguments_opt directives_opt selection_set_opt {
                $$ = Field{
                        Name: $1,
                        Arguments: $4,
                        Directives: $5,
                        SelectionSet: $6,
                        Alias: $3,
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
        : NAME ':' value_const { $$ = Argument{Name: $1, Value: $3} }
        ;

/* 2.2.6 Fragments */
fragment_spread
        : SPREAD NAME directives_opt { $$ = FragmentSpread{ Name: $2, Directives: $3 } }
        ;

inline_fragment
        : SPREAD ON type_condition directives_opt selection_set
                {
                        $$ = Fragment{
                                TypeCondition: $3,
                                Directives: $4,
                                SelectionSet: $5,
                        }
                }
        | SPREAD directives_opt selection_set
                {
                        $$ = Fragment{
                                Directives: $2,
                                SelectionSet: $3,
                        }
                }
        ;

fragment_definition
        : FRAGMENT FRAGMENT_NAME ON type_condition directives_opt selection_set
                {
                        $$ = Fragment{
                                        Name: $2,
                                        TypeCondition: $4,
                                        Directives: $5,
                                        SelectionSet: $6,
                                };
                }
        ;

type_condition
        : NAME { $$ = $1 }
        ;

/* 2.2.7 Input Values */

value_const
        : VALUE { $$ = $1 }
        | NAME { $$ = reflect.ValueOf($1) }
        | list_value { $$ = reflect.ValueOf($1) }
        | VARIABLE { $$ = reflect.ValueOf($1) }
        ;

/* 2.2.7.6 List Value */

list_value
        : '[' value_list ']' { $$ = $2 }

value_list
        : /* nothing */ { $$ = nil }
        | value_list value_const { $$ = append($1, $2) }
        ;

/* 2.2.9 Types */

type
        : NAME { $$ = $1 }
        ;

/* 2.2.10 Directives */

directives_opt
        : /* nothing */ { $$ = nil }
        | directives_opt directive { $$ = append($1, $2) }
        ;

directives
        : directives_opt directive { $$ = append($1, $2) }
        ;

directive
        : '@' NAME arguments_opt { $$ = Directive{ Name: $2, Arguments: $3}; }
        ;

/* todo subscriptions */