package main

func (service *Service) Start() (string, error) {
	go mlbPing()
	return "", nil
}

func (service *Service) Stop() (string, error) {
	return "", nil
}