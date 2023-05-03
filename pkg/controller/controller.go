package controller

import (
	"context"
	messagev1alpha1 "github.com/myoperator/messageoperator/pkg/apis/message/v1alpha1"
	"github.com/myoperator/messageoperator/pkg/sysconfig"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/recorder"
)

var (
	recorderProvider recorder.Provider
)

type MessageController struct {
	client.Client
	EventRecord record.EventRecorder
}

func NewMessageController() *MessageController {
	r := recorderProvider.GetEventRecorderFor("message-operator")
	return &MessageController{EventRecord: r}
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

	mutateMessageRes, err := controllerutil.CreateOrUpdate(ctx, r.Client, message, func() error {
		err = sysconfig.AppConfig(message)
		if err != nil {
			klog.Error("appconfig error: ", err)
			r.EventRecord.Event(message, corev1.EventTypeWarning, "UpdateFail", "appconfig update fail...")
			return err
		}
		// 加入事件
		r.EventRecord.Event(message, corev1.EventTypeNormal, "Update", "appconfig update...")
		return nil
	})
	if err != nil {
		klog.Error("reconcile error: ", err)
		return reconcile.Result{}, err
	}
	klog.Info("CreateOrUpdate ", "Message ", mutateMessageRes)
	return reconcile.Result{}, nil
}

// InjectClient 使用controller-runtime 需要注入的client
func (r *MessageController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}

// TODO: 删除逻辑并未处理
