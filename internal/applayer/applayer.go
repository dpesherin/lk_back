package applayer

import "lk_back/internal/storelayer"

type App struct {
	store storelayer.Store
}

type AppLayer interface {
}

func New(store storelayer.Store) *App {
	return &App{
		store: store,
	}
}
