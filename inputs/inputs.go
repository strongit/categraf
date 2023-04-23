package inputs

import (
	"flashcat.cloud/categraf/config"
	"flashcat.cloud/categraf/types"
)

type Initializer interface {
	Init() error
}

type SampleGatherer interface {
	Gather(*types.SampleList)
}

type Dropper interface {
	Drop()
}

type InstancesGetter interface {
	GetInstances() []Instance
}

func MayInit(t interface{}) error {
	if initializer, ok := t.(Initializer); ok {
		return initializer.Init()
	}
	return nil
}

func MayGather(t interface{}, slist *types.SampleList) {
	if gather, ok := t.(SampleGatherer); ok {
		gather.Gather(slist)
	}
}

func MayDrop(t interface{}) {
	if dropper, ok := t.(Dropper); ok {
		dropper.Drop()
	}
}

func MayGetInstances(t interface{}) []Instance {
	if instancesGetter, ok := t.(InstancesGetter); ok {
		return instancesGetter.GetInstances()
	}
	return nil
}

type (
	Cloneable interface {
		Clone() Input
	}

	Input interface {
		Cloneable
		Name() string
		GetLabels() map[string]string
		GetInterval() config.Duration
		InitInternalConfig() error
		Process(*types.SampleList) *types.SampleList
	}
)

type Creator func() Input

var InputCreators = map[string][]Creator{}

func Add(name string, creator Creator) {

	if _, ok := InputCreators[name]; !ok {
		InputCreators[name] = []Creator{creator}
	} else {
		InputCreators[name] = append(InputCreators[name], creator)
	}
}

type Instance interface {
	GetLabels() map[string]string
	GetIntervalTimes() int64
	InitInternalConfig() error
	Process(*types.SampleList) *types.SampleList
}
