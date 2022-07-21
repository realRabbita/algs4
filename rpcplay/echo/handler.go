package echo

// Handler handles the Echo RPC calls for the remote server.
type Handler struct{}

// Echo must comply with the RPC function signature.
// 方法必须满足Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法
func (h *Handler) Echo(request string, reply *string) error {
	*reply = request
	return nil
}
