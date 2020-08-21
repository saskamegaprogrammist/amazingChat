package utils

var statusCodes = map[string]int{
	"OK" :          200,
	"Created" :     201,
	"Bad Request" : 400,
	"Unauthorized" : 401,
	"Not Found" : 404,
	"Method Not Allowed" : 405,
	"Conflict" : 409,
	"Internal Server Error": 500,
	"Not Implemented": 501,
	"Bad Gateway" : 502,
	"Service Unavailable" : 503,
}

const (
	NO_ERROR = iota
	USER_ERROR
	SERVER_ERROR
)

func StatusCode(mess string) int {
	return statusCodes[mess]
}

const ERROR_ID  = 0
const LogFile  = "log.log"
const DBName = "amazing_chat"
const PortNum = ":9000"

