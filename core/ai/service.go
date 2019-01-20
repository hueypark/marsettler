package ai

type Service struct {
	interval float64
	time     float64
	fn       func()
}

func (service *Service) Update(delta float64) {
	service.time += delta

	if service.interval <= service.time {
		service.time -= service.interval
		service.fn()
	}
}

func NewService(interval float64, fn func()) *Service {
	service := &Service{
		interval,
		interval,
		fn,
	}

	return service
}
