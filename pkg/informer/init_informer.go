package informer

import (
	"context"
	"fmt"
	"github.com/myoperator/messageoperator/pkg/k8sconfig"
	"github.com/myoperator/messageoperator/pkg/send"
	"github.com/myoperator/messageoperator/pkg/workqueue"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"log"
)

type K8sConfig struct {
	// Queue 工作队列接口
	workqueue.Queue
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{
		Queue: workqueue.NewWorkQueue(5),
	}
}

func (k *K8sConfig) InitClient() kubernetes.Interface {

	client, err := kubernetes.NewForConfig(k8sconfig.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// InitInformerFactory 初始化informer对象
func (k *K8sConfig) InitInformerFactory() informers.SharedInformerFactory {

	depHandler := NewDeploymentHandler(k.Queue)
	podHandler := NewPodHandler(k.Queue)
	svcHandler := NewServiceHandler(k.Queue)

	fact := informers.NewSharedInformerFactory(k.InitClient(), 0)
	deploymentInformer := fact.Apps().V1().Deployments()
	deploymentInformer.Informer().AddEventHandler(depHandler)

	podInformer := fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(podHandler)

	serviceInformer := fact.Core().V1().Services()
	serviceInformer.Informer().AddEventHandler(svcHandler)
	fact.Start(wait.NeverStop)
	klog.Info("informer start !!")
	k.start(context.Background())

	return fact
}

// Start 启动工作队列
func (k *K8sConfig) start(ctx context.Context) {
	klog.Info("worker queue start...")
	go func() {
		for {
			select {
			case <-ctx.Done():
				klog.Info("exit work queue...")
				k.Close()
				return
			default:
			}

			// 不断由队列中获取元素处理
			obj, err := k.Pop()
			if err != nil {
				klog.Errorf("work queue pop error: %s\n", err)
				continue
			}

			// 如果自己的业务逻辑发生问题，可以重新放回队列。
			if err = k.handleObject(obj); err != nil {
				klog.Errorf("handle obj from work queue error: %s\n", err)
				// 重新入列
				_ = k.ReQueue(obj)
			} else {
				// 完成就结束
				k.Finish(obj)
			}
		}
	}()
}

// handleObject 处理 work queue 传入对象
func (k *K8sConfig) handleObject(obj *workqueue.QueueResource) error {

	// FIXME: 不用每次的实例化，这里可以优化
	// 发送 email
	ss := send.NewFeishuClient()
	err := ss.Send(obj)
	if err != nil {
		klog.Errorf("send feishu error: %s", obj.Kind, obj.Name, obj.Namespace, obj.EventType, err)
		return err
	}

	// 发送 email
	s := send.NewSender()
	err = s.Send(fmt.Sprintf("your cluster resource is changing!"),
		fmt.Sprintf("[%s](%s/%s) is [%s]", obj.Kind, obj.Name, obj.Namespace, obj.EventType))
	if err != nil {
		klog.Errorf("send [%s](%s/%s) [%s] error: %s", obj.Kind, obj.Name, obj.Namespace, obj.EventType, err)
		return err
	}
	klog.Info("send [%s](%s/%s) [%s] success...", obj.Kind, obj.Name, obj.Namespace, obj.EventType)

	// 发送 飞书

	return nil
}
