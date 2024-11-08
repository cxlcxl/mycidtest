package promotion

import (
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
)

type Service struct {
	C         *config.Config
	DbConnect *data.Data
}
