package main

import (
	restful "github.com/emicklei/go-restful/v3"
)

const (
	apiPrefix        = "/scheduler"
	bindPath         = apiPrefix + "/bind"
	preemptionPath   = apiPrefix + "/preemption"
	predicatesPrefix = apiPrefix + "/predicates"
	prioritiesPrefix = apiPrefix + "/priorities"
)

type Resource struct{}

func (s Resource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path(apiPrefix).
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	//tags := []string{"scheduler"}

	ws.Route(ws.GET("/version").To(s.version))

	return ws
}

func (s Resource) version(req *restful.Request, res *restful.Response) {
	res.WriteEntity(map[string]string{"version": "1.0.0"})
}

func main() {

}
