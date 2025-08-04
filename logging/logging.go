package logging

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel tanımları
// Log seviyeleri: ERROR, WARN, INFO, DEBUG
// INFO varsayılan seviye olarak ayarlanır
// DEBUG seviyesi daha ayrıntılı loglar için kullanılabilir
type LogLevel int

const (
	ERROR LogLevel = iota // 0
	WARN                  // 1
	INFO                  // 2
	DEBUG                 // 3
)

var LogLevelNames = map[LogLevel]string{
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
}

var currentLogLevel LogLevel

func SetLogLevel() {
	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
		currentLogLevel = INFO
		return
	}

	switch lvl {
	case "ERROR":
		currentLogLevel = ERROR
	case "WARN":
		currentLogLevel = WARN
	case "INFO":
		currentLogLevel = INFO
	case "DEBUG":
		currentLogLevel = DEBUG
	default:
		currentLogLevel = INFO // Default to INFO if invalid level
	}
}

// Genel uygulama logları için logger
var AppLogger *log.Logger

// Audit (kim, ne yaptı) logları için logger
var AuditLogger *log.Logger

func init() {
	SetLogLevel() // Log seviyesini ayarla

	// Uygulama log dosyası
	appLogFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Uygulama log dosyası açılamadı: %v", err)
	}
	AppLogger = log.New(appLogFile, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Audit log dosyası
	auditLogFile, err := os.OpenFile("audit.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Audit log dosyası açılamadı: %v", err)
	}
	AuditLogger = log.New(auditLogFile, "AUDIT: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogApp genel uygulama logları için fonksiyon
// Log seviyesine göre log yazılır
// Örnek: LogApp(INFO, "Uygulama başlatıldı")
func LogApp(level LogLevel, format string, v ...interface{}) {
	if level <= currentLogLevel { // Sadece geçerli log seviyesindeki loglar yazılır
		AppLogger.Printf("[%s] %s", LogLevelNames[level], fmt.Sprintf(format, v...))
	}
}

// Audit log için struct
type AuditLog struct {
	Timestamp  string      `json:"timestamp"`
	User       string      `json:"user"`
	Endpoint   string      `json:"endpoint"`
	Method     string      `json:"method"`
	StatusCode int         `json:"status_code"`
	ClientIP   string      `json:"client_ip"`
	Params     interface{} `json:"params,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func LogAudit(user, endpoint, method string, statusCode int, clientIP string, params interface{}, message string) {
	auditLog := AuditLog{
		Timestamp:  time.Now().Format(time.RFC3339),
		User:       user,
		Endpoint:   endpoint,
		Method:     method,
		StatusCode: statusCode,
		ClientIP:   clientIP,
		Params:     params,
		Message:    message,
	}

	jsonEntry, err := json.Marshal(auditLog)
	if err != nil {
		AppLogger.Printf("Audit log oluşturulamadı: %v", err)
		return
	}

	AuditLogger.Println(string(jsonEntry))
}
