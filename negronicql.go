package negronicql

import (
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/context"
)

type Negronicql struct {
	Cluster     *gocql.ClusterConfig
	Ips         []string
	Keyspace    string
	Consistency gocql.Consistency
	Session     *gocql.Session
}

func New() *Negronicql {
	return &Negronicql{}
}

// be sure to defer session.close()
func (m *Negronicql) Connect() error {

	//default to localhost
	if len(m.Ips) < 1 {

		m.Ips = []string{"127.0.0.1"}

	}

	//create cluster config
	m.Cluster = gocql.NewCluster(m.Ips[0])

	session, err := m.Cluster.CreateSession()

	m.Session = session

	if err != nil {

		return err

	}

	return nil

}

// The middleware handler
func (m *Negronicql) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//attach the session
	context.Set(r, "Session", m.Session)

	// Call the next middleware handler
	next(rw, r)

}
