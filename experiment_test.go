package myasthurts_test

import (
	"go/parser"
	"go/token"

	myasthurts "github.com/jamillosantos/go-my-ast-hurts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("My AST Hurts", func() {

	It("should parse a User struct", func() {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "data/models2.sample", nil, parser.AllErrors)
		Expect(err).ToNot(HaveOccurred())
		Expect(f).ToNot(BeNil())
		Expect(f.Decls).ToNot(BeNil())

		env := &myasthurts.Environment{}
		myasthurts.Parse(f, env)

		// ---------- Test User struct - models2.sample ----------
		Expect(env.Packages).To(HaveLen(1))
		Expect(env.Packages[0].Structs).To(HaveLen(1))

		s := env.Packages[0].Structs[0]

		Expect(s.Name).To(Equal("User"))
		Expect(s.Fields).To(HaveLen(6))
		//Expect(s.Comment).To(Equal("User is a model.")) TODO
	})

	// This test depend "should parse a User struct"
	It("should parse fields of User struct", func() {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "data/models2.sample", nil, parser.AllErrors)
		Expect(err).ToNot(HaveOccurred())
		Expect(f).ToNot(BeNil())
		Expect(f.Decls).ToNot(BeNil())

		env := &myasthurts.Environment{}
		myasthurts.Parse(f, env)

		// ---------- Test User struct - models2.sample ----------
		fields := env.Packages[0].Structs[0].Fields

		Expect(fields[0].Name).To(Equal("ID"))
		Expect(fields[1].Name).To(Equal("Name"))
		Expect(fields[2].Name).To(Equal("Email"))
		Expect(fields[3].Name).To(Equal("Password"))
		Expect(fields[4].Name).To(Equal("CreatedAt"))
		Expect(fields[5].Name).To(Equal("UpdatedAt"))

	})

	/*It("should test somethhing", func() {

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "data/models1.sample", nil, parser.AllErrors)
		Expect(err).ToNot(HaveOccurred())
		Expect(f).ToNot(BeNil())
		Expect(f.Decls).ToNot(BeNil())

		fmt.Println("---------- Reading ----------")

		ast.Inspect(f, func(x ast.Node) bool {

			switch x.(type) {
			case *ast.StructType:
				s, ok := x.(*ast.StructType)
				if !ok {
					break
				}

				for _, fieldList := range s.Fields.List {

					if fieldList.Names[0].Name == "ID" {
						fieldList.Tag.Value = fmt.Sprintf("%s", "`json:\"IDTestChange\"`")
					}

					fmt.Printf("%s %s %s\n", fieldList.Names[0], fieldList.Type, fieldList.Tag.Value)

				}

				/*for _, field := range s.Fields.List {
					typeExpr := field.Type

					start := typeExpr.Pos() - 1
					end := typeExpr.End() - 1

					typeInSource := src[start:end]

					fmt.Println(typeInSource)
				}
			case *ast.FuncDecl:
				fmt.Println("---------- Reading FuncDecl ----------")
				s, ok := x.(*ast.FuncDecl)
				if !ok {
					break
				}

				params := ""
				for _, p := range s.Type.Params.List {
					for i := 0; i < len(p.Names); i++ {
						params += fmt.Sprintf("%s %s, ", p.Names[i], p.Type)
					}
				}
				params = strings.TrimSuffix(params, ", ")
				fmt.Printf("%s %s %b -> Params(%s)\n\n", s.Name.Obj.Kind, s.Name.Name, s.Type.Params.NumFields(), params)
			case *ast.TypeSpec:
				s, ok := x.(*ast.TypeSpec)
				if !ok {
					break
				}
				fmt.Printf("%s %s\n", s.Name.Obj.Kind, s.Name.String())
			}
			return true
		})

		fmt.Println("---------- Scope ----------")

		fmt.Println(f.Scope.String())

		fmt.Println(f.Scope.Objects["Test"].Pos())

		fmt.Printf("%s\n", fset.Position(f.Scope.Objects["Test"].Pos()))

		fmt.Println("---------- Change field ----------")

		cc := reflect.TypeOf(f.Scope.Objects["User"])
		//f1 := cc.Field(0)

		fmt.Println(cc)

	})*/
})

/*
if file.Imports != nil {
		fmt.Println("---------- Reading Imports/start ----------")

		importsStr := "import (\n"
		for _, i := range file.Imports {
			importsStr += "\t"
			if i == nil {
				fmt.Println("<nil>")
			} else {
				importsStr += fmt.Sprintf("%s %s", i.Name.String(), i.Path.Value)
			}
			importsStr += "\n"
		}
		importsStr += ")"
		fmt.Println(importsStr)
		fmt.Print("---------- Reading Imports/end ----------\n\n")
	}

	if file.Imports != nil {
		fmt.Println("---------- Reading Functions/start ----------")

		for _, i := range file.Decls {
			fn, ok := i.(*ast.FuncDecl)
			if !ok {
				continue
			}

			params := ""
			for _, p := range fn.Type.Params.List {
				for i := 0; i < len(p.Names); i++ {
					params += fmt.Sprintf("%s %s, ", p.Names[i], p.Type)
				}
			}
			params = strings.TrimSuffix(params, ", ")
			fmt.Printf("%s %s %b -> Params(%s)\n", fn.Name.Obj.Kind, fn.Name.Name, fn.Type.Params.NumFields(), params)
		}

		fmt.Print("---------- Reading Functions/end ----------\n\n")
	}

	if file.Decls != nil {
		fmt.Println("---------- Reading Struct/start ----------")

		var v visitor

		ast.Walk(v, file)

		/*ast.Inspect(file, func(x ast.Node) bool {
			s, okk := x.(*ast.StructType)
			if !okk {
				return true
			}
			for _, field := range s.Fields.List {
				fmt.Printf("%s %s %s\n", field.Names[0], field.Type, field.Tag.Value)
			}
			return false
		})

		fmt.Print("---------- Reading Struct/end ----------\n\n")

	}

	/*fmt.Printf("file: %v\n\n---\n\n", file)
	fmt.Println("--------------------")
	fmt.Print("Doc:")
	if file.Doc == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println(file.Doc.Text())
	}
	fmt.Println("Name:", file.Name.String())
	fmt.Println("Declarations")
	for _, d := range file.Decls {
		// fmt.Printf("%T\n", d)
		switch dcl := d.(type) {
		case *ast.GenDecl:
			fmt.Print("  Doc: ")
			if dcl.Doc == nil {
				fmt.Println("<nil>")
			} else {
				fmt.Println(dcl.Doc.Text())
			}

			for _, s := range dcl.Specs {
				fmt.Printf("%T\n", s)
				switch spec := s.(type) {
				case *ast.ImportSpec:
					fmt.Println("    ", spec.Name.String(), spec.Path.Value)
				}
			}
			// fmt.Println("Name: ", dcl.)
		default:
			fmt.Printf("  unknown type: %T\n", dcl)
		}
	}
	fmt.Println("--------------------")
	fmt.Println()
})*/
