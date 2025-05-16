package shadowing

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var ShadowedVarAnalyzer = &analysis.Analyzer{
	Name: "shadowedvars",
	Doc:  "reports shadowed variables in inner scopes",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	visitedScopes := make(map[*types.Scope]bool)

	var dfs func(scope *types.Scope, shadowedVars map[string]token.Pos)
	dfs = func(scope *types.Scope, parentsVars map[string]token.Pos) {
		if scope == nil || visitedScopes[scope] {
			return
		}
		visitedScopes[scope] = true

		nodeVars := map[string]token.Pos{}

		for i := 0; i < len(scope.Names()); i++ {
			name := scope.Names()[i]
			if obj := scope.Lookup(name); obj != nil {
				if _, ok := obj.(*types.Var); !ok {
					continue
				}
				if prevPos, exists := parentsVars[name]; exists {
					pass.Reportf(obj.Pos(), "variable %q shadows an existing variable declared at %v", name, pass.Fset.Position(prevPos))
				} else {
					nodeVars[name] = obj.Pos()
				}
			}
		}

		for i := 0; i < scope.NumChildren(); i++ {
			unionVars := map[string]token.Pos{}
			for key, value := range parentsVars {
				unionVars[key] = value
			}

			for key, value := range nodeVars {
				unionVars[key] = value
			}

			dfs(scope.Child(i), unionVars)
		}
	}

	for _, file := range pass.Files {
		if scope := pass.TypesInfo.Scopes[file]; scope != nil {
			shadowedVars := map[string]token.Pos{}
			dfs(scope, shadowedVars)
		}
	}

	return nil, nil
}
