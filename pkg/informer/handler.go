package informer

import (
	"github.com/myoperator/messageoperator/pkg/workqueue"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

// TODO: 这里可以改成 annotation来过滤，不要所有的资源都发送消息
// TODO: 配置不同资源的email通知

const (
	MessageAnnotationKey   = "message"
	MessageAnnotationValue = "api.practice.com/send"
)

type DeploymentHandler struct {
	// Queue 工作队列接口
	workqueue.Queue
}

func NewDeploymentHandler(q workqueue.Queue) *DeploymentHandler {
	return &DeploymentHandler{Queue: q}
}

func (d DeploymentHandler) OnAdd(obj interface{}) {

	return
}

func (d DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {

	dep := newObj.(*appsv1.Deployment)
	_, ok := dep.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}

	obj, ok := newObj.(runtime.Object)
	if !ok {
		return
	}
	rr := &workqueue.QueueResource{
		Object:    obj,
		Name:      dep.GetName(),
		Namespace: dep.GetNamespace(),
		Kind:      "Deployment",
		EventType: workqueue.UpdateEvent,
	}
	d.Push(rr)

	//// TODO: 使用初始化实例 可以解决目前的bug
	//// 发送邮件
	//s := NewSender()
	//err := s.Send(fmt.Sprintf("deplyment is updated: %s", dep.GetName()),
	//	fmt.Sprintf("deplyment is updated: %s", dep.GetName()))
	//if err != nil {
	//	klog.Error("send deployment update error: ", err)
	//	return
	//}
	//klog.Info("send deployment update success...")

	return
}

func (d DeploymentHandler) OnDelete(obj interface{}) {

	dep := obj.(*appsv1.Deployment)
	// 判断是否有特殊的annotation，没有就不进行业务逻辑
	_, ok := dep.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	objj, ok := obj.(runtime.Object)
	if !ok {
		return
	}
	rr := &workqueue.QueueResource{
		Object:    objj,
		Name:      dep.GetName(),
		Namespace: dep.GetNamespace(),
		Kind:      "Deployment",
		EventType: workqueue.DeleteEvent,
	}
	d.Push(rr)

	return
}

var _ cache.ResourceEventHandler = &DeploymentHandler{}

type PodHandler struct {
	// Queue 工作队列接口
	workqueue.Queue
}

func NewPodHandler(q workqueue.Queue) *PodHandler {
	return &PodHandler{Queue: q}
}

func (p PodHandler) OnAdd(obj interface{}) {
	return
}

func (p PodHandler) OnUpdate(oldObj, newObj interface{}) {

	pod := newObj.(*corev1.Pod)
	_, ok := pod.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	obj, ok := newObj.(runtime.Object)
	if !ok {
		return
	}
	rr := &workqueue.QueueResource{
		Object:    obj,
		Name:      pod.GetName(),
		Namespace: pod.GetNamespace(),
		Kind:      "Pod",
		EventType: workqueue.UpdateEvent,
	}
	p.Push(rr)
	return
}

func (p PodHandler) OnDelete(obj interface{}) {

	pod := obj.(*corev1.Pod)
	_, ok := pod.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	objj, ok := obj.(runtime.Object)
	if !ok {
		return
	}
	rr := &workqueue.QueueResource{
		Object:    objj,
		Name:      pod.GetName(),
		Namespace: pod.GetNamespace(),
		Kind:      "Pod",
		EventType: workqueue.DeleteEvent,
	}
	p.Push(rr)

	return
}

var _ cache.ResourceEventHandler = &PodHandler{}

type ServiceHandler struct {
	// Queue 工作队列接口
	workqueue.Queue
}

func NewServiceHandler(q workqueue.Queue) *ServiceHandler {
	return &ServiceHandler{Queue: q}
}

func (s ServiceHandler) OnAdd(obj interface{}) {
	return
}

func (s ServiceHandler) OnUpdate(oldObj, newObj interface{}) {

	svc := newObj.(*corev1.Service)
	_, ok := svc.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	obj, ok := newObj.(runtime.Object)
	if !ok {
		return
	}
	rr := &workqueue.QueueResource{
		Object:    obj,
		Name:      svc.GetName(),
		Namespace: svc.GetNamespace(),
		Kind:      "Service",
		EventType: workqueue.UpdateEvent,
	}
	s.Push(rr)

	return
}

func (s ServiceHandler) OnDelete(obj interface{}) {

	svc := obj.(*corev1.Service)
	_, ok := svc.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	// TODO: 使用初始化实例 可以解决目前的bug
	objj, ok := obj.(runtime.Object)
	if !ok {
		return
	}
	rr := &workqueue.QueueResource{
		Object:    objj,
		Name:      svc.GetName(),
		Namespace: svc.GetNamespace(),
		Kind:      "Service",
		EventType: workqueue.DeleteEvent,
	}
	s.Push(rr)

	return
}

var _ cache.ResourceEventHandler = &ServiceHandler{}
