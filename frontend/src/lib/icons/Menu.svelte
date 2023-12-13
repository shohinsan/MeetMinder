<script lang="ts">
	export let hamburger = false
	export let menuOpen = (): void => {
		hamburger = !hamburger
	}
	export let stroke = 'currentColor'
	export let ariaLabel = 'toggle menu'
	export let width: string | number = 32
</script>

<button on:click={menuOpen} aria-expanded={hamburger} aria-label={ariaLabel}>
	<svg class:hamburger viewBox="0 0 100 100" fill="none" {stroke} stroke-width="5" {width}>
		<path
			class="top"
			d="m 30,33 h 40 c 3.722839,0 7.5,3.126468 7.5,8.578427 0,5.451959 -2.727029,8.421573 -7.5,8.421573 h -20" />
		<path class="middle" d="m 30,50 h 40" />
		<path
			class="bottom"
			d="m 70,67 h -40 c 0,0 -7.5,-0.802118 -7.5,-8.365747 0,-7.563629 7.5,-8.634253 7.5,-8.634253 h 20" />
	</svg>
</button>

<style lang="postcss">

	button {
		cursor: pointer;
		display: flex;
		align-items: center;
		border: none;
		background: transparent;
		overflow: hidden;

		&:hover {
			color: green;
		}
	}

	svg {
		transition: transform var(--transition-duration);
	}

	.top {
		stroke-dasharray: 40 160;
		transition: stroke-dashoffset var(--transition-duration);
	}

	.middle {
		transform-origin: 50%;
		transition: transform var(--transition-duration);
	}

	.bottom {
		stroke-dasharray: 40 85;
		transition: stroke-dashoffset var(--transition-duration);
	}

	.hamburger {
		transform: rotate(45deg);
	}

	.hamburger .top,
	.hamburger .bottom {
		stroke-dashoffset: -64px;
	}

	.hamburger .middle {
		transform: rotate(90deg);
	}
</style>
