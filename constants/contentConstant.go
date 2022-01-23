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
	INVALID_CSV_REQUEST          = "invalid csv form body"
	INVALID_CONTENT_ID           = "invalid contentId"
	ERROR_IN_HASHING_PASSWORD    = "error while hashing password"
	ERROR_IN_STORING_UNIQUE_USER = "the user already exists"
	ERROR_IN_AUTHENTICATING_USER = "the password is incorrect"
	USER_ID_NOT_ARRAY_OF_INT     = "userIds must be an array of integers"
	ERROR_NO_USER_EXIST          = "User does not exist"
	FILE_SIZE_EXCEEDED           = "file size exceeded"
	CONTENT_DOES_NOT_BELONG_TO_USER = "content does not belong to user"
)

const (
	MAX_FILE_SIZE = 10 * 1_024 * 1_024
)

const (
	API_TYPE_CREATE_CONTENT = "createContent"
	API_TYPE_EDIT_CONTENT   = "editContent"
	DELETE_CONTENT_MESSAGE = "deleteContent"
	FETCH_CONTENT_MESSAGE  = "fetchContent"
)

// response messages
const (
	UPLOAD_CSV_SUCCESS_MESSAGE = "CSV File Uploaded Successfully"
	CREATE_CONTENT_MESSAGE     = "Content Created Successfully"
	EDIT_CONTENT_MESSAGE       = "Content Edited Successfully"
)

var (
	CREATE_CONTENT_REQUIRED_FIELDS = []string{"title", "story", "userId"}
	CREATE_CONTENT_OPTIONAL_FIELDS = []string{}

	EDIT_CONTENT_REQUIRED_FIELDS = []string{"title", "story", "userId"}
	EDIT_CONTENT_OPTIONAL_FIELDS = []string{}
)

const (
	MIN_LENGTH_OF_CONTENT_TITLE = 5
	MAX_LENGTH_OF_CONTENT_TITLE = 100
)
