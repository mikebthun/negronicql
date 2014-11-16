package negronicql_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/Dal-Papa/negronicql"
	"github.com/gocql/gocql"
)

var (
	keyspace     = "mikebthun_negronicql_middleware_tests"
	ips          []string
	columnFamily = "go_db_package_test"
	email        = "test@test.com"
	ip           *string
	session      *gocql.Session
)

func TestIndex(t *testing.T) {

	conn := negronicql.New()
	err := conn.Connect()

	if err != nil {

		t.Fatalf("Is cassandra running?: %s", err)

	}

	session = conn.Session

}

func TestSetup(t *testing.T) {

	ip = flag.String("ip", "127.0.0.1", "Cassandra Ip")
	ips = []string{*ip}
	flag.Parse()

	//create the keyspace if does not exist
	cql := fmt.Sprintf(`

    CREATE KEYSPACE %s WITH REPLICATION = { 
      'class' : 'SimpleStrategy', 
      'replication_factor' : 1 

    }`, keyspace)

	session.Query(cql).Exec()

	cql = fmt.Sprintf("DROP TABLE %s.%s", keyspace, columnFamily)

	session.Query(cql).Exec()

	cql = fmt.Sprintf(`
    
    CREATE TABLE %s.%s 
    ( 
      email text, 
      first text, 
      last text, 
      PRIMARY KEY ( email ) 

    ) 

    `, keyspace, columnFamily)

	if err := session.Query(cql).Exec(); err != nil {

		t.Errorf("%s", err)

	}

}

func TestInsertWithParams(t *testing.T) {

	cql := fmt.Sprintf(`
    
    INSERT INTO %s.%s 
    (email, first, last) 
    VALUES ( ?, ?, ? )

    `, keyspace, columnFamily)

	session.Query(cql, email, "Mike", "B").Exec()

}

func TestSelectWithParams(t *testing.T) {

	cql := fmt.Sprintf(`

    SELECT email
    FROM %s.%s 
    WHERE email = ? 
    LIMIT 1

    `, keyspace, columnFamily)

	var check_email string

	if err := session.Query(cql, email).Consistency(gocql.One).Scan(&check_email); err != nil {

		t.Fatalf("Query failed %s", err.Error())

	}

	if check_email != email {

		t.Errorf("Email should be %s but is %s.", email, check_email)

	}

}

func TestCleanup(t *testing.T) {

	cql := fmt.Sprintf("DROP KEYSPACE %s", keyspace)

	session.Query(cql).Exec()

	defer session.Close()

}
