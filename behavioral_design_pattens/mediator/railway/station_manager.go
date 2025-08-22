package railway

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		trainQueue:     make([]Train, 0),
	}
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	s.isPlatformFree = true
	if len(s.trainQueue) > 0 {
		firstTrain := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrain.PermitArrival()
	}
}
