package rhdiff

// https://librsync.github.io/delta_8c.html

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"

	"github.com/resin-os/circbuf"
)

func Delta(sig *SignatureType, i io.Reader, output io.Writer) error {
	input := bufio.NewReader(i)

	err := binary.Write(output, binary.BigEndian, DELTA_MAGIC)
	if err != nil {
		return err
	}

	prevByte := byte(0)
	m := match{output: output}

	hash := NewRollingHash()
	block, _ := circbuf.NewBuffer(int64(sig.chunkLen))
	block.WriteByte(0)
	pos := 0

	for {
		pos += 1
		in, err := input.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if block.TotalWritten() > 0 {
			prevByte, err = block.Get(0)
			if err != nil {
				return err
			}
		}
		block.WriteByte(in)
		hash.Rollin(in)

		if hash.count < uint64(sig.chunkLen) {
			continue
		}

		if hash.count > uint64(sig.chunkLen) {
			err := m.add(MATCH_KIND_LITERAL, uint64(prevByte), 1)
			if err != nil {
				return err
			}
			hash.Rollout(prevByte)
		}

		if blockIdx, ok := sig.hash[hash.Digest()]; ok {
			strong2, _ := CalcStrongSum(block.Bytes(), sig.strongLen)
			if bytes.Equal(sig.strongHash[blockIdx], strong2) {
				hash.Reset()
				block.Reset()
				err := m.add(MATCH_KIND_COPY, uint64(blockIdx)*uint64(sig.chunkLen), uint64(sig.chunkLen))
				if err != nil {
					return err
				}
			}
		}
	}

	for _, b := range block.Bytes() {
		err := m.add(MATCH_KIND_LITERAL, uint64(b), 1)
		if err != nil {
			return err
		}
	}

	if err := m.flush(); err != nil {
		return err
	}

	return binary.Write(output, binary.BigEndian, OP_END)
}
