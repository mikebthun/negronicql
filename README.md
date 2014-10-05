# negronicql


Golang [Negroni](https://github.com/codegangsta/negroni) middleware for the [gocql package](https://github.com/gocql/gocql).

### Dependencies

Depends on [Gorilla Context Package](http://www.gorillatoolkit.org/pkg/context)
`

### Usage
 
```
import(

"github.com/mikebthun/negronicql"
"github.com/gorilla/mux"
"github.com/gorilla/context"

)

```


```
 
 router := mux.NewRouter()

 router.HandleFunc("/", MyHandler ).Methods("POST")
 n := negroni.Classic()
 cqldb := negronicql.NewNegronicql()

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

     session = context.Get( req, "Session" ).(*gocql.Session)


     
}
```

Run your queries like normal on the gocql session:

```

session.Query( `SELECT * FROM blah` ).Exec()

```
 
### License 

The MIT License (MIT)
