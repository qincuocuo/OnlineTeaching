package mservices

import (
	"errors"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"time"
)

type ServiceType int

const (
	_ ServiceType = iota
	ServiceProxy
	ServiceDebugger
)

func NewService(optionHandlers ...OptionHandler) Service {
	option := &Option{}
	for _, oh := range optionHandlers {
		oh(option)
	}

	switch option.ServerType {
	case ServiceProxy:
		return &Proxy{Option: option}
	case ServiceDebugger:
		return &Debugger{Option: option}
	default:
		panic("unsupported service type:" + strconv.Itoa(int(option.ServerType)))
	}
}

func initRouter(o *Option) (*http.Server, error) {
	switch o.ServerType {
	case ServiceDebugger:
		return initDebuggerRouter(o)
	case ServiceProxy:
		return initProxyRouter(o)
	default:
		return nil, errors.New("without router")
	}
}

func healthCheckProxyHandler(o *Option) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		v := mux.Vars(request)
		server := v["server"]
		if server == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		host, err := o.register.Query(server)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := http.Head("http://" + host + "/healthcheck")
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			writer.WriteHeader(http.StatusBadGateway)
			return
		}

		_, _ = io.Copy(ioutil.Discard, resp.Body)
		writer.WriteHeader(resp.StatusCode)
	}
}

func healthCheckHandler(o *Option) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !o.hc.HealthCheck() {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}

func initDebuggerRouter(o *Option) (*http.Server, error) {
	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", healthCheckHandler(o)).Methods(http.MethodHead)
	router.HandleFunc("/debug/{type}", func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		t := v["type"]
		switch t {
		case "trace":
			err := trace.Start(w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer trace.Stop()

			time.Sleep(5 * time.Second)
		case "cpu":
			err := pprof.StartCPUProfile(w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer pprof.StopCPUProfile()

			time.Sleep(5 * time.Second)
		case "":
			w.WriteHeader(http.StatusBadRequest)
		default:
			prof := pprof.Lookup(t)
			if prof == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			err := prof.WriteTo(w, 0)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}).Methods(http.MethodGet)

	return &http.Server{Handler: router}, nil
}

func initProxyRouter(o *Option) (*http.Server, error) {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheckHandler(o)).Methods(http.MethodHead)
	router.HandleFunc("/healthcheck/{server}", healthCheckProxyHandler(o)).Methods(http.MethodHead)

	router.Path("/servers/{server}/types/{type}").
		Methods(http.MethodGet).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)

			server := v["server"]
			t := v["type"]
			if t == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			sn, err := o.register.Query(server)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp, err := http.Get("http://" + sn + "/debug/" + t)
			if resp != nil {
				defer resp.Body.Close()
			}
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				return
			}

			w.Header().Set("Content-Disposition", "attachment;filename="+t+".pprof")
			w.Header().Set("Content-Type", "application/octet-stream")
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		})

	return &http.Server{Handler: router}, nil
}
