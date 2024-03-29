package common

import "regexp"

const regexID = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile(regexID)
	return r.MatchString(uuid)
}
