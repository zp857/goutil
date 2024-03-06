package v1

// Logger Levels
const (
	DebugLevel  = "debug"
	InfoLevel   = "info"
	WarnLevel   = "warn"
	ErrorLevel  = "error"
	SlientLevel = "slient"
)

// Logger Keys
const (
	TimeKey  = "time"
	StackKey = "stack"
	ErrorKey = "error"

	RequestKey = "request"
	TopicKey   = "topicName"
)

// Logger Names
const (
	GinLogger  = "[gin]"
	GormLogger = "[gorm]"

	HTTPClientLogger = "[http-client]"

	KafkaProducerLogger = "[kafka-producer]"
	KafkaCustomerLogger = "[kafka-consumer]"
)

// Debug
const (
	Request  = "request => [%v]\n%v"
	Response = "response =>\n%v"

	ClientRequest  = "request => [%v]\n%v"
	ClientResponse = "response => [%v]\n%v"

	ProducerRequest = "request => [%v] partition=%v offset=%v msg.bytes=%v\n%v"
)

// Info
const (
	HTTPRouter = "%v %v --> %v"
	HTTPServe  = "http server will listening at %v"

	ConsumerRouter = "%v --> %v"
	ConsumerServe  = "kafka consumer will listening at %v"

	InitDataExist   = "the initial data for the table %v already exists"
	InitDataError   = "initialize table(%v) data failed: %v"
	InitDataSuccess = "initialize table(%v) data success"

	SignalExit = "signal received: %v, program exit\n\n"

	InitEngineSuccess = "init engine success =>\n%v"
)

// Error
const (
	HTTPServeError = "http server ListenAndServe() err: %v"

	RequestError = "request err: %v"

	ConsumerInitError     = "consumer init err: %v"
	ConsumerCloseError    = "consumer close err: %v"
	GetPartitionsError    = "get partitions err: %v"
	ConsumePartitionError = "consume partition [%v] err: %v"

	RecoverCommon      = "[recover from panic] error => %v\n"
	RecoverWithStack   = RecoverCommon + "stack =>\n%v\n"
	RecoverWithRequest = RecoverCommon + "request =>\n%v\n"
	RecoverWithAll     = RecoverCommon + "request =>\n%v\nstack =>\n%v\n"

	EmptyDBNameError     = "dbName can't be empty"
	UnSupportDBTypeError = "dbType only support mysql && pgsql"

	ParseConfigError = "error parse config %v err: %v\n"
	InitLoggerError  = "error Init logger err: %v\n"

	ConnDBError       = "conn to the db err: %v"
	InitEngineError   = "init engine err: %v"
	ParseRequestError = "parse request err: %v"
	InvalidOperation  = "invalid operation: %v"
)
