package goversioninfo

import (
	"bytes"
	"encoding/binary"
	"os"

	"GoPackager/coff"
	"GoPackager/ico"
)

const (
	rtIcon      = coff.RT_ICON
	rtGroupIcon = coff.RT_GROUP_ICON
	rtManifest  = coff.RT_MANIFEST
)

// on storing icons, see: http://blogs.msdn.com/b/oldnewthing/archive/2012/07/20/10331787.aspx
type gRPICONDIR struct {
	ico.ICONDIR
	Entries []gRPICONDIRENTRY
}

func (group gRPICONDIR) Size() int64 {
	return int64(binary.Size(group.ICONDIR) + len(group.Entries)*binary.Size(group.Entries[0]))
}

type gRPICONDIRENTRY struct {
	ico.IconDirEntryCommon
	ID uint16
}

func addIcon(coff *coff.Coff, fname string, newID <-chan uint16) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	icons, err := ico.DecodeHeaders(f)
	if err != nil {
		return err
	}

	if len(icons) > 0 {
		// RT_ICONs
		group := gRPICONDIR{ICONDIR: ico.ICONDIR{
			Reserved: 0, // magic num.
			Type:     1, // magic num.
			Count:    uint16(len(icons)),
		}}
		gid := <-newID
		for _, icon := range icons {
			id := <-newID
			buff, err := bufferIcon(f, int64(icon.ImageOffset), int(icon.BytesInRes))
			if err != nil {
				return err
			}
			coff.AddResource(rtIcon, id, buff)
			group.Entries = append(group.Entries, gRPICONDIRENTRY{IconDirEntryCommon: icon.IconDirEntryCommon, ID: id})
		}
		coff.AddResource(rtGroupIcon, gid, group)
	}

	return nil
}

func bufferIcon(f *os.File, offset int64, size int) (*bytes.Reader, error) {
	data := make([]byte, size)
	_, err := f.ReadAt(data, offset)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
