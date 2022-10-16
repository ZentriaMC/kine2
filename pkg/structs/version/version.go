package version

// etcd/version/version.go
var (
	ClusterVersionNotDecided = "not_decided"
	Version           = "2.3.8"


	GitSHA = "Not provided (use ./build instead of go build)"
)

type Versions struct {
	Server  string `json:"etcdserver"`
	Cluster string `json:"etcdcluster"`
}
