package trouter

type RouterGroup struct {
	Handler  HandlersChain
	basePath string
	router   *TRouter
}


