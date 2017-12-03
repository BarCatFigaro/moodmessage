package stringsem

import (
    language "cloud.google.com/go/language/apiv1"
    "context"
    languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
    "log"
    "fmt"
)

func IsGood(msg string) bool {
    score, magnitude := getSentiment(msg)

    fmt.Printf("Score: %f, Magnitude: %f\n",score,magnitude)

    if score >= 0.5 && magnitude >= 0.5 {
        fmt.Println("good")
        return true
    } else {
        fmt.Println("bad")
        return false
    }
}

func IsHappy(msg []string) bool {
    happyScore := float32(0.0)

    for _, item := range msg {
        score, magnitude := getSentiment(item)
        happyScore += score*magnitude
    }

    if happyScore >= 0 {
        return true
    } else {
        return false
    }
}

func getSentiment(msg string) (float32, float32) {
    ctx := context.Background()

    client, err := language.NewClient(ctx)
    if err != nil {
        log.Fatalf("Failed to create client: %v",err)
    }

    sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
        Document: &languagepb.Document{
            Source: &languagepb.Document_Content{
                Content: msg,
                },
                Type: languagepb.Document_PLAIN_TEXT,
                },
                EncodingType: languagepb.EncodingType_UTF8,
                })
    if err != nil {
        log.Fatalf("Failed to analyze msg: %v", err)
    }

    return sentiment.DocumentSentiment.Score, sentiment.DocumentSentiment.Magnitude
}


