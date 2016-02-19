package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	log "github.com/parampavar/estimationgame/Godeps/_workspace/src/github.com/cihub/seelog"
	_ "github.com/parampavar/estimationgame/Godeps/_workspace/src/github.com/lib/pq"
	"github.com/parampavar/estimationgame/Godeps/_workspace/src/gopkg.in/redis.v3"
	"github.com/parampavar/estimationgame/models"
)

//import "github.com/parampavar/estimationgame/Godeps/_workspace/src/gopkg.in/redis.v3"
import "strconv"

const (
	DB_USER     = "u311d07be533d42da8c704a4c29f0d573"
	DB_PASSWORD = "c9e75db43e744176a5970138c3b7f080"
	DB_LOCATION = "10.72.6.110:5432"
	DB_NAME     = "d311d07be533d42da8c704a4c29f0d573"
	DB_SSLMODE  = "disable" //verify-full"
)

const (
	DEFAULT_PORT = "9000"
)

var isInCloud = false

// redisConnectAddrt
var client *redis.Client
var pgDbConnectionString = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", "postgres", "postgres", "localhost", "postgres", "disable")
var redisConnectAddr = fmt.Sprintf("%s:%s", "localhost", "6379")
var redisConnectPassword = ""
var redisConnectPort float64 = 6379

func FloatToString(input_num float64, number_of_decimals int) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', number_of_decimals, 64)
}
func RedisInit() {
	// client = redis.NewClient(&redis.Options{
	//   Addr: redisConnectAddr,
	// })
	log.Info("RedisInit Started...")
	log.Info("RedisInit connecting to " + redisConnectAddr + ",  password: " + redisConnectPassword)
	client := redis.NewClient(&redis.Options{
		Addr:     redisConnectAddr,
		Password: redisConnectPassword,
		DB:       0, // use default DB
	})

	client.FlushDb()
}

func RedisGetValue(key string) (interface{}, error) {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		log.Info("'" + key + "' does not exists")
		val = ""
	} else if err != nil {
		log.Info(err)
		val = ""
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
	fmt.Fprintln(w, "Hello, myWorld!n")
	log.Info("HomeHandler Ending")
}

func getDBData(tbl string) (interface{}, error) {
	var rows interface{}
	var err error
	var db *sql.DB
	var getFromDB = false

	if isInCloud == true {
		rows, err = RedisGetValue(tbl)
		if err == redis.Nil {
			log.Info("getDBData: Data not in Redis cache. Getting from DB.")
			getFromDB = true
		}
	} else {
		getFromDB = true
	}
	if getFromDB == true {
		log.Info("getDBData: DB ConnectString = ", pgDbConnectionString)
		db, err = sql.Open("postgres", pgDbConnectionString)
		if err != nil {
			log.Info(err)
		}
		log.Info("getDBData: DB Connection successful")

		err = db.Ping()
		if err != nil {
			log.Info("getDBData: Ping error")
			log.Info(err) // log.Critical(err)
		}
		log.Info("getDBData: DB Ping successful")

		rows, err = db.Query("SELECT *  FROM " + tbl)
		if err != nil {
			log.Info("getDBData: Query error")
			log.Info(err) // log.Critical(err)
		}
		log.Info("getDBData: DB Query successful")
		if isInCloud == true {
			log.Info("getDBData: Storing in Redis")
			RedisSetValue(tbl, rows)
		}
		// fmt.Fprintf(w, "Rows of "+tbl)
		// log.Info("DB Reading thru the rows")
	}
	return rows, err
}
func DBHandler(w http.ResponseWriter, r *http.Request) {
	tbl := "tbl" + r.URL.Path[1:len(r.URL.Path)] //Get the table name from the url

	log.Info("DBHandler url.Path=" + tbl)
	genrows, err := getDBData(tbl)
	if err != nil {
		log.Info("DBHandler getDBData error")
		log.Info(err)
	}
	retString := ""

	if rows, ok := genrows.(*sql.Rows); ok {
		log.Info("DBHandler genrows is sql.Rows")
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
	} else {
		/* not sql.Rows */
		log.Info("DBHandler genrows is NOT sql.Rows")
		log.Info(reflect.TypeOf(genrows))
		log.Info("DBHandler genrows is NOT sql.Rows")
	}

	// if "tbltool" == tbl {
	//  grows, err := models.ScanTools(rows) // ScanUsers was auto-generated!
	// }

	log.Info(retString)
	fmt.Fprintf(w, retString)
	w.Write([]byte("Gorilla2\n"))
}

