package crud_test

import (
	"testing"

	"bitbucket.com/tunaiku/generator/crud"
)

func TestCrudGeneratorShouldGenerateCrudCode(t *testing.T) {
	expectedGeneratedCode := "type AccountRepository struct{}\nfunc (e AccountRepository) Entity() string {return \"AccountRepository\" }\n"
	config := crud.CrudGeneratorConfig{
		Entity: "Account",
	}
	crudGenerator := crud.NewCrudGenerator(config)
	accountRepositoryCode := crudGenerator.Generate()
	if accountRepositoryCode != expectedGeneratedCode {
		t.Errorf("expect crudGenerator.Generate() to return %s but it was %s", expectedGeneratedCode, accountRepositoryCode)
	}
}
