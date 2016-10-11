package godjan

import "fmt"

type DataSource interface {
	Init()
	Save(m Model, data DataEntity) error
	Get(Model, string) interface{}
}

type MemoryDataSource struct {
	Data map[string]interface{}
}

func (mds *MemoryDataSource) Init() {
	mds.Data = make(map[string]interface{}, 10)
	fmt.Println("initted")
}

func (mds MemoryDataSource) Save(m Model, data DataEntity) error {
	key := fmt.Sprintf("%s:%s", m.DbID(), data.DbID())
	mds.Data[key] = data
	return nil
}

func (mds MemoryDataSource) Get(m Model, id string) interface{} {
	key := fmt.Sprintf("%s:%s", m.DbID(), id)
	return mds.Data[key]
}
