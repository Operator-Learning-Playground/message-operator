package informer

import (
	"fmt"
	. "github.com/myoperator/messageoperator/pkg/send"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

// TODO: 这里可以改成 annotation来过滤，不要所有的资源都发送消息

const (
	MessageAnnotationKey = "message"
	MessageAnnotationValue = "api.practice.com/send"
)

type DeploymentHandler struct {
}

func NewDeploymentHandler() *DeploymentHandler {
	return &DeploymentHandler{}
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
	// TODO: 使用初始化实例 可以解决目前的bug
	// 发送邮件
	s := NewSender()
	s.Send(fmt.Sprintf("deplyment is updated: %s", dep.GetName()),
		fmt.Sprintf("deplyment is updated: %s", dep.GetName()))
	klog.Info("发送成功")

	return
}

func (d DeploymentHandler) OnDelete(obj interface{}) {

	dep := obj.(*appsv1.Deployment)
	// 判断是否有特殊的annotation，没有就不进行业务逻辑
	_, ok := dep.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	// 发送邮件
	s := NewSender()
	s.Send(fmt.Sprintf("deplyment is deleted: %s", dep.GetName()),
		fmt.Sprintf("deplyment is deleted: %s", dep.GetName()))
	klog.Info("发送成功")


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

	pod := newObj.(*corev1.Pod)
	_, ok := pod.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	// 发送邮件
	s := NewSender()
	s.Send(fmt.Sprintf("pod is updated: %s", pod.GetName()),
		fmt.Sprintf("pod is updated: %s", pod.GetName()))

	klog.Info("发送成功")
	return
}

func (p PodHandler) OnDelete(obj interface{}) {

	pod := obj.(*corev1.Pod)
	_, ok := pod.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	// 发送邮件
	s := NewSender()
	s.Send(fmt.Sprintf("pod is deleted: %s", pod.GetName()),
		fmt.Sprintf("pod is deleted: %s", pod.GetName()))
	klog.Info("发送成功")

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

	svc := newObj.(*corev1.Service)
	_, ok := svc.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	// 发送邮件
	ss := NewSender()
	ss.Send(fmt.Sprintf("service is updated: %s", svc.GetName()),
		fmt.Sprintf("service is updated: %s", svc.GetName()))
	klog.Info("发送成功")

	return
}

func (s ServiceHandler) OnDelete(obj interface{}) {

	svc := obj.(*corev1.Service)
	_, ok := svc.GetAnnotations()[MessageAnnotationKey]
	if !ok {
		return
	}
	// TODO: 使用初始化实例 可以解决目前的bug
	// 发送邮件
	ss := NewSender()
	ss.Send(fmt.Sprintf("service is deleted: %s", svc.GetName()),
		fmt.Sprintf("service is deleted: %s", svc.GetName()))
	klog.Info("发送成功")

	return
}

var _ cache.ResourceEventHandler = &ServiceHandler{}
