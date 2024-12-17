package constant

import "github.com/google/uuid"

func GenTxID() string {
	u := uuid.New()
	return u.String()
}
