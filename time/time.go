package time

import "time"

/* CurrentDate proporciona la data actual.
 * El format de la data és: dd-mm-aaaa
 */
func CurrentDate() string {
	return time.Now().Format("02-01-2006")
}

/* CurrentTime proporciona l'hora actual.
 * El format de l'hora és: hh:mm:ss
 */
func CurrentTime() string {
	return time.Now().Format("15:04:05")
}

/* DateTime proporciona la data i hora actuals.
 * El format de la date-time és: dd-mm-aaa hh:mm:ss
 */
func DateTime() string {
	return time.Now().Format("02-01-2006 15:04:05")
}

/* Unix retorna l'hora actual en UnixNano
 */
func Unix() int64 {
	return time.Now().UnixNano()
}
