package main

import (
	"encoding/json"
	"go-embedded-scripting-comparison/pkg/file"
	"go-embedded-scripting-comparison/pkg/jsonnet"
	"go-embedded-scripting-comparison/pkg/lua"
	"go-embedded-scripting-comparison/pkg/model"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/adimunteanu/gluamapper"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Set up zerolog time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Set pretty logging on
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	args := os.Args[1:]

	scriptPath := args[0]

	if !file.CheckIfFileExists(scriptPath) {
		log.Fatal().Msgf("script file not found %s", scriptPath)
	}

	script, err := ioutil.ReadFile(scriptPath)

	if err != nil {
		log.Fatal().Err(err).Msg("Could not read script from file")
	}
	start := time.Now()
	input := model.RawWindDataPayload{
		ParkId:       "b86163c9-346e-45c9-93bb-dc77a22a5813",
		TurbineId:    "4deca05f-ceb7-474d-9c28-1e00b0c7521c",
		Region:       "Berlin",
		Date:         "2022-11-19",
		Interval:     19,
		Timezone:     "GMT+2",
		Value:        0.03,
		Availability: 95,
	}

	if strings.HasSuffix(scriptPath, ".lua") {
		for i := 1; i < 1000; i++ {
			output, err := lua.RunScript(script, "processing", input)

			if err != nil {
				log.Error().Err(err).Msg("issue running script")
				return
			}

			var processedEvent model.AnomalyDetectionPayload
			opt := gluamapper.Option{ErrorUnused: true, ErrorUnset: true, WeaklyTypedInput: false}
			mapper := gluamapper.NewMapper(opt)
			err = mapper.Map(output, &processedEvent)

			if err != nil {
				log.Error().Err(err).Msg("failed to cast a lua table")
				return
			}

			_, err = json.Marshal(processedEvent)

			if err != nil {
				log.Error().Err(err).Msg("could not marshall output of script")
				return
			}

			// log.Info().Msgf("Output: %s", string(msgData))
		}

	} else {
		var inputJson, err = json.Marshal(input)
		for i := 1; i < 1000; i++ {

			if err != nil {
				log.Error().Err(err).Msg("could not marshall input of script")
				return
			}

			_, err := jsonnet.WithJsonnet(script, inputJson)
			if err != nil {
				log.Error().Err(err).Msg("could not marshall output of script")
				return
			}
			// log.Info().Msgf("Output: %s", string(output))
		}
	}

	elapsed := time.Since(start)
	log.Info().Msgf("Output: %s", elapsed)
}
