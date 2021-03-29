package rhdiff

// Rolling Hash
// https://librsync.github.io/rabinkarp_8h_source.html

const (

	/** The RabinKarp seed value.
	 *
	 * The seed ensures different length zero blocks have different hashes. It
	 * effectively encodes the length into the hash. */
	RABINKARP_SEED = 1
	/** The RabinKarp multiplier.
	 *
	 * This multiplier has a bit pattern of 1's getting sparser with significance,
	 * is the product of 2 large primes, and matches the characterstics for a good
	 * LCG multiplier. */
	RABINKARP_MULT MagicNumber = 0x08104225

	/** The RabinKarp inverse multiplier.
		*
	  * This is the inverse of RABINKARP_MULT modular 2^32. Multiplying by this is
		* equivalent to dividing by RABINKARP_MULT. */
	RABINKARP_INVM MagicNumber = 0x98f009ad

	/** The RabinKarp seed adjustment.
	*
	* This is a factor used to adjust for the seed when rolling out values. It's
	* equal to; (RABINKARP_MULT - 1) * RABINKARP_SEED */
	RABINKARP_ADJ MagicNumber = 0x08104224
)

type RollingHash struct {
	count uint64
	hash  uint32
	multi MagicNumber
}

func NewRollingHash() RollingHash {
	return RollingHash{
		count: 0,
		hash:  RABINKARP_SEED,
		multi: 1,
	}
}

/* Table of RABINKARP_MULT^(2^(i+1)) for power lookups. */
var RABINKARP_MULT_POW2 = [32]uint32{
	0x08104225,
	0xa5b71959,
	0xf9c080f1,
	0x7c71e2e1,
	0x0bb409c1,
	0x4dc72381,
	0xd17a8701,
	0x96260e01,
	0x55101c01,
	0x2d303801,
	0x66a07001,
	0xfe40e001,
	0xc081c001,
	0x91038001,
	0x62070001,
	0xc40e0001,
	0x881c0001,
	0x10380001,
	0x20700001,
	0x40e00001,
	0x81c00001,
	0x03800001,
	0x07000001,
	0x0e000001,
	0x1c000001,
	0x38000001,
	0x70000001,
	0xe0000001,
	0xc0000001,
	0x80000001,
	0x00000001,
	0x00000001,
}

func (r *RollingHash) Update(data []byte) {
	data_len := len(data)
	// hash := r.hash

	for n := 0; n < data_len; n++ {
		r.hash = r.hash*uint32(RABINKARP_MULT) + uint32(data[n])
		r.multi *= RABINKARP_MULT
		r.count++
	}
	// r.hash = hash
	// r.count += uint64(data_len)

}

func (r *RollingHash) Rotate(out, in byte) {
	r.hash = r.hash*uint32(RABINKARP_MULT) + uint32(in) - uint32(r.multi)*(uint32(out)+uint32(RABINKARP_ADJ))
}

func (r *RollingHash) Rollin(in byte) {
	r.hash = r.hash*uint32(RABINKARP_MULT) + uint32(in)
	r.count++
	r.multi *= RABINKARP_MULT
}

func (r *RollingHash) Rollout(out byte) {
	r.count--
	r.multi *= RABINKARP_INVM
	r.hash -= uint32(r.multi) * (uint32(out) + uint32(RABINKARP_ADJ))
}

func (r *RollingHash) Digest() uint32 {
	return r.hash
}

func (r *RollingHash) Reset() {
	r.count = 0
	r.hash = RABINKARP_SEED
	r.multi = 1
}
