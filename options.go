package gofh

// Option is a functional option for configuring an App
type Option func(*App)

// WithAddr sets the address for the server to listen on
func WithAddr(addr string) Option {
	return func(a *App) {
		a.server.SetAddr(addr)
	}
}

// WithTemplateDir sets the directory for HTML templates
func WithTemplateDir(dir string) Option {
	return func(a *App) {
		a.core.SetTemplateDir(dir)
	}
}

// WithStaticDir sets the directory for serving static files
func WithStaticDir(dir string) Option {
	return func(a *App) {
		a.server.SetStaticDir(dir)
	}
}

// WithDebug enables debug mode
func WithDebug(debug bool) Option {
	return func(a *App) {
		// TODO: Implement debug mode in core and server
	}
}

// Additional options can be added here
