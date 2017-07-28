package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	eval "github.com/matteo-merzi/gopl/ch07/7.14"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml", "evaluate.gohtml"))
}

type Result struct {
	Expression string
	Variables  string
	Value      float64
	Err        string
}

func evaluate(w http.ResponseWriter, req *http.Request) {
	expStr := req.FormValue("expression")
	envStr := req.FormValue("variables")
	res := Result{
		Expression: expStr,
		Variables:  envStr,
	}
	env := eval.Env{}
	assignments := strings.Split(envStr, " ")
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			res.Err = fmt.Sprintf("incorrect variable assignment: %v", fields)
			tpl.ExecuteTemplate(w, "evaluate.gohtml", res)
			return
		}
		name, valueStr := fields[0], fields[1]
		fmt.Printf("name: %s, value: %s\n", name, valueStr)
		val, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			res.Err = fmt.Sprintf("unparsable value for %s", name)
			tpl.ExecuteTemplate(w, "evaluate.gohtml", res)
			return
		}
		env[eval.Var(name)] = val
	}

	exp, err := eval.Parse(expStr)
	if err != nil {
		res.Err = fmt.Sprintf("unparsable expression: %s", expStr)
		tpl.ExecuteTemplate(w, "evaluate.gohtml", res)
		return
	}
	fmt.Printf("expression: %s\n", exp.String())

	vars := make(map[eval.Var]bool)
	err = exp.Check(vars)
	if err != nil {
		for _, v := range vars {
			fmt.Println(v)
		}
		res.Err = fmt.Sprintf("error checking expression: %s", expStr)
		tpl.ExecuteTemplate(w, "evaluate.gohtml", res)
		return
	}
	res.Value = exp.Eval(env)
	fmt.Println(res)
	tpl.ExecuteTemplate(w, "evaluate.gohtml", res)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/evaluate", evaluate)
	http.ListenAndServe(":8080", nil)
}
