package routine

import "github.com/PatchouliG/wisckey-db/lsm"

// get,delete,put ReadResponse
// return by work routine
type ReadResponse struct {
}

type WriteResponse struct {
	//	return new lsm if build new sstable
	newLsm *lsm.Lsm
}