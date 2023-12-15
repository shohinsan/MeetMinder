<script lang="ts">
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from 'svelte/elements';

	type Element = $$Generic<'button' | 'a'>;

	interface ButtonComponentElements {
		button: HTMLButtonAttributes;
		a: HTMLAnchorAttributes;
	}

	type $$Props = ButtonComponentElements[Element] & {
		element: Element;
		color?: 'red' | 'green' | 'cyan';
		className?: string;
		ariaLabel?: string;
		ariaPressed?: boolean;
		fullSize?: boolean;
		size?: 'sm' | 'md' | 'lg';
		variant?: 'solid' | 'outline' | 'text';
	};

	export let element: Element;
	export let color: 'red' | 'green' | 'cyan' = 'red';
	export let size: 'sm' | 'md' | 'lg' = 'md';
	export let ariaLabel: string = '';
	export let ariaPressed: boolean = false;
	export let className: string = '';
	export let variant: 'solid' | 'outline' | 'text' = 'solid';
	export let fullSize: boolean = false;
	let inputId = crypto.randomUUID();

	$: styles = `
			${className}
			${size in rounded ? rounded[size] : rounded['md']}
			${variant === 'outline' || variant === 'text' ? '' : color in colors ? colors[color] : colors['red']}
			${size in sizes ? sizes[size] : sizes['md']}
			${variant === 'text' ? '' : variant in variants ? variants[variant] : variants['solid']}
			${fullSize ? 'w-full' : ''}
`;
</script>

<svelte:element
	this={element}
	id={inputId}
	class={styles}
	on:click
	aria-label={ariaLabel}
	aria-pressed={ariaPressed}
	role="button"
	tabindex="0"
	{...$$restProps}
>
	<slot />
</svelte:element>

<script context="module">
	const sizes = {
		sm: 'px-1 py-3',
		md: 'px-5 py-3',
		lg: 'px-7 py-5',
	};

	const colors = {
		red: 'bg-red-600 hover:bg-red-800',
		green: 'bg-green-600 hover:bg-green-800',
		cyan: 'bg-cyan-600',
	};

	const rounded = {
		sm: 'rounded-sm',
		md: 'rounded-md',
		lg: 'rounded-lg',
	};

	const variants = {
		solid: 'text-white',
		outline: 'border border-solid border-current transition-colors duration-300',
		text: 'text-current',
	};

</script>
