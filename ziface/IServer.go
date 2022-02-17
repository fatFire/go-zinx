package ziface

type IServer interface {
	// create new server
	Start()
	// stop server
	Stop()
	// run server
	Serve()
}
