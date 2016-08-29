package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type Form struct {
	Action string  `json:"action"`
	Method string  `json:"method"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Placeholder string `json:"placeholder"`
	Label       string `json:"label"`
	Required    bool   `json:"required"`
}

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.New("form.tmpl").Parse(`{{range .Fields}}
<div class="form-group">
    <label for="{{.Name}}">{{.Label}} {{if .Required}}*{{end}}</label>
    
    {{ if eq .Type "textarea" }}
        <textarea
            class="form-control"
            rows="5"
            name="{{.Name}}">{{.Placeholder}}</textarea>
    {{else}}
        <input 
            type="{{.Type}}" 
            class="form-control" 
            name="{{.Name}}" 
            placeholder="{{.Placeholder}}" 
            {{if .Required}}required{{end}}>
    {{end}}
</div>
{{end}}`)

	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to parse template"))
	}
}

func GenerateForm(form Form, writer io.Writer) error {
	return errors.Wrap(tmpl.Execute(writer, form), "unable to execute template")
}

func html(input, output string) error {
	form := &Form{}
	payload, err := ioutil.ReadFile(input)
	if err := json.Unmarshal(payload, form); err != nil {
		return errors.Wrap(err, "unable to unmarshal json")
	}

	file, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 755)
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}

	if err := tmpl.Execute(file, form); err != nil {
		return errors.Wrap(err, "unable to execute template")
	}

	return errors.Wrap(file.Close(), "unable to close file")
}

func main() {
	app := cli.NewApp()
	app.Name = "bootstrap-form-builder"
	app.Usage = "can build bootstrap form input, and code for cfdb ( a wordpress plugin ) so that building large forms stays easy"
	app.Commands = []cli.Command{
		{
			Name:  "html",
			Usage: "input.json output.html - reads the fields from input.json, and generates a list of input fields in output.html",
			Action: func(c *cli.Context) {
				if c.NArg() != 2 {
					log.Fatal(errors.New("invalid parameters count"))
				}

				args := c.Args()

				err := html(args.Get(0), args.Get(1))
				if err != nil {
					log.Fatal(errors.Wrap(err, "unable to execute"))
				}
			},
		},
	}
	app.Run(os.Args)
}
