package rootmulti

import (
	"errors"
	"fmt"
	"math"

	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	protoio "github.com/cosmos/gogoproto/io"

	"github.com/crypto-org-chain/cronos/memiavl"
)

// Implements interface Snapshotter
func (rs *Store) Snapshot(height uint64, protoWriter protoio.Writer) (returnErr error) {
	if height > math.MaxUint32 {
		return fmt.Errorf("height overflows uint32: %d", height)
	}
	version := uint32(height)

	exporter, err := memiavl.NewMultiTreeExporter(rs.dir, version, rs.supportExportNonSnapshotVersion)
	if err != nil {
		return err
	}

	defer func() {
		returnErr = errors.Join(returnErr, exporter.Close())
	}()

	for {
		item, err := exporter.Next()
		if err != nil {
			if err == memiavl.ErrorExportDone {
				break
			}

			return err
		}

		switch item := item.(type) {
		case *memiavl.ExportNode:
			if err := protoWriter.WriteMsg(&snapshottypes.SnapshotItem{
				Item: &snapshottypes.SnapshotItem_IAVL{
					IAVL: &snapshottypes.SnapshotIAVLItem{
						Key:     item.Key,
						Value:   item.Value,
						Height:  int32(item.Height),
						Version: item.Version,
					},
				},
			}); err != nil {
				return err
			}
		case string:
			if err := protoWriter.WriteMsg(&snapshottypes.SnapshotItem{
				Item: &snapshottypes.SnapshotItem_Store{
					Store: &snapshottypes.SnapshotStoreItem{
						Name: item,
					},
				},
			}); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown item type %T", item)
		}
	}

	return nil
}
