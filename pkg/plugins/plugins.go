package plugins

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// 插件名称
const Name = "wanna-scheduler"

type Args struct {
	FavoriteColor  string `json:"favorite_color,omitempty"`
	FavoriteNumber int    `json:"favorite_number,omitempty"`
	ThanksTo       string `json:"thanks_to,omitempty"`
}

type Sample struct {
	args   *Args
	handle framework.Handle
}

var _ framework.FilterPlugin = &Sample{}
var _ framework.PreBindPlugin = &Sample{}

func (s *Sample) Name() string {
	return Name
}

// PreFilter(ctx context.Context, state *CycleState, p *v1.Pod) (*PreFilterResult, *Status)
//func (s *Sample) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
//	klog.Infof("prefilter pod: %v", pod.Name)
//	return framework.NewStatus(framework.Success, "")
//}

// Filter(ctx context.Context, state *CycleState, pod *v1.Pod, nodeInfo *NodeInfo) *Status
func (s *Sample) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.Infof("filter pod: %v, node: %v", pod.Name, nodeInfo.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

// PreBind(ctx context.Context, state *CycleState, p *v1.Pod, nodeName string) *Status
func (s *Sample) PreBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	nodeInfo, error := s.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
	if error != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("prebind get node info error: %+v", nodeName))
	}
	klog.Infof("prebind node info: %+v", nodeInfo.Node())
	return framework.NewStatus(framework.Success, "")
}

// func(configuration runtime.Object, f framework.Handle) (framework.Plugin, error)
func New(configuration runtime.Object, f framework.Handle) (framework.Plugin, error) {
	args := &Args{}
	instance := configuration.DeepCopyObject()
	klog.Infof("instance str %+v", instance)
	obj, _ := instance.(*runtime.Unknown)
	klog.Infof("obj raw str %+v", string(obj.Raw))
	if err := json.Unmarshal(obj.Raw, args); err != nil {
		klog.Infof("json.Unmarshal err: %+v", err)
	}
	//if err := framework.DecodeInto(configuration, args); err != nil {
	//	return nil, err
	//}
	klog.Infof("get plugin config args: %+v", args)
	return &Sample{
		args:   args,
		handle: f,
	}, nil
}
