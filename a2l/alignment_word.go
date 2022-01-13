package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

//alignmentFloat32Ieee is necessary because in complex objects (maps and axis) the alignment of a value may not coincide with the bitwidth of a value.
//This keyword is used to define the alignment in the case of words.
type alignmentWord struct {
	//alignmentBorder describes the border at which the value is aligned to, i.e.
	//its memory address must be dividable by the value AlignmentBorder.
	alignmentBorder    uint16
	alignmentBorderSet bool
}

func parseAlignmentWord(tok *tokenGenerator) (alignmentWord, error) {
	aw := alignmentWord{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("alignmentWord could not be parsed")
	} else if !aw.alignmentBorderSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("alignmentWord alignmentBorder could not be parsed")
		}
		aw.alignmentBorder = uint16(buf)
		aw.alignmentBorderSet = true
		log.Info().Msg("alignmentWord alignmentBorder successfully parsed")
	}
	return aw, err
}
