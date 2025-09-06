package infrastructure

import (
	"context"
	"encoding/json"
	"log"
	
	domain "github.com/abrshodin/ethio-fb-backend/Domain"
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
				"language" : {Type: genai.TypeString},
			},
			Required: []string{"topic", "teams", "league"},
		},
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(
			`System Prompt: Detect League Context: When a user asks about a match, standings, results, 
			live scores, or any league-related query, identify that the request is 
			about football.Default assumption: unless a specific league is mentioned, provide 
			information for Ethiopian Premier League (ETH) and English Premier League (EPL) in 
			the specified order. Response Order: Step 1: Provide data for the Ethiopian Premier 
			League (ETH). Step 2: Provide data for the English Premier League (EPL). Returns only 
			shorts for premier league 'ETH' for Ethiopian 'EPL' for English Keep the 
			order consistent: ETH first, EPL second. Language Handling: If the user writes in 
			Amharic or explicitly wants to interact in Amharic, all responses, including headings 
			and match details, should be in Amharic. Include a field in the response 'language': 
			'amharic'. If the user writes in English or wants English responses, respond in English 
			and set language": "english.Auto-detect language preference from the user prompt and adjust 
			accordingly.` + "\n\nuser prompt" + text,
		),
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
