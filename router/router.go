package trouter

type TRouterI interface {
	TRoutes
	GET(string, ...HandlerFunc) TRoutesI
	POST(string, ...HandlerFunc) TRoutesI
	DELETE(string, ...HandlerFunc) TRoutesI
	PUT(string, ...HandlerFunc) TRoutesI
}

type TRouter struct {
}

type TRoutesI interface {
}

type TRoutes struct {
}
