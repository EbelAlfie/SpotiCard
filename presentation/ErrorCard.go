package presentation

import (
	"fmt"
	"html/template"
	"strings"
)

func ErrorCard(errorModel ErrorModel) string {
    cardScale := 2
    cardModifier := CardModifier{
		Width:  300 * cardScale,
		Height: 100 * cardScale,
		Radius: 8,
	}

    config := struct {
		Card   CardModifier
        Text   TextModifier
	}{
		Card:      cardModifier,
		Text:   TextModifier {
            X:    80 * cardScale,
		    Y:    15 * cardScale,
            Text: errorModel.Error.Error(),
        },
	}

	card := `
		<svg 
            height="{{.Card.Height}}" 
            width="{{.Card.Width}}" 
            xmlns="http://www.w3.org/2000/svg"  
            xmlns:xlink="http://www.w3.org/1999/xlink"
        > 
            <defs>
                <style>
                    .warning-text{
                        fill: #ffffff;
                        font-size: 30 ;
                    }
                </style>    
            </defs>

            <rect 
                height="{{.Card.Height}}" 
                width="{{.Card.Width}}" 
                x="0" 
                y="0"
                rx="{{.Card.Radius}}"
                ry="{{.Card.Radius}}"
            />

            <text class="warning-text" x="{{.Text.X}}" y="{{.Text.Y}}">{{.Text.Text}}</text>
        </svg>
	`

    var result strings.Builder
    
	template, err := template.New("ErrorCard").Parse(card)
	if err != nil {
		return fmt.Sprintf("Error Parsing String %s", err)
	}
	err = template.Execute(&result, config)
	if err != nil {
		return fmt.Sprintf("Error Parsing String %s", err)
	}

	return result.String()
}
