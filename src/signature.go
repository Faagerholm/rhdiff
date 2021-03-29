package rhdiff

import (
	"encoding/binary"
	"fmt"
	"io"

	"golang.org/x/crypto/md4"
)

type MagicNumber uint32

const (
	/** A delta file.
	 *
	 * At present, there's only one delta format.
	 *
	 * The four-byte literal \c "rs\x026". */
	DELTA_MAGIC MagicNumber = 0x72730236

	/** A signature file with RabinKarp rollsum and BLAKE2 hash.
	*
	* Uses a faster/safer rollsum together with the safer BLAKE2 hash. This is
	* the recommended default supported since librsync 2.2.0.
	*
	* The four-byte literal \c "rs\x01G".
	*
	* \sa rs_sig_begin() */
	RS_RK_BLAKE2_SIG_MAGIC MagicNumber = 0x72730147
)

const (
	RK_BLAKE2_LENGTH = 32
)

type SignatureType struct {
	chunkLen   uint32
	strongLen  uint32
	strongHash [][]byte
	hash       map[uint32]int
}

func Signature(input io.Reader, output io.Writer, chunkLen, strongLen uint32) (*SignatureType, error) {

	if strongLen > RK_BLAKE2_LENGTH {
		return nil, fmt.Errorf("strongLen exceeded limit of 32 (%d)", strongLen)
	}

	err := binary.Write(output, binary.BigEndian, RS_RK_BLAKE2_SIG_MAGIC)
	if err != nil {
		return nil, err
	}
	err = binary.Write(output, binary.BigEndian, chunkLen)
	if err != nil {
		return nil, err
	}
	block := make([]byte, chunkLen)

	var res SignatureType
	res.hash = make(map[uint32]int)
	res.chunkLen = chunkLen
	res.strongLen = strongLen

	for {
		raw, err := input.Read(block)
		if err == io.ErrUnexpectedEOF || err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		data := block[:raw]

		// hash := CheckWeakSum(data)
		// TODO: Get hash value from data
		hash := uint32(0)
		err = binary.Write(output, binary.BigEndian, hash)
		if err != nil {
			return nil, err
		}
		strong, _ := CalcStrongSum(data, strongLen)
		output.Write(strong)
		res.hash[hash] = len(res.strongHash)
		res.strongHash = append(res.strongHash, strong)
	}
	return &res, nil
}

func CalcStrongSum(data []byte, strongLen uint32) ([]byte, error) {
	d := md4.New()
	d.Write(data)
	return d.Sum(nil)[:strongLen], nil
}

func LoadSigFile(signature io.Reader, strongLen uint32) (*SignatureType, error) {
	var magic MagicNumber

	err := binary.Read(signature, binary.BigEndian, &magic)
	if err != nil {
		return nil, fmt.Errorf("Unable to read signature file")
	}

	if magic != RS_RK_BLAKE2_SIG_MAGIC {
		return nil, fmt.Errorf("Got magic number %x, expected value %x", magic, RS_RK_BLAKE2_SIG_MAGIC)
	}

	var chunkLen uint32
	err = binary.Read(signature, binary.BigEndian, &chunkLen)
	if err != nil {
		return nil, fmt.Errorf("Unable to determin chunk length")
	}

	var sig SignatureType
	sig.hash = make(map[uint32]int)
	sig.strongLen = strongLen
	sig.chunkLen = chunkLen

	for {
		var hash uint32
		err = binary.Read(signature, binary.BigEndian, &hash)

		if err == io.ErrUnexpectedEOF || err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		var strong = make([]byte, strongLen)
		err = binary.Read(signature, binary.BigEndian, &strong)
		if err == io.ErrUnexpectedEOF || err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		sig.hash[hash] = len(sig.strongHash)
		sig.strongHash = append(sig.strongHash, strong)
	}

	return &sig, nil
}
