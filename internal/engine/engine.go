package engine

import (
	"log"
	"strings"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/code-gorilla-au/pyrotic/internal/parser"
	"github.com/code-gorilla-au/pyrotic/internal/writer"
)

const (
	metaDelimiter         = ","
	metaKeyValueDelimiter = "="
)

func New(dryrun bool, dirPath string, sharedPath string, fileSuffix string) (Core, error) {
	if dryrun {
		log.Println(chalk.Cyan("DRYRUN MODE"))
	}
	tmpl, err := parser.New(dirPath, sharedPath, fileSuffix)
	if err != nil {
		return Core{}, err
	}
	return Core{
		parser: tmpl,
		fwr:    writer.New(dryrun),
	}, nil
}

// Generate - generates code from
func (c *Core) Generate(data Data) error {

	parsedOutput, err := c.parser.Parse(parser.TemplateData{
		Name: data.Name,
		ParseData: parser.ParseData{
			Meta: generateMeta(data.MetaArgs),
		},
	})
	if err != nil {
		return err
	}

	for _, item := range parsedOutput {
		switch item.ParseData.Action {
		case parser.ActionAppend:
			if err := c.fwr.AppendFile(item.To, item.Output); err != nil {
				log.Println("error appending file ", err)
				return err
			}
		case parser.ActionInject:
			if err := c.fwr.InjectIntoFile(item.To, item.Output, writer.Inject{
				Matcher: item.InjectMatcher,
				Clause:  writer.InjectClause(item.InjectClause),
			}); err != nil {
				return err
			}
		default:
			if err := c.fwr.WriteFile(item.To, item.Output, 0750); err != nil {
				log.Println("error writing to file ", err)
				return err
			}
		}

	}

	return nil
}

func generateMeta(meta string) (result map[string]string) {
	result = make(map[string]string)

	if len(meta) == 0 {
		return result
	}

	parts := splitIntoParts(meta)

	for _, part := range parts {
		key, value, ok := parseKeyValue(part)
		if !ok {
			return make(map[string]string)
		}
		result[key] = value
	}

	return
}

func splitIntoParts(meta string) []string {
	var parts []string
	var currentPart strings.Builder
	inQuotes := false

	for i := 0; i < len(meta); i++ {
		char := meta[i]

		if char == '"' {
			inQuotes = !inQuotes
		}

		if char == ',' && !inQuotes {
			parts = append(parts, currentPart.String())
			currentPart.Reset()
			continue
		}

		currentPart.WriteByte(char)
	}

	if currentPart.Len() > 0 {
		parts = append(parts, currentPart.String())
	}

	return parts
}

func processValue(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}

	if len(value) >= 2 && value[0] == '"' {
		quoteCount := strings.Count(value, `"`)
		if quoteCount >= 2 && value[len(value)-1] == '"' {
			return value[1 : len(value)-1]
		} else if quoteCount == 1 {
			return value[1:]
		}
	}
	return value
}

func parseKeyValue(part string) (string, string, bool) {
	part = strings.TrimSpace(part)
	if part == "" {
		return "", "", false
	}

	kv := strings.SplitN(part, "=", 2)
	if len(kv) != 2 {
		return "", "", false
	}

	key := strings.TrimSpace(kv[0])
	value := processValue(kv[1])

	// Validate key and unquoted empty value
	if key == "" || (value == "" && kv[1] == "") {
		return "", "", false
	}

	return key, value, true
}
