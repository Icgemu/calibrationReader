package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*Description of rescaling the axis values of an adjustable object. A rescale axis consists
mainly of a number of rescaling axis points pairs (axis i , virtual i ) which describe a rescale
mapping between the axis points and a virtual axis that is used for the access of the table
function values deposited in the control unit. Between two pairs the mapping is linear.
Both, the axis points and the virtual axis points must be in ascending order. Consider, for
example, the three rescale pairs (0x00, 0x00), (0x64, 0xC0) and (0xD8, 0xFF). Then all
axis points between 0x00 and 0x64 are mapped linear to the virtual axis [0x00, 0xC0], and
all axis points between 0x64 and 0xD8 are mapped linear to the virtual axis [0xC0, 0xFF]:
Accordingly, to each axis point there is a virtual axis point. The virtual axis points are
distributed equidistantly on the virtual axis including the axis limits, e.g. the virtual axis
points can be derived from the size of the virtual axis and the number of axis points.
According to the rescale mapping the axis point can be computed from the virtual axis
points. The following algorithm can be applied, where D is the length of the (equidistant)
intervals on virtual axis:
The following example makes clear how the evaluation of the formula can be used to
derive the actual axis points. We have no_of_rescale_pairs = 3 and virtual 1 = 0x00 = 0,
virtual 2 = 0xC0 = 192, virtual 3 = 0xFF = 255, axis 1 = 0x00 = 0, axis 2 = 0x64 = 100, axis 3 =
0xD8 = 216. Assume no_axis_pts = 9, and therefore D = 32. The first of the two
executions of the inner loop (j-loop) is on virtual 2 – virtual 1 / D = 192/32 = 6 iterations. For
each iteration (axis 2 – axis 1 )/(virtual 2 – virtual 1 ) = 100/192, and therefore
X 2 = 0 + 32 * 100/192 = 16,666,
X 3 = 0 + 64 * 100/192 = 33,333,
X 4 = 0 + 96 * 100/192 = 50,
X 5 = 0 + 128 * 100/192 =66,666,
X 6 = 0 + 160 * 100/192 = 83,333.
For the second execution there are virtual 3 – virtual 2 / D = 2 iterations with (axis 3 –
axis 2 )/(virtual 3 – virtual 2 ) = 116/64. Consequently
X 7 = 100 + (192 – 192) * 116/64 = 100 and
X 8 = 100 + (224 – 192) * 116/64 = 158.
Also X 1 = axis 1 = 0 and X 9 = axis 3 = 216.*/
type axisRescale4 struct {
	position                   uint16
	positionSet                bool
	datatype                   dataTypeEnum
	datatypeSet                bool
	maxNumberOfRescalePairs    uint16
	maxNumberOfRescalePairsSet bool
	indexIncr                  IndexOrderEnum
	indexIncrSet               bool
	adressing                  AddrTypeEnum
	adressingSet               bool
}

func parseAxisRescale4(tok *tokenGenerator) (axisRescale4, error) {
	ar4 := axisRescale4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisRescale4 could not be parsed")
			break forLoop
		} else if !ar4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisRescale4 position could not be parsed")
				break forLoop
			}
			ar4.position = uint16(buf)
			ar4.positionSet = true
			log.Info().Msg("axisRescale4 position successfully parsed")
		} else if !ar4.datatypeSet {
			ar4.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescale4 datatype could not be parsed")
				break forLoop
			}
			ar4.datatypeSet = true
			log.Info().Msg("axisRescale4 datatype successfully parsed")
		} else if !ar4.maxNumberOfRescalePairsSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisRescale4 maxNumberOfRescalePairs could not be parsed")
				break forLoop
			}
			ar4.maxNumberOfRescalePairs = uint16(buf)
			ar4.maxNumberOfRescalePairsSet = true
			log.Info().Msg("axisRescale4 maxNumberOfRescalePairs successfully parsed")
		} else if !ar4.indexIncrSet {
			ar4.indexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescale4 indexIncr could not be parsed")
				break forLoop
			}
			ar4.indexIncrSet = true
			log.Info().Msg("axisRescale4 indexIncr successfully parsed")
		} else if !ar4.adressingSet {
			ar4.adressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescale4 adressing could not be parsed")
				break forLoop
			}
			ar4.adressingSet = true
			log.Info().Msg("axisRescale4 adressing successfully parsed")
			break forLoop
		}
	}
	return ar4, err
}
