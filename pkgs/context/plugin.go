package context

type Plugin interface {
	Signature() any
	PreloadConfig(appCtx *Application) error
	ConfigLoaded(appCtx *Application) error
	PreRun(appCtx *Application) error
}
