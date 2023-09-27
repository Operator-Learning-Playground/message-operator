package controller

import (
	"context"
	"github.com/go-logr/logr"
	messagev1alpha1 "github.com/myoperator/messageoperator/pkg/apis/message/v1alpha1"
	"github.com/myoperator/messageoperator/pkg/sysconfig"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type MessageController struct {
	client.Client
	EventRecorder record.EventRecorder // 事件管理器
	logr.Logger
}

func NewMessageController(eventRecorder record.EventRecorder, log logr.Logger) *MessageController {
	return &MessageController{EventRecorder: eventRecorder, Logger: log}
}

// Reconcile 调协loop
func (r *MessageController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	message := &messagev1alpha1.Message{}
	err := r.Get(ctx, req.NamespacedName, message)
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			klog.Error("get message error: ", err)
			return reconcile.Result{}, err
		}
		// 如果未找到的错误，不再进入调协
		return reconcile.Result{}, nil
	}
	klog.Info(message)

	// 当前正在删
	if !message.DeletionTimestamp.IsZero() {
		klog.Info("delete the message....")
		return reconcile.Result{}, nil
	}

	err = sysconfig.AppConfig(message)
	if err != nil {
		klog.Error("appconfig error: ", err)
		r.EventRecorder.Event(message, corev1.EventTypeWarning, "UpdateFailed", "update app config fail...")
		return reconcile.Result{}, nil
	}
	r.EventRecorder.Event(message, corev1.EventTypeNormal, "Update", "update app config...")

	klog.Info("CreateOrUpdate ", "Message: ", message.Name, "/", message.Namespace)
	return reconcile.Result{}, nil
}

// InjectClient 使用 controller-runtime 需要注入的client
func (r *MessageController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}
