package ihex32

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestParseFromFile(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	hexPath := "testing/ASAP2_Demo_V171.hex"
	h, err := ParseFromFile(hexPath)
	if err == nil {
		if len(h) != 137324 {
			err = errors.New("wrong length of dataset")
			t.Fatalf("failed parsing with error: %s.", err)
		}
	} else {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func BenchmarkParseFromFile(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		startTime := time.Now()
		hexPath := "testing/ASAP2_Demo_V171.hex"
		h, err := ParseFromFile(hexPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
			log.Info().Msg("length of data in hex file: " + fmt.Sprint(len(h)))
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Warn().Msg("time for parsing ihex32 bench file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
	}
}
