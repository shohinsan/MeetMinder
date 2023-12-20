<!-- src/routes/availability.svelte -->
<script lang="ts">
	import Panel from '$lib/components/Panel.svelte';
	import Tabs from '$lib/components/Tabs.svelte';

	let items = ['Upcoming', 'Pending', 'Past'];
	let activeItem = 'Upcoming';

	const tabChange = (e: CustomEvent<string>) => {
		activeItem = e.detail;
	};
</script>

<h1 class="text-3xl md:text-5xl sm:text-md mb-10">Scheduled events</h1>

<Tabs {activeItem} {items} on:tabChange={tabChange} />
{#if activeItem === 'Upcoming'}
	<div class="mt-10">
		<!-- Add all sorts of ssr search, filter functionality here -->
		<Panel>
			Hey Upcoming List
		</Panel>
	</div>
{:else if activeItem === 'Pending'}
<div class="mt-10">
	<Panel>
		Hey Pending List
	</Panel>
</div>
{:else if activeItem === 'Past'}
<div class="mt-10">
	<Panel>
		Hey Past List
	</Panel>
</div>
{/if}
