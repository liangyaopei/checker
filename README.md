# Overview
This repo provides flexible, configurable validation method when validating struct in Golang.

# Install
```
go get -u github.com/liangyaopei/checker
```

# Description
## Rule
Users can customize their various validation rules attached every filed in struct,
 including complex,composite rules.
Below rules all accepts a `fieldExpr` to locate the field in struct, and:
- `NewAndRule` accepts an array of rules, it passed when all rules are passed.
- `NewAndRule` accepts an array of rules, it passed when any rule is passed.
- `NewSliceRule` accepts slice/array input and a rule, applies the to every element in array/slice.
- `NewLengthRule` accepts slice/array/string/map, it checks the length of input is less than or equal to limit.
- `NewEnumRuleXXX` accepts a slice, it checks whether value of related struct filed exist in input enum slice. It has
  `NewEnumRuleString`,`NewEnumRuleInt`, `NewEnumRuleUint`,`NewEnumRuleFloat`.
- `NewEqRuleXXX` accepts a value, it checks whether value of related struct filed is equal to value. It 
    has `NewEnumRuleString`, `NewEnumRuleInt`, `NewEnumRuleUint`,`NewEnumRuleFloat`.
- `NewNotEqRuleXXX` is similar to `NewEqRuleXXX`. It checks whether value of related struct filed is not equal to value.
- `NewRangeRuleXXX` accepts min and max value, it checks whether value of related struct filed is `min<=v<=max`. It has
   `NewRangeRuleInt`, `NewRangeRuleUint`,`NewRangeRuleFloat`.
# Example
