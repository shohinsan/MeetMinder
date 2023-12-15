<script>
	export let width = 64;
	export let height = 64;
	export let stroke = 'currentColor';
	export let fill = 'none';
</script>

<svg
    {width}
    {height}
	xmlns="http://www.w3.org/2000/svg"
	{fill}
	viewBox="0 0 24 24"
	stroke-width="1.5"
	{stroke}
	class="w-6 h-6"
>
	<path
		stroke-linecap="round"
		stroke-linejoin="round"
		d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"
	/>
</svg>
