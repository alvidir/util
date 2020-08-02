package observer

import "sync"

type subject struct {
	sync.Map
}

func (sbj *subject) Register(obs Observer) {
	sbj.Store(obs, nil)
}

func (sbj *subject) Unregister(obs Observer) {
	sbj.Delete(obs)
}

func (sbj *subject) Broadcast(msg interface{}) {
	wg := &sync.WaitGroup{}
	sbj.Range(func(obs interface{}, _ interface{}) bool {
		wg.Add(1) // each iteration triggers a new goroutine

		go func(wg *sync.WaitGroup, obs Observer) {
			defer wg.Done()
			obs.OnNotification(msg)
		}(wg, obs.(Observer))

		return true
	})

	// Waiting for all observers to end its job
	wg.Wait()
}
