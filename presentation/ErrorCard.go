package presentation

func ErrorCard(message string) string {
	return `
		<svg 
            height="${cardModifier.height}" 
            width="${cardModifier.width}" 
            xmlns="http://www.w3.org/2000/svg"  
            xmlns:xlink="http://www.w3.org/1999/xlink"
        > 
            <defs>
                <script>
                    setInterval(() => { location.reload() }, 60000) ;
                </script>
                <style>
                    .warning-text{
                        fill: #ffffff;
                        font-size: 30 ;
                    }
                </style>    
            </defs>

            <rect 
                height="${cardModifier.height}" 
                width="${cardModifier.width}" 
                x="0" 
                y="0"
                rx="${cardModifier.radius}"
                ry="${cardModifier.radius}"
            />

            <text class="warning-text" x="${error.x}" y="${error.y}">${error.text}</text>
        </svg>
	`
}
