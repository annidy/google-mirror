package render_test

import (
	"html/template"
	"os"
	"strings"
	"testing"
)

func TestMust1(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse("{{.Baz.Sun}}\n"))
	tmpl.Execute(os.Stdout, struct {
		Foo string
		Bar int
		Baz struct {
			Sun string
		}
	}{Foo: "hello", Bar: 42, Baz: struct{ Sun string }{Sun: "太阳"}})
}

func TestMust2(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse(`{{range .}}
		<li>{{.}}</li>
	{{end}}\n`))
	tmpl.Execute(os.Stdout, []string{"foo", "bar", "baz"})
}

func TestMust3(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse(`{{range $index, $element := .}}
		<li>{{$index}} - {{$element}}</li>
	{{end}}\n`))
	tmpl.Execute(os.Stdout, []string{"foo", "bar", "baz"})
}

func TestMust4(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse(`{{range .}}
		<li>{{.}}</li>
	{{end}}\n`))
	tmpl.Execute(os.Stdout, map[string]string{"foo": "1", "bar": "2", "baz": "3"})
}

func TestMust5(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse(`{{range $key, $value := .}}
		<li>{{$key}} - {{$value}}</li>
	{{end}}\n`))
	tmpl.Execute(os.Stdout, map[string]string{"foo": "1", "bar": "2", "baz": "3"})
}

func TestIf1(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse("{{.}} - {{if . == 10}}真{{else}}假{{end}}\n"))
	tmpl.Execute(os.Stdout, true)
	tmpl.Execute(os.Stdout, false)
	tmpl.Execute(os.Stdout, 1)
	tmpl.Execute(os.Stdout, 0)
	tmpl.Execute(os.Stdout, "1")
	tmpl.Execute(os.Stdout, "0")
	tmpl.Execute(os.Stdout, "")
	tmpl.Execute(os.Stdout, nil)
}

func TestIf2(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse("{{.}} - {{if eq . 10}}真{{else}}假{{end}}\n"))
	tmpl.Execute(os.Stdout, 10)
	tmpl.Execute(os.Stdout, 11)
}

func TestIf3(t *testing.T) {
	tmpl := template.New("foo")
	tmpl = template.Must(tmpl.Parse("{{.}} - {{if and eq .Bar .Foo}}真{{else}}假{{end}}\n"))
	tmpl.Execute(os.Stdout, struct {
		Foo int
		Bar int
	}{Foo: 11, Bar: 10})
}

func TestFunc1(t *testing.T) {
	tmpl := template.New("foo").Funcs(template.FuncMap{"upper": func(s string) string { return strings.ToUpper(s) }})
	tmpl = template.Must(tmpl.Parse("{{.}} - {{. | upper}}\n"))
	tmpl.Execute(os.Stdout, "foo")
}
