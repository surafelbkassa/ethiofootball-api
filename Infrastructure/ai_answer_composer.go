package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIAnswerComposer struct {
	apiKey string
}

func NewAIAnswerComposer(apiKey string) *AIAnswerComposer {
	return &AIAnswerComposer{apiKey: apiKey}
}

func (c *AIAnswerComposer) ComposeAnswer(dCtx domain.AnswerContext) (*domain.Answer, error) {
	// If the topic is "compare", use the special JSON generation logic.
	if dCtx.Topic == "compare" {
		return c.composeComparisonJSON(dCtx)
	}

	// Otherwise, use the default markdown generation logic.
	return c.composeMarkdown(dCtx)
}

// --- Private helper for generating Markdown ---
func (c *AIAnswerComposer) composeMarkdown(dCtx domain.AnswerContext) (*domain.Answer, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(c.apiKey))
	if err != nil {
		return nil, domain.ErrUnexpected
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-latest")
	contextBytes, _ := json.MarshalIndent(dCtx.ContextData, "", "  ")

	language := "English"
	if dCtx.Language == "am" {
		language = "Amharic"
	}

	prompt := fmt.Sprintf(`You are a helpful and concise football assistant for Ethiopian fans.
	Your task is to write a short, friendly summary in %s using ONLY the data provided below.

	**Rules:**
	- Use ONLY the provided data. Do not make up scores, fixtures, or facts.
	- If a piece of information is missing from the data, say "it is not available" or "is not confirmed."
	- The output MUST be markdown.
	- The tone should be friendly and respectful of all clubs.
	- NO betting or gambling language.

	**Provided Data (JSON format):**
	%s

	Now, write the summary:`, language, string(contextBytes))

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, domain.ErrUnexpected
	}

	var markdownContent string
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil { // Simplified check
		part := resp.Candidates[0].Content.Parts[0]
		if txt, ok := part.(genai.Text); ok {
			markdownContent = string(txt)
		}
	}
	if markdownContent == "" {
		return nil, domain.ErrUnexpected
	}

	// Construct an Answer object with ONLY the Markdown field populated.
	answer := &domain.Answer{
		Markdown:  markdownContent,
		Source:    dCtx.Source,
		Freshness: dCtx.Freshness,
	}
	return answer, nil
}

func (c *AIAnswerComposer) composeComparisonJSON(dCtx domain.AnswerContext) (*domain.Answer, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(c.apiKey))
	if err != nil {
		log.Printf("Failed to create genai client: %v", err)
		return nil, domain.ErrUnexpected
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-latest")
	// **CRITICAL**: We tell the model to ONLY output JSON.
	model.ResponseMIMEType = "application/json"

	contextBytes, _ := json.MarshalIndent(dCtx.ContextData, "", "  ")

	// This is the new, highly specific prompt for generating JSON.
	jsonPrompt := fmt.Sprintf(`You are a data processing API. Your sole purpose is to populate a JSON object that compares two football teams based on the provided data.

	**Your response MUST be ONLY the JSON object and nothing else.**

	**JSON Schema to populate:**
	{
	"team_a": {
		"name": "string",
		"honors": ["string"],
		"recent_form": ["W", "D", "L", "W", "W"],
		"notable_players": ["string"],
		"fanbase_notes": "string"
	},
	"team_b": {
		"name": "string",
		"honors": ["string"],
		"recent_form": ["W", "D", "L", "W", "W"],
		"notable_players": ["string"],
		"fanbase_notes": "string"
	}
	}

	**Provided Data:**
	%s

	Fill the JSON object using the data above. If a piece of information is missing, use an empty array [] or an empty string "".`, string(contextBytes))

	resp, err := model.GenerateContent(ctx, genai.Text(jsonPrompt))
	if err != nil {
		log.Printf("Failed to generate comparison JSON: %v", err)
		return nil, domain.ErrUnexpected
	}

	// Extract the JSON string from the response
	var jsonText string
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		part := resp.Candidates[0].Content.Parts[0]
		if txt, ok := part.(genai.Text); ok {
			jsonText = string(txt)
		}
	}
	if jsonText == "" {
		return nil, errors.New("gemini returned empty JSON for comparison")
	}

	// Unmarshal the generated JSON string into our domain struct
	var comparisonData domain.ComparisonData
	if err := json.Unmarshal([]byte(jsonText), &comparisonData); err != nil {
		log.Printf("Failed to unmarshal comparison JSON: %v. Raw response: %s", err, jsonText)
		return nil, domain.ErrUnexpected
	}

	// Construct an Answer object with ONLY the ComparisonData field populated.
	answer := &domain.Answer{
		ComparisonData: &comparisonData,
		Source:         dCtx.Source,
		Freshness:      dCtx.Freshness,
	}
	return answer, nil
}
