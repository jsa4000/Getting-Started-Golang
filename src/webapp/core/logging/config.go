package logging

import (
	"fmt"
	"io"
	"os"
	"strings"
	"webapp/core/config"
)

// Config config for configuration
type Config struct {
	Output io.Writer
	Level  Level
	Format Format
}

var levelName = []string{
	PanicLevel: "panic",
	FatalLevel: "fatal",
	ErrorLevel: "error",
	WarnLevel:  "warn",
	InfoLevel:  "info",
	DebugLevel: "debug",
	TraceLevel: "trace",
}

var formatName = []string{
	TextFormat: "text",
	JSONFormat: "json",
}

func outputByName(s string) (io.Writer, error) {
	switch s {
	case "stderr":
		return os.Stderr, nil
	case "stdout":
		return os.Stdout, nil
	default:
		return nil, fmt.Errorf("Output '%s' cannot be converted", s)
	}
}

func levelByName(s string) (Level, error) {
	s = strings.ToLower(s)
	for i, name := range levelName {
		if name == s {
			return Level(i), nil
		}
	}
	return 0, fmt.Errorf("Level '%s' cannot be converted", s)
}

func formatByName(s string) (Format, error) {
	s = strings.ToLower(s)
	for i, name := range formatName {
		if name == s {
			return Format(i), nil
		}
	}
	return 0, fmt.Errorf("Format '%s' cannot be converted", s)
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := struct {
		Output string `config:"logging.output:stderr"`
		Level  string `config:"logging.level:debug"`
		Format string `config:"logging.format:text"`
	}{}
	config.ReadFields(&c)

	//Output
	output, err := outputByName(c.Output)
	if err != nil {
		output = os.Stderr
	}

	//Level
	level, err := levelByName(c.Level)
	if err != nil {
		level = DebugLevel
	}

	//Format
	format, err := formatByName(c.Format)
	if err != nil {
		format = TextFormat
	}

	return &Config{
		Output: output,
		Level:  level,
		Format: format,
	}
}
