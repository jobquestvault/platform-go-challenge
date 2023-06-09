package pg

import "github.com/jobquestvault/platform-go-challenge/internal/sys/errors"

var (
	NoConnectionErr      = errors.NewError("no connection error")
	UnsupportedAssetErr  = errors.NewError("unsupported asset type")
	NoRecordsAffectedErr = errors.NewError("no records affected")
)
