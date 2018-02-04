package main

import (
	"github.com/golang/glog"
	clientset "github.com/radu-matei/learning-operator/pkg/client/clientset/versioned"
	sscheme "github.com/radu-matei/learning-operator/pkg/client/clientset/versioned/scheme"
	informers "github.com/radu-matei/learning-operator/pkg/client/informers/externalversions"
	listers "github.com/radu-matei/learning-operator/pkg/client/listers/example/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

const controllerAgentName = "simple-example-controller"

// Controller is the controller implementation for Foo resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// exampleclientnet is a clientset for our own API group
	exampleclientset clientset.Interface

	simpleExampleLister listers.SimpleExampleLister
	simpleExampleSynced cache.InformerSynced

	workqueue workqueue.RateLimitingInterface
	recorder  record.EventRecorder
}

// NewController returns a new instance of a controller
func NewController(
	kubeclientset kubernetes.Interface,
	exampleclientset clientset.Interface,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
	sampleInformerFactory informers.SharedInformerFactory) *Controller {

	simpleExampleInformer := sampleInformerFactory.Example().V1alpha1().SimpleExamples()
	sscheme.AddToScheme(scheme.Scheme)

	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:       kubeclientset,
		exampleclientset:    exampleclientset,
		simpleExampleLister: simpleExampleInformer.Lister(),
		simpleExampleSynced: simpleExampleInformer.Informer().HasSynced,
		workqueue:           workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "SimpleExamples"),
		recorder:            recorder,
	}

	glog.Info("Setting up event handlers")
	// Set up an event handler for when SimpleExample resources change
	simpleExampleInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueSimpleExample,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueSimpleExample(new)
		},
	})

	return controller
}

func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {

	return nil
}

// enqueueSimpleExample takes a SimpleExample resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than SimpleExample.
func (c *Controller) enqueueSimpleExample(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}
	c.workqueue.AddRateLimited(key)
}
