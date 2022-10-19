/* eslint-disable @typescript-eslint/no-var-requires */
const ExplicitImports = require('./rules/explicitImports.js')
const PrefixPrivateFunctionsWithUnderscore = require('./rules/prefixPrivateFunctionsWithUnderscore.js')
const PrefixInternalFunctionsWithUnderscore = require('./rules/prefixInternalFunctionsWithUnderscore.js')
const InheritedConstructorArgsNotInContractDefinition = require('./rules/inheritedConstructorArgsNotInContractDefinition.js')
const PrefixStorageVariablesWithSUnderscore = require('./rules/prefixStorageVariablesWithSUnderscore.js')
const PrefixImmutableVariablesWithI = require('./rules/prefixImmutableVariablesWithI.js')

module.exports = [
  ExplicitImports,
  PrefixPrivateFunctionsWithUnderscore,
  PrefixInternalFunctionsWithUnderscore,
  InheritedConstructorArgsNotInContractDefinition,
  PrefixStorageVariablesWithSUnderscore,
  PrefixImmutableVariablesWithI,
]
