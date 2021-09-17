package plantuml

import "testing"

func TestEncodeStatement(t *testing.T) {
    testCase := func(input, expected string) {
        got, err := EncodeStatement(input)
        if err != nil {
            t.Error(err)
        }
        if got != expected {
            t.Errorf("expected '%s' got '%s'", expected, got)
        }
    }
    testCase("@startuml\nBob -> Alice : hello\n@enduml", "~1UDfoA2v9B2efpStXSifFKj2rKt3CoKnELR1Io4ZDoSddSaZDIodDpG44003___W93C00")
}

