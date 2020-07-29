package observer

import "sync"

type subject struct {
	observers sync.Map
}

func (s *subject) Register(obs Observer) {
	s.observers.Store(obs, nil)
}

func (s *subject) Unregister(obs Observer) {
	s.observers.Delete(obs)
}

func (s *subject) Broadcast(msg interface{}) {
	wg := &sync.WaitGroup{}
	s.observers.Range(func(obs interface{}, _ interface{}) bool {
		wg.Add(1) // each iteration triggers a new goroutine

		go func(wg *sync.WaitGroup, obs Observer) {
			defer wg.Done()
			obs.OnUpdate(msg)
		}(wg, obs.(Observer))

		return true
	})

	// Waiting for all observers to end its job
	wg.Wait()
}
