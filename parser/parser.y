%{
package parser

import (
        "reflect"
        "github.com/mhamrah/gql"
)
%}

%union {
    str string
    strs []string
    val reflect.Value
    vals []reflect.Value
    doc gql.Document

    definitions []gql.Definition
    definition gql.Definition

    directives []gql.Directive
    directive gql.Directive

    arguments map[string]gql.Argument
    argument gql.Argument

    variables []gql.Variable
    variable gql.Variable

    selectionSet []gql.Selection
    selection gql.Selection

    operation gql.Operation
    opType gql.OpType
    field gql.Field
    fragment gql.Fragment
    fragmentSpread gql.FragmentSpread
    objectField gql.ObjectField
    objectFields []gql.ObjectField

    schema gql.Schema
    operationTypeDefinition gql.OperationTypeDefinition
    operationTypeDefinitions []gql.OperationTypeDefinition
    typeDefinition gql.TypeDefinition
    typeDefinitions map[string]gql.TypeDefinition
    inputDefinitions map[string]gql.InputValueDefinition
    inputDefinition gql.InputValueDefinition
    fieldDef gql.FieldDefinition
    fieldsDef []gql.FieldDefinition
    typeDesc gql.TypeDescription
}

%token <str> DIRECTIVE
%token <str> ENUM
%token <str> EXTEND
%token <str> FALSE
%token <str> FRAGMENT
%token <str> IMPLEMENTS
%token <str> INPUT
%token <str> INTERFACE
%token <str> MUTATION
%token <str> NULL
%token <str> QUERY
%token <str> ON
%token <str> SCALAR
%token <str> SCHEMA
%token <str> SUBSCRIPTION
%token <str> TRUE
%token <str> TYPE
%token <str> UNION

%token <str> VARIABLE
%token <str> IDENTIFIER
%token <val> VALUE

%token SPREAD ON

%token <str> FRAGMENT_NAME

%type <doc> start
%type <doc> document
%type <str> fragment_name
%type <str> name
%type <str> name_opt

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
%type <fragment> fragment_definition
%type <directive> directive
%type <variables> variable_definitions
%type <variables> variable_definition_list
%type <variable> variable_definition
%type <typeDesc> type
%type <str> type_name
%type <typeDesc> list_type
%type <typeDesc> non_null_type
%type <str> enum_value
%type <val> default_value_opt
%type <val> default_value
%type <val> value_const
%type <val> boolean_value
%type <vals> value_list
%type <vals> list_value
%type <fragmentSpread> fragment_spread
%type <fragment> inline_fragment
%type <directives> directives_opt
%type <directives> directives
%type <str> type_condition
%type <objectField> object_field
%type <objectFields> object_field_list
%type <objectFields> object_value

%type <schema> schema
%type <operationTypeDefinition> operation_type_definition
%type <operationTypeDefinitions>
operation_type_definition_list
%type <operationTypeDefinitions> schema_definition
%type <typeDefinition> scalar_type_definition
%type <typeDefinitions> type_definitions
%type <typeDefinition> object_type_definition
%type <strs> implements_interfaces_opt
%type <strs> type_name_list
%type <fieldDef> field_definition
%type <fieldsDef> field_definition_list
%type <inputDefinitions> arguments_definition_opt
%type <inputDefinitions> arguments_definition
%type <inputDefinitions> input_value_definition_list
%type <inputDefinition> input_value_definition
%type <typeDefinition> interface_type_definition
%type <typeDefinition> union_type_definition
%type <strs> union_members
%type <typeDefinition> enum_type_definition
%type <str> enum_value_definition
%type <strs> enum_value_definition_list
%type <typeDefinition> input_object_type_definition
%%

start:  document
                {
                        // TODO check error
                        yylex.(*lexer).doc = $1
                        return 0
                }
        ;

