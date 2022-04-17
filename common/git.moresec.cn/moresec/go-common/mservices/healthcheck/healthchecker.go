package healthcheck

type HealthChecker interface {
	HealthCheck() bool
}

type Default struct {
}

func (d *Default) HealthCheck() bool {
	return true
}
