package executor

import (
	"context"
	"fmt"
	"github.com/airren/echo-bio-backend/global"
	"github.com/airren/echo-bio-backend/service"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/airren/echo-bio-backend/model"
)

var JobChan chan model.AnalysisJob
var clientSet *kubernetes.Clientset
var informerFactory informers.SharedInformerFactory
var exector Executor

func init() {
	exector = Executor{}
}

func Run(ctx context.Context) {
	for {
		for job := range JobChan {
			newJob := job
			exector.CreateJob(ctx, &newJob)
		}
	}
}

func KubeInitializer() {
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		kubeconfig := filepath.Join("../conf", "kubeconf.yaml")
		if val := os.Getenv("KUBECONFIG"); len(val) != 0 {
			kubeconfig = val
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
			os.Exit(1)
		}
	}

	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("clinetset init failed")
		os.Exit(1)
	}

	informerFactory = informers.NewSharedInformerFactory(clientSet, time.Second*30)
}

type Executor struct {
}

func (e *Executor) CreateJob(ctx context.Context, analysisJob *model.AnalysisJob) (err error) {

	jobClient := clientSet.BatchV1().Jobs("echo-bio")
	jobImage, err := service.GetImageForAnalysisJob(analysisJob)
	if err != nil {
		log.Printf("this analysisJob doesn't bind a images: %v", analysisJob)
		return
	}

	command, err := service.GetCommandForAnalysisJob(analysisJob)
	if err != nil {
		global.Logger.Error("rend analysis_job command failed: ")
	}
	var backOffLimit int32 = 0

	jobSpec := &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("job-%d", analysisJob.Id),
			Namespace: "echo-bio",
		},
		Spec: v1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:    fmt.Sprintf("job-%d", analysisJob.Id),
							Image:   jobImage,
							Command: strings.Split(command, " "),
						},
					},
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
			BackoffLimit:   &backOffLimit,
			CompletionMode: nil,
		},
	}

	_, err = jobClient.Create(ctx, jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("failed to create K8s AnalysisJob err: ", err)
	}
	return
}
