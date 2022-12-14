package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	g := rpc{}
	protogen.Options{}.Run(g.Generate)
}

type rpc struct{}

// Generate generate service code
func (md *rpc) Generate(plugin *protogen.Plugin) error {

	for _, f := range plugin.Files {
		if len(f.Services) == 0 {
			continue
		}
		fileName := f.GeneratedFilenamePrefix + ".svr.go"
		t := plugin.NewGeneratedFile(fileName, f.GoImportPath)
		t.P("//Code generated by protoc-gen-tinrpc")
		t.P()
		pkg := fmt.Sprintf("package %s", f.GoPackageName)
		t.P(pkg)
		t.P()
		for _, s := range f.Services {
			serviceCode := fmt.Sprintf(`%stype%s struct{}`,
				getComments(s.Comments), s.Desc.Name())
			t.P(serviceCode)
			t.P()
			for _, m := range s.Methods {
				funcCode := fmt.Sprintf(`
				%sfunc	(this *%s) %s(args *%s,  	reply *%s)error{
					//define your service ...
					return nil
				}`,
					getComments(m.Comments),
					s.Desc.Name(),
					m.Desc.Name(),
					m.Input.Desc.Name(),
					m.Output.Desc.Name())
				t.P(funcCode)
			}
		}
	}
	return nil
}

// getComments get comment details
func getComments(comments protogen.CommentSet) string {
	c := make([]string, 0)
	c = append(c, strings.Split(string(comments.Leading), "\n")...)
	c = append(c, strings.Split(string(comments.Trailing), "\n")...)

	res := ""
	for _, comment := range c {
		if strings.TrimSpace(comment) == "" {
			continue
		}
		res += "//" + comment + "\n"
	}
	return res
}
