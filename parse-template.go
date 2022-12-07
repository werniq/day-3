package blog

import (
	"net/http"
	"text/template"
)

var (
	name   = ""
	skills = []string{}
)

func changeName(_name string) string {
	name = _name
	return name
}

func addOneSkills(newSkill string) []string {
	skills = append(skills, newSkill)
	return skills
}

func changeSkill(newSkills []string) []string {
	skills = nil
	skills = newSkills
	return skills
}

func Skills() []string {
	return skills
}

func Name() string {
	return name
}

func main(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.gohtml")

	if err != nil {
		panic(err)
	}
	data := struct {
		Name   string
		Skills []string
	}{
		Name:   Name(),
		Skills: Skills(),
	}

	errr := t.Execute(w, data)

	if errr != nil {
		panic(err)
	}
}
