package db

import (
	"strconv"
	"time"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

type CassandraDB struct {
	Hosts    []string
	Keyspace string
	Session  *gocql.Session
}

func NewCassandra(hosts []string, keyspace string) (*CassandraDB, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &CassandraDB{
		Hosts:    hosts,
		Keyspace: keyspace,
		Session:  session,
	}, nil
}

func (c *CassandraDB) Name() string { return "cassandra" }

func (c *CassandraDB) WriteTest(n int) (time.Duration, error) {
	start := time.Now()
	for i := 0; i < n; i++ {
		if err := c.Session.Query(
			"INSERT INTO tests (id, value) VALUES (?, ?)",
			i, "value-"+strconv.Itoa(i),
		).Exec(); err != nil {
			return 0, err
		}
	}
	return time.Since(start), nil
}

func (c *CassandraDB) ReadTest(n int) (time.Duration, error) {
	start := time.Now()
	for i := 0; i < n; i++ {
		var value string
		err := c.Session.Query("SELECT value FROM tests WHERE id = ?", i).Scan(&value)
		if err != nil && err != gocql.ErrNotFound {
			return 0, err
		}
	}
	return time.Since(start), nil
}
