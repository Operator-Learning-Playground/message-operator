package controller

import (
	"context"
	messagev1alpha1 "github.com/myoperator/messageoperator/pkg/apis/message/v1alpha1"
	"github.com/myoperator/messageoperator/pkg/sysconfig"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)


type MessageController struct {
	client.Client
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

// Reconcile 调协loop
func (r *MessageController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	message := &messagev1alpha1.Message{}
	err := r.Get(ctx, req.NamespacedName, message)
	if err != nil {
		return reconcile.Result{}, err
	}
	klog.Info(message)

	err = sysconfig.AppConfig(message)
	if err != nil {
		return reconcile.Result{}, nil
	}

	return reconcile.Result{}, nil
}

// InjectClient 使用controller-runtime 需要注入的client
func (r *MessageController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}

// TODO: 删除逻辑并未处理
