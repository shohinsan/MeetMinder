<script>
	export let width = 32
	export let height = 32
	export let fill = '#000'
</script>

<svg {width} {height} role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
	{#each Array.from({ length: 12 }) as _, i}
		<circle
			{fill}
			class="spinner"
			cx={12 + 8.5 * Math.cos((i * Math.PI) / 6)}
			cy={12 + 8.5 * Math.sin((i * Math.PI) / 6)}
			r="0" />
	{/each}
</svg>

<style lang="scss">
	@mixin animate($delay) {
		animation: load 1.2s infinite;
		animation-delay: $delay;
	}

	@for $i from 0 through 11 {
		.spinner:nth-child(#{$i + 1}) {
			@include animate($i * 0.1s);
		}
	}

	@keyframes load {
		0%,
		50% {
			animation-timing-function: cubic-bezier(0, 1, 0, 1);
			r: 0;
		}
		10% {
			animation-timing-function: cubic-bezier(0.53, 0, 0.61, 0.73);
			r: 2px;
		}
	}
</style>
