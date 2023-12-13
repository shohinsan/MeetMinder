<script lang="ts">

	import type {
		HTMLInputAttributes,
		HTMLInputTypeAttribute,
		HTMLTextareaAttributes,
	} from 'svelte/elements';

	// Generic type for the element
	type Element = 'input' | 'textarea' | 'type';

	// Input Elements interface
	interface InputComponentElements {
		input: HTMLInputAttributes;
		textarea: HTMLTextareaAttributes;
		type: HTMLInputTypeAttribute;
	}

	// Generic props for Input component
	type $$Props = InputComponentElements[Element] & {
		element: Element;
		variant?: 'solid' | 'outline' | 'text';
		typeis?: 'text' | 'tel' | 'email' | 'password' | 'search' | 'date' | 'file' | 'range' | 'radio' | 'checkbox' | 'select' | 'textarea';
		color?: 'red' | 'green' | 'cyan'; // Only one color now
		className?: string;
		value?: string;
		ariaPlaceholder?: string;
		ariaLabelledBy?: string;
		fullSize?: boolean;
		label?: string;
	};

	export let element: Element;
	export let variant: 'solid' | 'outline' | 'text' = 'solid';
	export let typeis: $$Props['typeis'] = 'text';
	export let size: 'sm' | 'md' | 'lg' = 'md';

	export let color: $$Props['color'] = 'cyan'; // Default color is 'cyan'
	export let ariaPlaceholder: $$Props['ariaPlaceholder'] = '';
	export let ariaLabelledBy: $$Props['ariaLabelledBy'] = '';
	export let className: $$Props['className'] = '';
	export let fullSize: boolean = false;
	export let value: $$Props['value'] = '';
	export let label: $$Props['label'] = '';
	let inputId = crypto.randomUUID();

	$: styles = `
			${className}
			${size in rounded ? rounded[size] : rounded['md']}
			${variant === 'outline' || variant === 'text' ? '' : `text-${color}`}
			${size in sizes ? sizes[size] : sizes['md']}
			${variant === 'text' ? '' : variant in variants ? variants[variant] : variants['solid']}
			${fullSize ? 'w-full' : ''}
`;

</script>

<div class="flex flex-col py-2">

	{#if label}
		<label for={inputId}>{label}</label>
	{/if}

	<svelte:element
		this={element}
		id={inputId}
		class={styles}
		value={value}
		on:on:submit
		type={typeis}
		aria-placeholder={ariaPlaceholder}
		aria-labelledby={ariaLabelledBy}
		role="textbox"
		tabindex="0"
		{...$$restProps}
	/>
</div>

<script context="module">
	const sizes = {
		sm: 'px-1 py-3',
		md: 'px-5 py-3',
		lg: 'px-7 py-5',
	};

	const rounded = {
		sm: 'rounded-sm',
		md: 'rounded-md',
		lg: 'rounded-lg',
	};

	const variants = {
		solid: 'text-white',
		outline: 'border border-solid border-current transition-colors duration-300 hover:border-red-600 ' +
			'active:border-red-900 focus:border-red-800 focus:outline-none',
		text: 'text-current',
	};
</script>