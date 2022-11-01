package drapper

import "log"

func LogInfo(msg string) {
	log.Printf("INFO - %v", msg)
}
func LogError(msg string) {
	log.Printf("ERROR - %v", msg)
}
func LogWarn(msg string) {
	log.Printf("WARN - %v", msg)
}