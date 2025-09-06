package infrastructure

import (
	"context"
	"encoding/json"
	"log"

	"github.com/abrshodin/ethio-fb-backend/Domain"

	"google.golang.org/genai"
)

type AIIntentParser struct {
	apiKey string
}

func NewAIIntentParser(apiKey string) *AIIntentParser {
	return &AIIntentParser{apiKey}
}

func (ip AIIntentParser) Parse(text string) (*domain.Intent, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Print(err)
		return nil, domain.ErrUnexpected
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"topic": {
					Type: genai.TypeString,
					Enum: []string{"fixture", "table", "compare", "news", "fact"},
				},
				"teams": {
					Type: genai.TypeArray,
					Items: &genai.Schema{
						Type: genai.TypeString,
					},
				},
				"league":    {Type: genai.TypeString},
				"date":      {Type: genai.TypeString},
				"follow_up": {Type: genai.TypeString},
			},
			Required: []string{"topic", "teams", "league"},
		},
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(text),
		config,
	)

	if err != nil {
		log.Print(err)
		return nil, domain.ErrUnexpected
	}

	var parsed domain.Intent
	err = json.Unmarshal([]byte(result.Text()), &parsed)
	if err != nil {
		log.Print(err)
		return nil, domain.ErrUnexpected
	}

	return &parsed, nil
}
