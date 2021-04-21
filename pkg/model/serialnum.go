package model

import (
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// IDType is the data type of IDs
type IDType int64

// SerialNum represents the 6 byte value for serial numbers on
// Autronica devices
type SerialNum [6]byte

// SerialNumValues contains the decoded values of SocketID
type SerialNumValues struct {
	A    uint16
	B    uint16
	Year uint16
	Day  uint16
}

// ParseSerialNum string format
func ParseSerialNum(s string) (SerialNum, error) {
	parts := strings.Split(s, ".")

	a1, err := strconv.ParseInt(parts[0], 10, 8)
	if err != nil {
		return SerialNum{}, err
	}

	a2, err := strconv.ParseInt(parts[1], 10, 8)
	if err != nil {
		return SerialNum{}, err
	}

	year, err := strconv.ParseInt(parts[2], 10, 16)
	if err != nil {
		return SerialNum{}, err

	}
	day, err := strconv.ParseInt(parts[3], 10, 16)
	if err != nil {
		return SerialNum{}, err

	}
	b, err := strconv.ParseInt(parts[4], 10, 16)
	if err != nil {
		return SerialNum{}, err
	}

	return SerialNum([6]byte{
		byte(a1),
		byte(a2),
		byte(b >> 8),
		byte(b & 0xff),
		byte(day >> 1),
		byte(((year - 1990) & 0x7f) | (day << 7)),
	}), nil
}

// FromBytes takes a byte representation and creates a SerialNum
func FromBytes(bb []byte) SerialNum {
	var buf [6]byte
	if len(bb) != 6 {
		return SerialNum{}
	}

	for i := 0; i < 6; i++ {
		buf[i] = bb[i]
	}
	return SerialNum(buf)
}

// FromUint64 creates a SerialNum from uint64
func FromUint64(ui uint64) SerialNum {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, ui)
	var serial SerialNum
	copy(serial[:], buf[2:])
	return serial
}

// String returns
func (s SerialNum) String() string {
	v := s.ToSerialNumValues()
	return fmt.Sprintf("%03d.%03.03d.%03d.%03d.%05d", v.A>>8, 0xff&v.A, v.Year, v.Day, v.B)
}

// IDType returns serial number as IDType
func (s SerialNum) IDType() IDType {
	// This should be safe since we only use the lower 6 bytes
	return IDType(s.Uint64())
}

// ToSerialNumValues returns a SerialNumValues representation of the SerialNum
func (s SerialNum) ToSerialNumValues() SerialNumValues {
	return SerialNumValues{
		A:    binary.BigEndian.Uint16(s[0:2]),
		B:    binary.BigEndian.Uint16(s[2:4]),
		Year: binary.BigEndian.Uint16(s[4:6])&0x7f + 1990,
		Day:  binary.BigEndian.Uint16(s[4:6]) >> 7,
	}
}

// Uint64 converts socket ID to uint64 so that it can be used as an id
func (s SerialNum) Uint64() uint64 {
	return binary.BigEndian.Uint64(append([]byte{0, 0}, s[:]...))
}

// Value serializes this data type for the SQL API
func (s SerialNum) Value() (driver.Value, error) {
	return s.Uint64(), nil
}

// Scan deseralize this data type for SQL API
func (s *SerialNum) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.(uint64); ok {
		vv := FromUint64(v)
		copy(s[:], vv[:])
		return nil
	}

	return errors.New("failed to Scan SerialNum")
}
