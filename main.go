package main

import (
	"database/sql"
	"fmt"
	log "github.com/parampavar/estimationgame/Godeps/_workspace/src/github.com/cihub/seelog"
	_ "github.com/parampavar/estimationgame/Godeps/_workspace/src/github.com/lib/pq"
	"net/http"
	"os"
//	"reflect"
	"github.com/parampavar/estimationgame/models"
	// "encoding/json"
)
import "gopkg.in/redis.v3"

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "postgres"
    DB_LOCATION = "localhost"
    DB_NAME     = "postgres"
    DB_SSLMODE  = "disable"
)

// const (
// 	DB_USER     = "u311d07be533d42da8c704a4c29f0d573"
// 	DB_PASSWORD = "c9e75db43e744176a5970138c3b7f080"
// 	DB_LOCATION = "10.72.6.110:5432"
// 	DB_NAME     = "d311d07be533d42da8c704a4c29f0d573"
// 	DB_SSLMODE  = "disable" //verify-full"
// )

const (
	DEFAULT_PORT = "9000"
)
// Redis related - Start
var client *redis.Client
func Init() {
	client = redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	client.FlushDb()
}

func RedisGetValue(key string) (interface{}, error) {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
        log.Info("'" + key + "' does not exists")
    } else if err != nil {
        log.Info(err)
    } else {
        log.Info("'" + key + "' exists with val='" + val + "'")
    }
    return val, err
}

func RedisSetValue(key string, val interface{}) error {
	err := client.Set(key, val, 0).Err()
	if err != nil {
        log.Info(err)
    }
    return err
}
// Redis related - End


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("HomeHandler Starting")
	log.Info("HomeHandler url.Path=" + r.URL.Path[1:len(r.URL.Path)])
	fmt.Fprintln(w, "Hello, myWorld!n")
	log.Info("HomeHandler Ending")
}

func DBHandler(w http.ResponseWriter, r *http.Request) {
	tbl := "tbl" + r.URL.Path[1:len(r.URL.Path)] //Get the table name from the url

	log.Info("DBHandler url.Path=" + tbl)

	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", DB_USER, DB_PASSWORD, DB_LOCATION, DB_NAME, DB_SSLMODE)
	log.Info("DB ConnectString = ", dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Info(err)
	}
	log.Info("DB Connection successful")

	err = db.Ping()
	if err != nil {
		log.Info("DBHandler Ping")
		log.Info(err) // log.Critical(err)
	}
	log.Info("DB Ping successful")

	rows, err := db.Query("SELECT *  FROM " + tbl )
	if err != nil {
		log.Info("DBHandler Query")
		log.Info(err) // log.Critical(err)
	}
	log.Info("DB Query successful")

	fmt.Fprintf(w, "Rows of " + tbl)
	log.Info("DB Reading thru the rows")
	retString := ""
	if "tbltoy" == tbl {
		grows, err := models.ScanToys(rows) // ScanUsers was auto-generated!
		if err != nil {
		    log.Info(err)
		}
		retString = models.ToysJson(grows)
	} 
	if "tbltool" == tbl {
		grows, err := models.ScanTools(rows) // ScanUsers was auto-generated!
		if err != nil {
		    log.Info(err)
		}
		retString = models.ToolsJson(grows)
	} 
	if "tbluser" == tbl {
		grows, err := models.ScanUsers(rows) // ScanUsers was auto-generated!
		if err != nil {
		    log.Info(err)
		}
		retString = models.UsersJson(grows)
	} 
	// if "tbltool" == tbl {
	// 	grows, err := models.ScanTools(rows) // ScanUsers was auto-generated!
	// }

	
    log.Info(retString)
	fmt.Fprintf(w, retString)
	w.Write([]byte("Gorilla2\n"))
}


func main() {
	// defer log.Flush()
	log.Info("App Started")

	log.Info("Connecting to redis")

	Init()
	RedisGetValue("abced")
	RedisSetValue("abced", "1111111111111111111111111111111")
	val, _ := RedisGetValue("abced")
	log.Info(val)
	RedisGetValue("abced")
	// if err != nil {
	// 	log.Info("Connecting to redis errored")
	// 	log.Info(err)
	// }
	log.Info("Connecting to redis successful")


	// router := mux.NewRouter()
	// router.HandleFunc("/", HomeHandler)
	// router.HandleFunc("/db", DBHandler)
	// // Bind to a port and pass our router in
	// http.ListenAndServe(":8000", nil)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Info("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/home", HomeHandler)

	http.HandleFunc("/user", DBHandler)
	http.HandleFunc("/toy", DBHandler)
	http.HandleFunc("/tool", DBHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Info("ListenAndServe: ", err)
	}

}
