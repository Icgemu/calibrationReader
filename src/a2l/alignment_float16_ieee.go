package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

//alignmentFloat32Ieee is necessary because in complex objects (maps and axis) the alignment of a value may not coincide with the bitwidth of a value.
//This keyword is used to define the alignment in the case of 16bit floats.
type alignmentFloat16Ieee struct {
	//alignmentBorder describes the border at which the value is aligned to, i.e.
	//its memory address must be dividable by the value AlignmentBorder.
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentFloat16Ieee(tok *tokenGenerator) (alignmentFloat16Ieee, error) {
	af16 := alignmentFloat16Ieee{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentFloat16Ieee could not be parsed")
	} else if !af16.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentFloat16Ieee alignmentBorder could not be parsed")
		}
		af16.alignmentBorder = uint16(buf)
		af16.alignmentBorderSet = true
		log.Info().Msg("alignmentFloat16Ieee alignmentBorder successfully parsed")
	}
	return af16, err
}