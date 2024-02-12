package model

type Config struct {
	Clusters       []*Cluster `json:"clusters"`
	Users          []*User    `json:"users"`
	Contexts       []*Context
	CurrentContext string `json:"current-context"`
}

type Context struct {
	Name    string `json:"name"`
	Context struct {
		Cluster   string `json:"cluster"`
		User      string `json:"user"`
		Namespace string `json:"namespace"`
	}
}

type Cluster struct {
	Name string `json:"name"`
}

type User struct {
	Name string `json:"name"`
}
