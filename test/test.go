package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Main() {
	pkgPath := "Sickle/entity/database"
	pkg, err := parser.ParseDir(token.NewFileSet(), pkgPath, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, astFile := range pkg {
		for _, astObj := range astFile.Scope.Objects {
			if astObj.Kind == ast.Typ {
				typeSpec, ok := astObj.Decl.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				structName := typeSpec.Name.Name
				fmt.Println("Struct:", structName)

				for _, field := range structType.Fields.List {
					fieldName := ""
					if len(field.Names) > 0 {
						fieldName = field.Names[0].Name
					}

					fmt.Println("Field:", fieldName)
				}

				fmt.Println("-----------")
			}
		}
	}
}
