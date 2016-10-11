package godjan

type Model interface {
	DbID() string
}

type DataEntity interface {
	DbID() string
}

type SampleData struct {
	ID    string
	value int
}

func (sd SampleData) DbID() string {
	return sd.ID
}
