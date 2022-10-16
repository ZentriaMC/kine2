package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZentriaMC/kine2/pkg/structs/version"
)

var versionStruct = version.Versions{
	Server:  version.Version,
	Cluster: version.ClusterVersionNotDecided,
}

type VersionHandler struct {
}

func (_ *VersionHandler) ServeVersion(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, &versionStruct)
}

func (vh *VersionHandler) Register(mux *mux.Router) {
	mux.HandleFunc("/version", vh.ServeVersion)
}
