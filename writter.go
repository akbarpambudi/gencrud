package crud

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

func WriteRepositoryFiles() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for _, entityConfig := range entities {
		crudGenerator := NewCrudGenerator(entityConfig)
		var fileName bytes.Buffer
		fileName.WriteString(cwd)
		fileName.WriteString(strings.ToLower(crudGenerator.State.RepositoryName))
		fileName.WriteString(".go")
		generatedRepoCode := crudGenerator.Generate()
		err := ioutil.WriteFile(fileName.String(), []byte(generatedRepoCode), 0644)
		if err != nil {
			panic(err)
		}
	}
}
