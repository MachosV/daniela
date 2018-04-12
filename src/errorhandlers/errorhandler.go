package errorhandlers

import "log"

/*LogError does
what it says
Logs an error in the format i like
*/
func LogError(err error, message string) {
	log.Println(err)
	log.Println(message)
}
