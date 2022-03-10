package v1alpha1

import (
	"github.com/mark8s/crd-demo/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"time"
)

func WatchResources(clientSet ExampleVAlphaInterface) cache.Store {
	store, controller := cache.NewInformer(&cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return clientSet.Projects("default").List(options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return clientSet.Projects("default").Watch(options)
		},
	}, &v1alpha1.Project{}, 1*time.Minute, cache.ResourceEventHandlerFuncs{})

	go controller.Run(wait.NeverStop)
	return store
}
