package godjan

type Plugin interface {
	Name() string
	Models() []Model
}
