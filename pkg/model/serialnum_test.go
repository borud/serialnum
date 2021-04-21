package model

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pair struct {
	sid SerialNum
	str string
}

// 000.037.2020.105.00140  --  C-adr: 252   [borud]
// 000.037.2020.093.00131  --  C-adr: 176   [stalehd]
// 000.037.2020.092.00118  --  C-adr: 188   [pkk]
// 000.037.2020.108.00039  --  C-adr: 254   [tlan]
// 000.037.2020.091.00031  --  C-adr: 223   [hjg]
// 000.037.2020.093.00022  --  C-adr: 172   [labben]

var testData = []pair{
	{
		sid: SerialNum([6]byte{0, 1, 2, 3, 4, 5}),
		str: "000.001.1995.008.00515",
	},
	{
		// tlan
		sid: SerialNum([6]byte{0, 37, 0, 39, 54, 30}),
		str: "000.037.2020.108.00039",
	},
	{
		// borud
		sid: SerialNum([6]byte{0, 37, 0, 140, 52, 158}),
		str: "000.037.2020.105.00140",
	},
	{
		// pkk
		sid: SerialNum([6]byte{0, 37, 0, 118, 46, 30}),
		str: "000.037.2020.092.00118",
	},
	{
		// stalehd
		sid: SerialNum([6]byte{0, 37, 0, 131, 46, 158}),
		str: "000.037.2020.093.00131",
	},
}

func TestSerialNum(t *testing.T) {
	// Simple equality test
	sid := SerialNum([6]byte{0x0, 0x25, 0x0, 0x1f, 0x2d, 0x9e})
	v := sid.ToSerialNumValues()
	assert.Equal(t, uint16(37), v.A)
	assert.Equal(t, uint16(31), v.B)
	assert.Equal(t, uint16(2020), v.Year)
	assert.Equal(t, uint16(91), v.Day)

	assert.Equal(t, uint64(158915833246), sid.Uint64())
	assert.Equal(t, uint64(158915833246), FromUint64(uint64(158915833246)).Uint64())

	scanSid := SerialNum{}
	scanSid.Scan(uint64(158915833246))
	assert.Equal(t, uint64(158915833246), scanSid.Uint64())

	// Make sure that we can convert to IDType and that none of the uint64/int64 assumptions bite us
	big := SerialNum([6]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	assert.Equal(t, big.IDType(), IDType(281474976710655))

	// Test the string formatting and parsing
	for _, p := range testData {
		parsed, err := ParseSerialNum(p.str)
		assert.Nil(t, err)
		assert.Equal(t, p.str, parsed.String())

		twice, err := ParseSerialNum(p.sid.String())
		assert.Nil(t, err)
		assert.Equal(t, p.str, twice.String())
	}

	// Test that the padding code behaves correctly
	assert.Equal(t, binary.BigEndian.Uint64([]byte{0, 0, 1, 2, 3, 4, 5, 6}), SerialNum([6]byte{1, 2, 3, 4, 5, 6}).Uint64())

}
