package pkg

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0) // #nosec G404
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id.String()
}