// Fragment names are all names (identifiers) except ON
fragment_name:  DIRECTIVE       { $$ = $1 }
        |       ENUM            { $$ = $1 }
        |       EXTEND          { $$ = $1 }
        |       FALSE           { $$ = $1 }
        |       FRAGMENT        { $$ = $1 }
        |       IDENTIFIER      { $$ = $1 }
        |       IMPLEMENTS      { $$ = $1 }
        |       INPUT           { $$ = $1 }
        |       INTERFACE       { $$ = $1 }
        |       MUTATION        { $$ = $1 }
        |       NULL            { $$ = $1 }
        |       QUERY           { $$ = $1 }
        |       SCALAR          { $$ = $1 }
        |       SCHEMA          { $$ = $1 }
        |       SUBSCRIPTION    { $$ = $1 }
        |       TRUE            { $$ = $1 }
        |       TYPE            { $$ = $1 }
        |       UNION           { $$ = $1 }
        ;

name: fragment_name { $$ = $1 }
        |       ON { $$ = $1 }
        ;

name_opt: /* nothing */ { $$ = "" }
        | name
        ;

/* 2.2 Document */

document
        : definition_list schema
                {
                       $$ = gql.Document{Definitions: $1, Schema: $2 }
                }
        ;


definition_list
        : /* nothing */ { $$ = nil }
        | definition_list definition { $$ = append($1, $2) }
        ;

definition
        : operation_definition { $$ = gql.Definition{ Operation: $1 } }
        | fragment_definition { $$ = gql.Definition{ Fragment: $1 } }
        ;

/* 2.2.1 Operations */
operation_definition
        : selection_set
                {
                        $$ = gql.Operation{OpType: gql.Query, SelectionSet: $1 }
                }
        | operation_type name_opt selection_set
                {
                        $$ = gql.Operation{OpType: $1, Name: $2, SelectionSet: $3}
                }
        | operation_type name_opt variable_definitions selection_set
                {
                        $$ = gql.Operation{OpType: $1, Name: $2, Variables: $3, SelectionSet: $4 }
                }
        | operation_type name_opt directives selection_set
                {
                        $$ = gql.Operation{OpType: $1, Name: $2, Directives: $3, SelectionSet: $4 }
                }
        | operation_type name_opt variable_definitions directives selection_set
                {
                        $$ = gql.Operation{OpType: $1, Name: $2, Variables: $3, Directives: $4, SelectionSet: $5 }
                }
        ;

operation_type
        : QUERY { $$ = gql.Query }
        | MUTATION { $$ = gql.Mutation }
        | SUBSCRIPTION { $$ = gql.Subscription }
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
                        $$ = gql.Variable{Name: $1, Type: $3, Default: $4}
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
        : field { $$ = gql.Selection{ Field: $1 } }
        | fragment_spread { $$ = gql.Selection{ FragmentSpread: $1 } }
        | inline_fragment { $$ = gql.Selection{ InlineFragment: $1 } }
        ;

field
        : name arguments_opt directives_opt selection_set_opt {
                $$ = gql.Field{
                        Name: $1,
                        Arguments: $2,
                        Directives: $3,
                        SelectionSet: $4,
                        }
                }
        | name ':' name arguments_opt directives_opt selection_set_opt {
                $$ = gql.Field{
                        Name: $1,
                        Arguments: $4,
                        Directives: $5,
                        SelectionSet: $6,
                        Alias: $3,
                        }
                }
        ;

arguments: '(' argument_list ')' { $$ = $2 }
        ;

arguments_opt: /* nothing */ { $$ = nil }
        | arguments { $$ = $1; }
        ;

argument_list: /* nothing */ { $$ = make(map[string]gql.Argument) }
        | argument_list argument { $1[$2.Name] = $2; $$ = $1; }
        ;

argument: name ':' value_const { $$ = gql.Argument{Name: $1, Value: $3} }
        ;

/* 2.2.6 Fragments */
fragment_spread
        : SPREAD fragment_name directives_opt { $$ = gql.FragmentSpread{ Name: $2, Directives: $3 } }
        ;

