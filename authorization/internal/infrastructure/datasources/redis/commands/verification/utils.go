package verification

import "fmt"

func getKey(accountId int) string {
	return fmt.Sprintf("verification_entities%v", accountId)
}
