package crud

import (
	"bytes"
	"fmt"
	"text/template"
)

type generatorState struct {
	RepositoryName string
}

type CrudGenerator struct {
	config CrudGeneratorConfig
	State  generatorState
}

func (c *CrudGenerator) InitiateState() {
	c.State.RepositoryName = fmt.Sprint(c.config.Entity, "Repository")
}

func (c CrudGenerator) generateMethods() string {
	var generatedCrudBuffer bytes.Buffer
	crudTemplate, err := template.New(c.config.Entity).Parse("func (e {{.State.RepositoryName}}) Entity() string {return \"{{.State.RepositoryName}}\" }")
	if err != nil {
		panic(err)
	}
	err = crudTemplate.Execute(&generatedCrudBuffer, c)
	if err != nil {
		panic(err)
	}
	return generatedCrudBuffer.String()
}
func (c CrudGenerator) generateStruct() string {
	var generatedCrudBuffer bytes.Buffer
	crudTemplate, err := template.New(c.config.Entity).Parse("type {{.State.RepositoryName}} struct{}")
	if err != nil {
		panic(err)
	}
	err = crudTemplate.Execute(&generatedCrudBuffer, c)
	if err != nil {
		panic(err)
	}
	return generatedCrudBuffer.String()
}
func (c CrudGenerator) Generate() string {
	var buffer bytes.Buffer
	buffer.WriteString(c.config.PackageName)
	buffer.WriteString(c.generateStruct())
	buffer.WriteString("\n")
	buffer.WriteString(c.generateMethods())
	buffer.WriteString("\n")

	return buffer.String()
}

func NewCrudGenerator(config CrudGeneratorConfig) *CrudGenerator {
	generator := &CrudGenerator{config: config}
	generator.InitiateState()
	return generator
}
