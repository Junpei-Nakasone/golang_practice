package main

// https://gist.github.com/rms1000watt/308ce7e525ebbf5981275981fa002a94

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

type School struct {
	Students []Person
	Name     string
}

func main() {
	templateStr := `Hello World:
My Name is: {{.Name}}
My Age is: {{.Age}}
{{ if and (eq .Age 33) (eq .Name "Mary")}}
Verified: I am 33 and my name is Mary.
{{ else }}
Not verified
{{ end }}
`

	templateStr2 := `Hello School:
School Name is: {{.Name}}
Students are: {{ range $student := .Students }}{{ $student.Name }}-{{ $student.Age }}, {{ end }}
`
	mary := Person{
		Name: "Mary",
		Age:  33,
	}

	joe := Person{
		Name: "Joe",
		Age:  33,
	}

	school := School{
		Name: "Ocean College",
		Students: []Person{
			mary,
			joe,
		},
	}

	t := template.New("Template Title...")
	t, err := t.Parse(templateStr)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(os.Stdout, mary)
	t.Execute(os.Stdout, joe)

	t, err = template.New("Template 2 Title...").Parse(templateStr2)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(os.Stdout, school)
}