func main() {
	defer log.Flush()
	log.Info("App Started")

	var appUrl = "http://localhost"
	var port = ""

	vcapApplication := os.Getenv("VCAP_APPLICATION")
	if vcapApplication == "" {
		log.Info("App Running locally.......")
		if port = os.Getenv("PORT"); len(port) == 0 {
			log.Info("Warning: PORT not set. Defaulting to ", DEFAULT_PORT)
			port = DEFAULT_PORT
		}
		appUrl = appUrl + ":" + port
	} else {
		log.Info("App Running in the cloud.......")

		var vcapApplicationJson map[string]interface{}
		_ = json.Unmarshal([]byte(vcapApplication), &vcapApplicationJson)
		appUrls := vcapApplicationJson["application_uris"].([]interface{})
		appUrl = "http://" + appUrls[0].(string)
		log.Info("application_uris=" + appUrl)
		pgDbConnectionString = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", DB_USER, DB_PASSWORD, DB_LOCATION, DB_NAME, DB_SSLMODE)
		isInCloud = true
		if port = os.Getenv("PORT"); len(port) == 0 {
			log.Info("Warning: PORT not set. Defaulting to ", DEFAULT_PORT)
			port = DEFAULT_PORT
		}
		appUrl = appUrl + ":" + port

		log.Info("Getting  VCAP_SERVICES")
		vcapServices := os.Getenv("VCAP_SERVICES")
		if vcapServices == "" {
			log.Info("No services bound to this application")
		} else {
			var vcapServicesJson map[string]interface{}
			_ = json.Unmarshal([]byte(vcapServices), &vcapServicesJson)
			cloudRedis := vcapServicesJson["redis-2"].([]interface{})
			log.Info(cloudRedis)
			cloudRedisMap := cloudRedis[0].(map[string]interface{})
			cloudRedisCredentialsMap := cloudRedisMap["credentials"].(map[string]interface{})
			log.Info(cloudRedisCredentialsMap["host"].(string))
			redisConnectAddr = cloudRedisCredentialsMap["host"].(string)
			log.Info(redisConnectAddr)
			redisConnectPassword = cloudRedisCredentialsMap["password"].(string)
			log.Info(redisConnectPassword)
			redisConnectPort = cloudRedisCredentialsMap["port"].(float64)
			log.Info(redisConnectPort)
			redisConnectAddr = redisConnectAddr + ":" + FloatToString(redisConnectPort, 0)
			log.Info(redisConnectAddr)
			log.Info("Connecting to redis")
			RedisInit()
			//RedisGetValue("abced")
			RedisSetValue("abced", "1111111111111111111111111111111")

			// val, _ := RedisGetValue("abced")
			// log.Info(val)
			// RedisGetValue("abced")
			// log.Info("Connecting to redis successful")
		}

	}

	// log.Info("Connecting to redis")

	// Init()
	// RedisGetValue("abced")
	// RedisSetValue("abced", "1111111111111111111111111111111")
	// val, _ := RedisGetValue("abced")
	// log.Info(val)
	// RedisGetValue("abced")
	// // if err != nil {
	// //   log.Info("Connecting to redis errored")
	// //   log.Info(err)
	// // }
	// log.Info("Connecting to redis successful")

	// router := mux.NewRouter()
	// router.HandleFunc("/", HomeHandler)
	// router.HandleFunc("/db", DBHandler)
	// // Bind to a port and pass our router in
	// http.ListenAndServe(":8000", nil)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/home", HomeHandler)

	http.HandleFunc("/user", DBHandler)
	http.HandleFunc("/toy", DBHandler)
	http.HandleFunc("/tool", DBHandler)

	log.Info("App Started at " + appUrl)
	//err := http.ListenAndServe("", nil)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Info("ListenAndServe: ", err)
	}

}
