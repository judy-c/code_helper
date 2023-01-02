package json_helper

type StructItemModel struct {
	ItemType 	string
	VarName		string
}

type StructTypeModel struct {
	TypeName 	string
	StructItems []StructItemModel
}




type NetworkModel struct {
	Name    string
	BandwidthProvided       float64
	BandwidthAlloc  float64
}

type InstancesItemModel struct {
	InstanceID      string
	Network NetworkModel
}

type RootModel struct {
	Sharing bool
	Instances       []InstancesItemModel
}
