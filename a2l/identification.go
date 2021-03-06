package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

type identification struct {
	position    uint16
	positionSet bool
	datatype    dataTypeEnum
	datatypeSet bool
}

func parseIdentification(tok *tokenGenerator) (identification, error) {
	i := identification{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("identification could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("identification could not be parsed")
			break forLoop
		} else if !i.positionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("identification position could not be parsed")
				break forLoop
			}
			i.position = uint16(buf)
			i.positionSet = true
			log.Info().Msg("identification position successfully parsed")
		} else if !i.datatypeSet {
			i.datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("identification position could not be parsed")
				break forLoop
			}
			i.datatypeSet = true
			log.Info().Msg("identification datatype successfully parsed")
			break forLoop
		}
	}
	return i, err
}
