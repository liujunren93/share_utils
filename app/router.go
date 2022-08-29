package app

func (app *App) initRouter() {
	if !app.LocalConf.EnableAutoRoute {
		return
	}

}
