package mservices

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"

	"git.moresec.cn/moresec/go-common/mlog"
)

// ServiceManagerOption represents options passed into the service manager constructor
type ServiceManagerOption func(*serviceManagerOptions)

// WithAllOrNothing specifies whether all services should exit upon a single service failure
func WithAllOrNothing() ServiceManagerOption {
	return func(o *serviceManagerOptions) {
		o.allOrNothing = true
	}
}

type serviceManagerOptions struct {
	allOrNothing bool
}

// Service is a service that is registered with and run by a ServiceManager
type Service interface {
	Name() string
	Serve() error
	Stop()
}

// ServiceManager controls an arbitrary number of running services.
type ServiceManager struct {
	sync.Mutex

	options *serviceManagerOptions

	allOrNothing bool

	services []Service
	stopped  bool

	// Channel that signals stopping the server
	stopChan chan struct{}
}

// NewServiceManager creates a new ServiceManager instance.
func NewServiceManager(options ...ServiceManagerOption) *ServiceManager {
	opts := &serviceManagerOptions{}
	for _, option := range options {
		option(opts)
	}
	return &ServiceManager{
		options: opts,
	}
}

// RegisterService registers a Service for management with a ServiceManager.
func (sm *ServiceManager) RegisterService(service Service) {
	sm.Lock()
	sm.services = append(sm.services, service)
	sm.Unlock()
}

// Run starts and runs all services registered with a ServiceManager. This
// function does not return until the ServiceManager is stopped via Stop.
func (sm *ServiceManager) Run() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		mlog.Info("Caught signal stopping services", zap.Reflect("singal", sig))
		sm.Stop()
	}()

	sm.Lock()
	sm.stopChan = make(chan struct{})
	services := sm.services
	sm.Unlock()

	wg := sync.WaitGroup{}

	go func() {
		<-sm.stopChan

		for _, service := range services {
			mlog.Info("Stopping service ", zap.String("name", service.Name()))
			service.Stop()
		}
	}()

	mlog.Info("Starting services ...")
	for _, service := range services {
		wg.Add(1)

		go func(service Service) {
			mlog.Info("Starting service ", zap.String("name", service.Name()))
			err := service.Serve()
			if sm.options.allOrNothing {
				sm.Stop()
			}
			mlog.Info("Serve()", zap.String("name", service.Name()), zap.Error(err))
			wg.Done()
		}(service)
	}

	// Block until goroutines have exited
	mlog.Info("Server is ready")
	wg.Wait()
}

// Stop stops a running ServiceManager. If the ServiceManager is not running,
// this function does nothing.
func (sm *ServiceManager) Stop() {
	sm.Lock()
	defer sm.Unlock()

	if !sm.stopped {
		close(sm.stopChan)

		sm.stopped = true
	}
}
