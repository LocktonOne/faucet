package handlers

import (
	"net/http"
	"net/url"
)

type RPCProxy struct {
	backend func(r *http.Request) *url.URL
	sender  string
	amount  uint32
}

func NewProxy(target *url.URL, sender string, amount uint32) *RPCProxy {
	backend := func(r *http.Request) *url.URL {
		u := *target
		u.Fragment = r.URL.Fragment
		u.Path = r.URL.Path
		u.RawQuery = r.URL.RawQuery
		return &u
	}

	return &RPCProxy{backend: backend, sender: sender, amount: amount}
}

func (rpc *RPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//request, err := requests.NewRequest(r)
	//if err != nil {
	//	w.Write([]byte(helpers.InvalidRequestError))
	//	return
	//}
	//
	//err = filters.RunFilter(request, rpc.filters, rpc.filtersParams)
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//
	//msg, err := json.Marshal(request)
	//if err != nil {
	//	w.Write([]byte(helpers.ParseError))
	//	return
	//}
	//
	//response, err := http.Post(rpc.Eth_rpc, "application/json", bytes.NewBuffer(msg))
	//defer response.Body.Close()
	//if err != nil {
	//	w.Write([]byte(helpers.InternalError))
	//	return
	//}
	//
	//responseMess, err := io.ReadAll(response.Body)
	//if err != nil {
	//	w.Write([]byte(helpers.ParseError))
	//	return
	//}

	//w.Write(responseMess)

}
