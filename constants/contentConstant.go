package constants

import "time"

// request related
const (
	HTTP_METHOD_GET  = "GET"
	HTTP_METHOD_POST = "POST"
)

// api status
const (
	API_FAILED_STATUS  = "Fail"
	API_SUCCESS_STATUS = "Success"
)

// DB constants
const (
	DB_READER = "reader"
	DB_WRITER = "writer"
)

// caching constants durations are in nano sec
var (
	CACHE_TTL_VERY_SHORT time.Duration = 60 * 1_000_000_000
	CACHE_TTL_SHORT      time.Duration = 300 * 1_000_000_000
	CACHE_TTL_MEDIUM     time.Duration = 1_800 * 1_000_000_000
	CACHE_TTL_LONG       time.Duration = 3_600 * 1_000_000_000
	CACHE_TTL_VERY_LONG  time.Duration = 86_400 * 1_000_000_000
)

// info messages
const (
	INFO_CACHE_DISABLED = "cache disabled"
)

// code alphabets
const (
	CODE_ALPHABET_SHORT = "abcdef0123456789"
	CODE_ALPHABET_LONG  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// error messages
const (
	INVALID_PASSWORD             = "incorrect password"
	INVALID_USER_ID              = "invalid userId"
	ERROR_USER_IDS_NOT_FOUND     = "userIds not found"
	INVALID_REQUEST              = "invalid json request body"
	INVALID_MAIL_ID              = "invalid email address provided"
	ERROR_IN_HASHING_PASSWORD    = "error while hashing password"
	ERROR_IN_STORING_UNIQUE_USER = "the user already exists"
	ERROR_IN_AUTHENTICATING_USER = "the password is incorrect"
	USER_ID_NOT_ARRAY_OF_INT     = "userIds must be an array of integers"
	ERROR_NO_USER_EXIST          = "User does not exist"
)

const (
	MIN_LENGTH_OF_FIRSTNAME = 2
	MAX_LENGTH_OF_FIRSTNAME = 50
	MIN_LENGTH_OF_LASTNAME  = 2
	MAX_LENGTH_OF_LASTNAME  = 50
	MIN_LENGTH_OF_PASSWORD  = 6
	MAX_LENGTH_OF_PASSWORD  = 100
)

// response messages
const (
	CREATE_USER_MESSAGE = "User Created Successfully"
	EDIT_USER_MESSAGE   = "Name Updated Successfully"
	DELETE_USER_MESSAGE = "User Deleted Successfully"
	FETCH_USER_MESSAGE  = "User Fetched Successfully"
)

// required request body fields
var (
	USER_REGISTER_REQUIRED_FIELDS = []string{"firstName", "lastName", "password", "email", "mobileNo"}
	USER_REGISTER_OPTIONAL_FIELDS = []string{"gender"}

	USER_EDIT_REQUIRED_FIELDS = []string{"firstName", "lastName"}
	USER_EDIT_OPTIONAL_FIELDS = []string{}

	USER_DELETE_REQUIRED_FIELDS = []string{"password"}
	USER_DELETE_OPTIONAL_FIELDS = []string{}
)

const (
	API_TYPE_CREATE_USER = "createUser"
	API_TYPE_EDIT_USER   = "editUser"
	API_TYPE_DELETE_USER = "deleteUser"
)
