package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"

)

var a = 1

func main() {
	bar()
	bar()
	var bar string = "bar"
}

func bar() {
	fmt.Println(time.Now())
}`

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.Scope.Outer)
	fmt.Println(f.Scope.Objects)
	ast.Inspect(f, func(n ast.Node) bool {
		fmt.Println("+++++++++++++++++++++++++++++++++++++++")
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
			if x.Obj != nil && x.Obj.Kind == ast.Fun {
				fmt.Print("type:", x.Obj.Kind)
				if x.Name == "bar" {
					x.Name = "foo"
				}
			}
		case *ast.FuncDecl:
			fmt.Println("xxxxxxxxxxxxxxxxxx", x.Name, "xxxxxxxxxxxxxxxxxxxxxxx")
			if x.Name.Name == "bar" {
				x.Name.Name = "foo"
			}
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		fmt.Println("+++++++++++++++++++++++++++++++++++++++")
		return true
	})

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
	fmt.Println("===================================")
	for _, decl := range f.Decls {
		if gen, ok := decl.(*ast.FuncDecl); ok {
			fmt.Println(gen.Name)
			for _, list := range gen.Body.List {
				fmt.Println(src[list.Pos()-1 : list.End()])
			}

		}

		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
			fmt.Println(src[gen.Pos()-1 : gen.End()])

			f.Decls = append(f.Decls, decl.(*ast.GenDecl))
		}

		if gen, ok := decl.(*ast.BadDecl); ok {
			fmt.Println("error: ", src[gen.Pos()-1:gen.End()])
		}
	}
	fmt.Println("===================================")

	ffff, err := os.Create("test.go")
	if err != nil {
		panic(err)
	}
	defer ffff.Close()
	list := []*ast.Comment{
		{
			Slash: 0,
			Text:  "// xxxxxxxxx",
		},
		{
			Slash: 5,
			Text:  "// 123",
		},
		{
			Slash: 3,
			Text:  "// 1232144",
		},
	}

	f.Imports[1].Comment = &ast.CommentGroup{
		List: list,
	}
	ast.SortImports(fset, f)

	_ = printer.Fprint(ffff, fset, f)

	//for _, ident := range f.Unresolved {
	//	fmt.Println(ident.Name, ident.NamePos)
	//}
	//fmt.Println("===============", len(f.Unresolved), "====================")
	//
	//if err := ast.Print(fset, f); err != nil {
	//	panic(err)
	//}
	// output:
	//
	// "fmt"
	// "time"
}
