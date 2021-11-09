package a2l

import (
	"github.com/rs/zerolog/log"
	"errors"
	"strconv"
)

type ripAddr4 struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseRipAddr4(tok *tokenGenerator) (ripAddr4, error) {
	ra4 := ripAddr4{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
				log.Err(err).Msg("ripAddr4 could not be parsed")
			break forLoop
		} else if !ra4.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
					log.Err(err).Msg("ripAddr4 position could not be parsed")
				break forLoop
			}
			ra4.position = uint16(buf)
			ra4.positionSet = true
				log.Info().Msg("ripAddr4 position successfully parsed")
		} else if !ra4.datatypeSet {
			var buf dataTypeEnum
			buf, err = parseDataTypeEnum(tok)
			if err != nil {
					log.Err(err).Msg("ripAddr4 datatype could not be parsed")
				break forLoop
			}
			ra4.datatype = buf
			ra4.datatypeSet = true
				log.Info().Msg("ripAddr4 datatype successfully parsed")
			break forLoop
		}
	}
	return ra4, err
}
