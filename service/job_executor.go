package service

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/airren/echo-bio-backend/dal"
	"github.com/airren/echo-bio-backend/model"
)

var JobChan chan model.AnalysisJob
var clientSet *kubernetes.Clientset
var informerFactory informers.SharedInformerFactory
var Exec = Executor{}

func KubeInitializer() {
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		kubeconfig := filepath.Join("./conf", "kubeconf.yaml")
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
	jobImage, err := GetImageForAnalysisJob(analysisJob)
	if err != nil {
		log.Printf("this analysisJob doesn't bind a images: %v", analysisJob)
		return
	}

	command, err := GetCommandForAnalysisJob(ctx, analysisJob)
	if err != nil {
		log.Print("rend analysis_job command failed: ")
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
							Name: fmt.Sprintf("job-%d", analysisJob.Id),
							Env: []apiv1.EnvVar{
								{"ACCOUNT_ID", analysisJob.AccountId, nil},
								{"ANALYSIS_JOB_ID", fmt.Sprint(analysisJob.Id), nil},
							},
							Image:           jobImage,
							ImagePullPolicy: "IfNotPresent",
							Command:         []string{"bash", "-c", command + ";ls -al && echo success ;sleep infinity"},
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

func GetCommandForAnalysisJob(ctx context.Context, job *model.AnalysisJob) (command string, err error) {

	command = "set -ex; mkdir -p /tmp/analysis_job/results && cd /tmp/analysis_job ; "

	staffDownload := "curl -s http://www.echo-bio.cn:8088/api/v1/internal/file_download/%v?account_id=${ACCOUNT_ID} -o %v;"

	compactAndUpload := `; tar -czvf results.tar.gz ./results;
curl -X POST http://www.echo-bio.cn:8088/api/v1/internal/file_upload?analysis_job_id=${ANALYSIS_JOB_ID}&account_id=${ACCOUNT_ID} \
--header "ACCOUNT_ID: ${ACCOUNT_ID}" \
--form 'file=@"/tmp/analysis_job/results.tar.gz"'`

	if len(job.Parameters) == 0 {
		err = job.TransferJsonToParameters()
		if err != nil {
			return
		}
	}

	algo, err := dal.GetAlgorithmByName(context.TODO(), job.Algorithm)
	if err != nil {
		log.Errorf("get algorithm by name failed: %v", err)
		return
	}

	algoParameters, err := dal.QueryParameter(ctx, &model.AlgoParameter{
		AlgorithmId: algo.Id,
	})
	if err != nil {
		log.Errorf("get algorithm's parameters failed: %v", err)
		return
	}

	fileParameterMap := make(map[string]string)
	for _, f := range algoParameters {
		if f.Type == model.ParamFile {
			if _, ok := job.Parameters[f.Name]; ok {
				fileParameterMap[f.Name] = job.Parameters[f.Name].(string)
			}
		}
	}
	// add download data files command
	for _, fileId := range fileParameterMap {
		command = command + fmt.Sprintf(staffDownload, fileId, fileId)
	}
	// use go template render the command
	algoCommand, err := GenerateAlgorithmCommand(job, algo)
	if err != nil {
		return
	}
	return command + algoCommand + compactAndUpload, err
}

func GenerateAlgorithmCommand(job *model.AnalysisJob, algo *model.Algorithm) (command string, err error) {
	tmpl, err := template.New("job_command").Parse(algo.Command)
	if err != nil {
		log.Fatalf("parse the command template failed: %v", err)
		return "", err
	}
	commandStr := bytes.Buffer{}
	err = tmpl.Execute(&commandStr, job)
	if err != nil {
		log.Fatalf("render the command template failed: %v", err)
	}
	return commandStr.String(), err
}
