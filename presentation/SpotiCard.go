package presentation

func SpotifyCard() string {
	return `
		<svg 
			height="100" 
			width="300" 
			xmlns="http://www.w3.org/2000/svg"  
			xmlns:xlink="http://www.w3.org/1999/xlink"
		>
			<style>
				.equalizer {
					stroke-width: 5;
					stroke: #1DB954;
					animation-iteration-count: infinite;
				}
				
			</style>
			<rect 
				height="100" 
				width="300" 
				x="0" 
				y="0"
			/>
			<text fill="white" x="80" y="25">HEHEHE</text>
			<text fill="white" x="80" y="50">HEHEHE</text>
			<image 
				height="64" 
				width="64" 
				x="10"
				y="10"
				id="album-image" 
				href="https://i.scdn.co/image/ab67616d0000b273232711f7d66a1e19e89e28c5" 
				alt="Track Cover"
			/>
			<line 
				class="equalizer"
				x1="8" y1="90" x2="8" y2="70"
			>
			</line> 
		</svg>
	`
}
