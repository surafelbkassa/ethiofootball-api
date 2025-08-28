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
	ctx := context.Background()

	// ✅ Correctly initialize the Gemini client
	client, err := genai.NewClient(ctx, option.WithAPIKey(c.apiKey))
	if err != nil {
		log.Printf("Failed to create genai client: %v", err)
		return nil, domain.ErrUnexpected
	}
	defer client.Close() // ✅ always close client

	// Marshal the context data for the prompt
	contextBytes, err := json.MarshalIndent(dCtx.ContextData, "", "  ")
	if err != nil {
		log.Printf("Could not marshal context data: %v", err)
		return nil, domain.ErrUnexpected
	}

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

	// ✅ Use the model directly
	model := client.GenerativeModel("gemini-1.5-flash-latest")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Failed to generate content: %v", err)
		return nil, domain.ErrUnexpected
	}

	// Extract markdown
	var markdownContent string
	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				markdownContent += string(text)
			}
		}
	}

	if markdownContent == "" {
		log.Println("Gemini returned an empty answer")
		return nil, errors.New("gemini returned an empty answer")
	}

	answer := &domain.Answer{
		Markdown:  markdownContent,
		Source:    dCtx.Source,
		Freshness: dCtx.Freshness,
	}

	return answer, nil
}
