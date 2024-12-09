package presentation

import (
	"fmt"
	"html/template"
	"strings"

	"spoti-card.com/domain/entity"
)

func SpotifyCard(track entity.TrackEntity) string {
	cardScale := 2

	cardModifier := CardModifier{
		Width:  300 * cardScale,
		Height: 100 * cardScale,
		Radius: 8,
	}

	imageModifier := ImageModifier{
		Width:  64 * cardScale,
		Height: 64 * cardScale,
		X:      10,
		Y:      10,
		Url:    track.Album.Images[0].Url,
	}

	audioModifier := AudioModifier{
		Url: track.PreviewUrl,
	}

	equalizer := EqualizerModifier{
		Y: 80 * cardScale,
	}

	title := TextModifier{
		X:    80 * cardScale,
		Y:    15 * cardScale,
		Text: track.Name,
	}

	caption := TextModifier{
		X:    80 * cardScale,
		Y:    25 * cardScale,
		Text: "",
	}

	status := TextModifier{
		X:    30,
		Y:    equalizer.Y,
		Text: "Is Playing",
	}

	config := struct {
		Card      CardModifier
		Image     ImageModifier
		Audio     AudioModifier
		Equalizer EqualizerModifier
		Title     TextModifier
		Caption   TextModifier
		Status    TextModifier
	}{
		Card:      cardModifier,
		Image:     imageModifier,
		Audio:     audioModifier,
		Equalizer: equalizer,
		Title:     title,
		Caption:   caption,
		Status:    status,
	}

	card := `
		<svg 
        height="{{.Card.Height}}" 
        width="{{.Card.Width}}" 
        xmlns="http://www.w3.org/2000/svg"  
        xmlns:xlink="http://www.w3.org/1999/xlink"
    >
        <style>
            .album-image {
                border-radius: 4px;
            }

            .song-title {
                fill: #ffffff;
                font-size: 20;
            }

            .song-artist {
                fill: #b3b3b3;
                font-size: 15;
            }

            .track_status {
                fill: #ffffff
                font-size: 10;
            }

            .equalizer {
                height: 2px;
                rx: 1;
                fill: #1DB954;
                transform-box: fill-box;
                transform: rotate(180deg) ;
                animation-iteration-count: infinite;
                animation-name: equalizer-anim;
            }

            .eq-slow {
                animation-duration: 1.2s;
            }

            .eq-medium {
                animation-duration: 1.0s;
            }

            .eq-quick {
                animation-duration: 0.8s;
            }

            .eq-fast {
                animation-duration: 0.6s;
            }

            @keyframes equalizer-anim {
                0%, 100% {
                    height: 2px;
                }
                50% {
                    height: 8px;
                }
                60% {
                    height: 11px;
                }
            }
        </style>
        <video autoplay loop>
            <source src="{{.Audio.Url}}" type="audio/mpeg" />
        </video>
        <script>
            setInterval(() => { location.reload() }, 60000) ;
        </script>
        <rect 
            height="{{.Card.Height}}" 
            width="{{.Card.Width}}" 
            x="0" 
            y="0"
            rx="{{.Card.Radius}}"
            ry="{{.Card.Radius}}"
        />
        <text 
            class="song-title" 
            fill="white" x="{{.Title.x}}" y="{{.Title.y}}"
        >
            {{.Title.Text}}
        </text>
        <text 
            class="song-artist" 
            fill="white" x="{{.Caption.X}}" y="{{.Caption.Y}}"
        >
            {.Caption.Text}}
        </text>
        <image 
            class="album-image" 
            height="{{.Image.Height}}" 
            width="{{.Image.Width}}" 
            x="{{.Image.X}}" 
            y="{{.Image.Y}}" 
            href="{{.Image.Url}}" 
            alt="Track Cover"
            clip-path="inset(0% round 15px)"
        />
        <rect class="equalizer eq-fast" x="16" y="{{.Equalizer.Y}}" width="2"/>
        <rect class="equalizer eq-medium" x="19" y="{{.Equalizer.Y}}" width="2"/>
        <rect class="equalizer eq-quick" x="22" y="{{.Equalizer.Y}}" width="2"/>
        <rect class="equalizer eq-slow" x="25" y="{{.Equalizer.Y}}" width="2"/>
        
        <text 
            class="track-status" 
            fill="white" 
            x="{{.Status.X}}" 
            y="{{.Status.Y}}"
        >
            {{.Status.Text}}
        </text>
    </svg>
	`

	var result strings.Builder
	template, err := template.New("SpotiCard").Parse(card)
	if err != nil {
		return fmt.Sprintf("Error Parsing String %s", err)
	}
	err = template.Execute(&result, config)
	if err != nil {
		return fmt.Sprintf("Error Parsing String %s", err)
	}

	return result.String()
}