inline_fragment
        : SPREAD ON type_condition directives_opt selection_set
                {
                        $$ = gql.Fragment{
                                TypeCondition: $3,
                                Directives: $4,
                                SelectionSet: $5,
                        }
                }
        | SPREAD directives_opt selection_set
                {
                        $$ = gql.Fragment{
                                Directives: $2,
                                SelectionSet: $3,
                        }
                }
        ;

fragment_definition
        : FRAGMENT fragment_name ON type_condition directives_opt selection_set
                {
                        $$ = gql.Fragment{
                                        Name: $2,
                                        TypeCondition: $4,
                                        Directives: $5,
                                        SelectionSet: $6,
                                };
                }
        ;

type_condition: type_name { $$ = $1 }
        ;

/* 2.2.7 Input Values */

value_const:    VALUE { $$ = $1 }
        |       boolean_value { $$ = $1 }
        |       VARIABLE { $$ = reflect.ValueOf($1) }
        |       list_value { $$ = reflect.ValueOf($1) }
        |       enum_value { $$ = reflect.ValueOf($1) }
        |       object_value { $$ = reflect.ValueOf($1) }
        |       NULL { $$ = reflect.ValueOf(nil) }
        ;

boolean_value:  TRUE { $$ = reflect.ValueOf(true) }
        |       FALSE { $$ = reflect.ValueOf(false) }
        ;

enum_value:     DIRECTIVE { $$ = $1; }
        |       ENUM { $$ = $1; }
        |       EXTEND { $$ = $1; }
        |       FRAGMENT { $$ = $1; }
        |       IDENTIFIER { $$ = $1; }
        |       IMPLEMENTS { $$ = $1; }
        |       INPUT { $$ = $1; }
        |       INTERFACE { $$ = $1; }
        |       MUTATION { $$ = $1; }
        |       ON { $$ = $1; }
        |       QUERY { $$ = $1; }
        |       SCALAR { $$ = $1; }
        |       SCHEMA { $$ = $1; }
        |       SUBSCRIPTION  { $$ = $1; }
        |       TYPE { $$ = $1; }
        |       UNION { $$ = $1; }
        ;

/* 2.2.7.6 List Value */

list_value: '[' value_list ']' { $$ = $2 }
        ;

value_list: /* nothing */ { $$ = nil }
        | value_list value_const { $$ = append($1, $2) }
        ;

/* 2.2.7.7 Object Value */
object_value: '{' object_field_list '}' { $$ = $2 }
        ;

object_field_list: /* nothing */ { $$ = nil }
        | object_field_list object_field { $$ = append($1, $2) }
        ;

object_field: name ':' value_const { $$ = gql.ObjectField{ Key: $1, Value: $3 }}
        ;

/* 2.2.9 Types */

type: type_name { $$ = gql.TypeDescription{ Name: $1 } }
        | list_type { $$ = $1 }
        | non_null_type { $$ = $1 }
        ;

type_name: name { $$ = $1 }
        ;

list_type: '[' type ']'
        {
                $2.Flags |= gql.List
                $$ = $2
        }
        ;

non_null_type: type_name '!'
                {
                        $$ = gql.TypeDescription{ Name: $1, Flags: gql.Required }
                }
        | list_type '!'
                {
                        $1.Flags |= gql.Required;
                        $$ = $1
                }
        ;

/* 2.2.10 Directives */

directives_opt: /* nothing */ { $$ = nil }
        | directives_opt directive { $$ = append($1, $2) }
        ;

directives: directives_opt directive { $$ = append($1, $2) }
        ;

directive: '@' name arguments_opt { $$ = gql.Directive{ Name: $2, Arguments: $3}; }
        ;

/* Schema Support */

schema: schema_definition type_definitions { $$ = gql.Schema{ OperationTypeDefinitions: $1, TypeDefinitions: $2 } }
      | type_definitions { $$ = gql.Schema{TypeDefinitions: $1}}
        ;


schema_definition: SCHEMA directives_opt '{' operation_type_definition_list '}'
        {
                $$ = $4
        }
        ;

