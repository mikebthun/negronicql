# negronicql


Golang [Negroni](https://github.com/codegangsta/negroni) middleware for the [gocql package](https://github.com/gocql/gocql).

### Dependencies

Depends on [Gorilla Context Package](http://www.gorillatoolkit.org/pkg/context)
`

### Usage
 


#### Backward Compatibility

Nov-16-2014: Changed from Session to CQLSession to avoid Session issues in Negroni by Dal-Papa. This pull will break older builds.


~~~ go
import(

"github.com/mikebthun/negronicql"
"github.com/gorilla/mux"
"github.com/gorilla/context"

)

 router := mux.NewRouter()

 router.HandleFunc("/", MyHandler ).Methods("POST")
 n := negroni.Classic()
 cqldb := negronicql.New()

 //set cluster options here if needed, defaults to localhost
 //cqldb.Ips = []string{"127.0.0.1", "127.0.0.2"}  
 //cqldb.Consistency = gocql.Quorum
 cqldb.Keyspace = "MyKeySpace"

 cqldb.Connect()
 
 //defer close here IMPORTANT
 defer cqldb.Session.Close()

 n.Use(cqldb)
 n.UseHandler(router)
 n.Run(":3000")
 

func MyHandler(w http.ResponseWriter, req *http.Request) {
 
     //grab the session here

     session = context.Get( req, "CQLSession" ).(*gocql.Session)


     
}
~~~

Run your queries like normal on the gocql session:

~~~ go

session.Query( `SELECT * FROM blah` ).Consistency(gocql.One).Exec()

~~~

If you want to customize your ClusterConfig object, you can instantiate one, give it its attributes and Connect().

~~~ go
  cqldb := negronicql.New()
  cqldb.Cluster = gocql.NewCluster("127.0.0.1", "127.0.0.2")
  cqldb.Cluster.Authenticator = gocql.PasswordAuthenticator{"user", "password"}
  cqldb.Cluster.Port = 4242
  cqldb.Cluster.Keyspace = "MyKeySpace"
  cqldb.Cluster.Consistency = gocql.Quorum
  // ...
  cqldb.Connect()

~~~

### Contributors

Author : Mike B Thun @mikebthun

Contrib : Clem DalPalu @Dal-Papa
 
### License 

The MIT License (MIT)
