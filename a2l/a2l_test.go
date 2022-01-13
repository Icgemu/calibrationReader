package a2l

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestParseFromFile(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	a2lPath := "testing/ASAP2_Demo_V171.a2l"
	startTime := time.Now()
	a, err := ParseFromFile(a2lPath)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
	log.Info().Msg("time for parsing a2l test file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
}

func BenchmarkParseFromFile(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a2lPath := "testing/ASAP2_Demo_V171.a2l"
		startTime := time.Now()
		a, err := ParseFromFile(a2lPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
		log.Warn().Msg("time for parsing a2l bench file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
	}
}