operation_type_definition_list: /* nothing */ { $$ = nil }
        | operation_type_definition_list operation_type_definition { $$ = append($1, $2)}
        ;

operation_type_definition: operation_type ':' type_name
        {
                $$ = gql.OperationTypeDefinition{OpType: $1, Name: $3}
        }
        ;

type_definitions: /* nothing */ { $$ = make(map[string]gql.TypeDefinition) }
        | type_definitions scalar_type_definition { $1[$2.NamedType()] = $2; $$ = $1 }
        | type_definitions object_type_definition { $1[$2.NamedType()] = $2; $$ = $1 }
        | type_definitions interface_type_definition { $1[$2.NamedType()] = $2; $$ = $1 }
        | type_definitions union_type_definition { $1[$2.NamedType()] = $2; $$ = $1 }
        | type_definitions enum_type_definition { $1[$2.NamedType()] = $2; $$ = $1}
        | type_definitions input_object_type_definition { $1[$2.NamedType()] = $2; $$ = $1 }
        ;

scalar_type_definition: SCALAR name directives_opt { $$ = gql.ScalarDefinition{Name: $2, Directives: $3 } }
        ;

object_type_definition: TYPE name implements_interfaces_opt directives_opt '{' field_definition_list '}'
                {
                        $$ = gql.ObjectDefinition{
                                Name: $2,
                                Implements: $3,
                                Fields: $6,
                                Directives: $4,
                                }
                }
        ;

implements_interfaces_opt: /* nothing */ { $$ = nil }
        | IMPLEMENTS type_name_list { $$ = $2 }
        ;

type_name_list: type_name { $$ = []string{$1}}
        | type_name_list type_name { $$ = append($1, $2)}
        ;


field_definition: name arguments_definition_opt ':' type directives_opt
                {
                        $$ = gql.FieldDefinition{ Name: $1, Arguments: $2, Type: $4, Directives: $5 }
                }
        ;

field_definition_list: field_definition { $$ = []gql.FieldDefinition{$1} }
        | field_definition_list field_definition { $$ = append($1, $2) }
        ;

arguments_definition_opt: /* nothing */ { $$ = nil; }
        | arguments_definition { $$ = $1; }
        ;

arguments_definition: '(' input_value_definition_list ')' { $$ = $2; }
        ;

input_value_definition_list: input_value_definition { $$ = make(map[string]gql.InputValueDefinition); $$[$1.Name] = $1; }
        | input_value_definition_list input_value_definition { $1[$2.Name] = $2; $$ = $1 }
        ;

input_value_definition: name ':' type default_value_opt directives_opt
        {
                $$ = gql.InputValueDefinition{ Name: $1, Type: $3, Default: $4, Directives: $5 }
        }
        ;

interface_type_definition: INTERFACE name directives_opt '{' field_definition_list '}'
        {
                $$ = gql.InterfaceDefinition{ Name: $2, Fields: $5, Directives: $3 }
        }
        ;

union_type_definition: UNION name directives_opt '=' union_members
        {
                $$ = gql.UnionDefinition{ Name: $2, Types: $5, Directives: $3 }
        }
        ;

union_members: type_name { $$ = []string{$1} }
        | union_members '|' type_name { $$ = append($1, $3)}
        ;

enum_type_definition: ENUM name directives_opt '{' enum_value_definition_list '}'
        {
                $$ = gql.EnumDefinition{ Name: $2, Values: $5, Directives: $3 }
        }
        ;

enum_value_definition: name directives_opt { $$ = $1 }
        ;

enum_value_definition_list: enum_value_definition { $$ = []string{$1} }
        | enum_value_definition_list enum_value_definition { $$ = append($1, $2) }
        ;

input_object_type_definition: INPUT name directives_opt '{' input_value_definition_list '}'
        {
                $$ = gql.InputObjectDefinition{
                        Name: $2,
                        InputDefs: $5,
                        Directives: $3,
                };
        }

/* type extension
 directive def
directive locs
*/