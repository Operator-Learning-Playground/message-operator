package informer

import (
	"fmt"
	. "github.com/myoperator/messageoperator/pkg/send"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/tools/cache"
)

type DeploymentHandler struct {
}

func NewDeploymentHandler() *DeploymentHandler {
	return &DeploymentHandler{}
}

func (d DeploymentHandler) OnAdd(obj interface{}) {
	//fmt.Println("有走到发送！！")
	//dep := obj.(*v1.Deployment)
	//handlerLock.Lock()
	//defer handlerLock.Unlock()
	//GlobalSend.Send(fmt.Sprintf("deplyment already create: %s", dep.GetName()),
	//	fmt.Sprintf("deplyment already create: %s", dep.GetName()))
	//fmt.Println("发送成功")
	return
}

func (d DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	return
}

func (d DeploymentHandler) OnDelete(obj interface{}) {

	dep := obj.(*v1.Deployment)
	// TODO: 使用初始化实例 可以解决目前的bug
	s := NewSender()
	s.Send(fmt.Sprintf("deplyment already delete: %s", dep.GetName()),
		fmt.Sprintf("deplyment already delete: %s", dep.GetName()))
	fmt.Println("发送成功")

	// 发送邮件

	return
}

var _ cache.ResourceEventHandler = &DeploymentHandler{}

type PodHandler struct {
}

func NewPodHandler() *PodHandler {
	return &PodHandler{}
}

func (p PodHandler) OnAdd(obj interface{}) {
	return
}

func (p PodHandler) OnUpdate(oldObj, newObj interface{}) {
	return
}

func (p PodHandler) OnDelete(obj interface{}) {
	return
}

var _ cache.ResourceEventHandler = &PodHandler{}

type ServiceHandler struct {
}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}

func (s ServiceHandler) OnAdd(obj interface{}) {
	return
}

func (s ServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	return
}

func (s ServiceHandler) OnDelete(obj interface{}) {
	return
}

var _ cache.ResourceEventHandler = &ServiceHandler{}
