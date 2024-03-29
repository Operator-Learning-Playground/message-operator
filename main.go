package main

import (
	messagev1alpha1 "github.com/myoperator/messageoperator/pkg/apis/message/v1alpha1"
	"github.com/myoperator/messageoperator/pkg/controller"
	"github.com/myoperator/messageoperator/pkg/httpserver"
	"github.com/myoperator/messageoperator/pkg/informer"
	"github.com/myoperator/messageoperator/pkg/k8sconfig"
	"github.com/myoperator/messageoperator/pkg/sysconfig"
	"time"

	//corev1 "k8s.io/api/core/v1"
	//"k8s.io/client-go/tools/record"
	_ "k8s.io/code-generator"
	"k8s.io/klog/v2"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

/*
	manager 主要用来管理Controller Admission Webhook 包括：
	访问资源对象的client cache scheme 并提供依赖注入机制 优雅关闭机制
	operator = crd + controller + webhook
*/

func main() {

	logf.SetLogger(zap.New())
	var d time.Duration = 0
	// 1. 管理器初始化
	mgr, err := manager.New(k8sconfig.K8sRestConfig(), manager.Options{
		Logger:     logf.Log.WithName("message-operator"),
		SyncPeriod: &d, // resync不设置触发
	})

	if err != nil {
		mgr.GetLogger().Error(err, "unable to set up manager")
		os.Exit(1)
	}

	// 2. ++ 注册进入序列化表
	err = messagev1alpha1.SchemeBuilder.AddToScheme(mgr.GetScheme())
	if err != nil {
		klog.Error(err, "unable add schema")
		os.Exit(1)
	}

	k8sConfig := informer.NewK8sConfig()
	k8sConfig.InitInformerFactory()

	// 3. 控制器相关
	messageCtl := controller.NewMessageController(mgr.GetEventRecorderFor("message-operator"), mgr.GetLogger())

	err = builder.ControllerManagedBy(mgr).For(&messagev1alpha1.Message{}).Complete(messageCtl)

	// 4. 载入业务配置
	if err = sysconfig.InitConfig(); err != nil {
		klog.Error("unable to load sysconfig error: ", err)
		os.Exit(1)
	}

	errC := make(chan error)

	// 5. 启动 controller 管理器
	go func() {
		klog.Info("controller start!! ")
		if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
			errC <- err
		}
	}()

	// 6 启动httpServer
	go func() {
		klog.Info("httpserver start!! ")
		if err = httpserver.HttpServer(); err != nil {
			errC <- err
		}
	}()

	// 阻塞
	getError := <-errC
	klog.Error(getError.Error())

}
