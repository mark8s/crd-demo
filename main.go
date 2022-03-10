package main

import (
	"flag"
	"fmt"
	. "github.com/mark8s/crd-demo/api/types/v1alpha1"
	"github.com/mark8s/crd-demo/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	"time"
)

var config *rest.Config

func init() {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	var err error
	config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err)
	}
}

func main() {
	AddToScheme(scheme.Scheme)
	projectClient, err := v1alpha1.NewForConfig(config)
	if err != nil {
		return
	}

	projectList := &ProjectList{}
	projectList, err = projectClient.Projects("default").List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(projectList)

	store := v1alpha1.WatchResources(projectClient)
	for {
		list := store.List()
		fmt.Printf("project in store: %d\n", len(list))
		time.Sleep(2 * time.Second)
	}

}
