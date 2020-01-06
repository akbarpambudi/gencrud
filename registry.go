package crud

import "reflect"

var (
	entities []CrudGeneratorConfig
)

func RegisterEntity(entity interface{}) {
	entityType := reflect.TypeOf(entity)
	entities = append(entities, CrudGeneratorConfig{
		Entity:      entityType.Name(),
		PackageName: entityType.PkgPath(),
	})
}
