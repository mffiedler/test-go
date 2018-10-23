package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	examplecomclientset "github.com/hongkailiu/test-go/codegen/pkg/client/clientset/versioned"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	exampleClient, err := examplecomclientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %v", err)
	}

	list, err := exampleClient.AppV1alpha1().SVTGos("ttt").List(metav1.ListOptions{})
	if err != nil {
		glog.Fatalf("Error listing all databases: %v", err)
	}

	for _, svtGo := range list.Items {
		fmt.Printf("svtGo %s with Size %d\n", svtGo.Name, svtGo.Spec.Size)
	}
}
