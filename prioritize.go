package main

import (
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kube-scheduler/extender/v1"
)

type Prioritize struct {
	Name string
	Func func(pod v1.Pod, nodes []v1.Node) (*schedulerapi.HostPriorityList, error)
}

func (p Prioritize) Handler(args schedulerapi.ExtenderArgs) (*schedulerapi.HostPriorityList, error) {
	return p.Func(*args.Pod, args.Nodes.Items)
}
