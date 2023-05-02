package informer

import (
	"github.com/myoperator/messageoperator/pkg/k8sconfig"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"log"
)

type K8sConfig struct{}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
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

	depHandler := NewDeploymentHandler()
	podHandler := NewPodHandler()
	svcHandler := NewServiceHandler()

	fact := informers.NewSharedInformerFactory(k.InitClient(), 0)
	deploymentInformer := fact.Apps().V1().Deployments()
	deploymentInformer.Informer().AddEventHandler(depHandler)

	podInformer := fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(podHandler)

	serviceInformer := fact.Core().V1().Services()
	serviceInformer.Informer().AddEventHandler(svcHandler)

	fact.Start(wait.NeverStop)
	klog.Info("informer start !!")

	return fact
}
