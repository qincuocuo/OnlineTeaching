package mservices

import (
	"context"
	"encoding/json"
	"git.moresec.cn/moresec/go-common/mlog"
	"go.uber.org/zap"
	"net/http"
	"net/http/pprof"
)

// ProfilingService is a service that returns profiling information via a
// HTTP server.

var (
	// defaultServeMux ...
	defaultServeMux = http.NewServeMux()
	routes          = []string{}
)

func init() {
	HandleFunc("/debug/pprof/", pprof.Index)
	HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	HandleFunc("/debug/pprof/profile", pprof.Profile)
	HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	HandleFunc("/debug/pprof/trace", pprof.Trace)
	HandleFunc("/routes", func(resp http.ResponseWriter, req *http.Request){
		json.NewEncoder(resp).Encode(routes)
	})
}

type ProfilingService struct {
	server *http.Server
	address string
}

// NewProfilingService creates a new ProfilingService instance bound to
// a specified address.
func NewProfilingService(address string) *ProfilingService {
	return &ProfilingService{
		address: address,
	}
}

// Name returns a human-readable name for a ProfilingService.
func (ps *ProfilingService) Name() string {
	return "Profiling HTTP endpoint"
}

// Serve runs a ProfilingService. It sets up the HTTP endpoint and services
// requests until the service is stopped. It runs on the calling Goroutine.
func (ps *ProfilingService) Serve() error {
	mlog.Info("Serving profiling HTTP endpoints on",
		zap.String("addr", ps.address))

	ps.server = &http.Server{
		Addr: ps.address,
		Handler: defaultServeMux,
	}

	err := ps.server.ListenAndServe()
	if err != nil {
		mlog.Error("Profiling HTTP error", zap.Error(err))
	}

	return err
}

// Stop stops a running ProfilingService.
func (ps *ProfilingService) Stop() {
	ps.server.Shutdown(context.Background())
}

// HandleFunc ...
func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// todo: 增加安全管控
	defaultServeMux.HandleFunc(pattern, handler)
	routes = append(routes, pattern)
}
