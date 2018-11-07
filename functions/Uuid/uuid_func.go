package uuid_func

import (
	"github.com/satori/go.uuid"
)

func Format(str string) uuid.UUID {
	return uuid.Must(uuid.FromString(str))
}
