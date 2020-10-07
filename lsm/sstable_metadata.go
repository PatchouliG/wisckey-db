package lsm

import (
	"github.com/PatchouliG/wisckey-db/record"
	"path"
	"sort"
)

// own by lsm
type sstableMetaData struct {
	id       sstableId
	startKey record.Key
	endKey   record.Key
	refCount int
}

func newSStableMetaData(id sstableId, startKey record.Key, endKey record.Key) sstableMetaData {
	return sstableMetaData{id, startKey, endKey, 1}
}

// to file name todo
func (sstm *sstableMetaData) fileName(id sstableId) string {
	return path.Join(dbDir, "sstable_"+id.Id.String())
}

// remove  sstable file if ref count is 0
func (sstm *sstableMetaData) decRef() {
	sstm.refCount--
	if sstm.refCount == 0 {
		sstm.deleteSStable()
	}
}

// todo use for gc
func (sstm *sstableMetaData) deleteSStable() {
	panic("")
}

func sortMetadata(sstms []sstableMetaData) {
	sort.Slice(sstms, func(i, j int) bool {
		return sstms[i].startKey.Value() < sstms[j].startKey.Value()
	})
}