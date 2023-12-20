<!-- src/routes/availability.svelte -->
<script lang="ts">
	import Panel from '$lib/components/Panel.svelte';
</script>

<h1 class="text-3xl md:text-5xl sm:text-md mb-10">Event types</h1>
<div class="flex flex-wrap -mx-4 w-full w-[400px]">
    <div class="w-full sm:w-full md:w-1/2 lg:min-w-[400px] lg:w-1/3 px-4 mb-4">
		<Panel>
			<h2 class="text-2xl">30 Minute Meeting</h2>
			<p class="text-sm">30 mins, One-on-One</p>
			<div class="flex mt-2 justify-between">
				<button class="bg-blue-500 text-white px-3 rounded-md py-1 mr-2">Edit</button>
				<button class="bg-red-500 text-white px-3 rounded-md py-1">Delete</button>
			</div>
		</Panel>
	</div>
    <div class="w-full sm:w-full md:w-1/2 lg:min-w-[400px] lg:w-1/3 px-4 mb-4">
		<Panel>
			<h2 class="text-2xl">30 Minute Meeting</h2>
			<p class="text-sm">30 mins, One-on-One</p>
			<div class="flex mt-2 justify-between">
				<button class="bg-blue-500 text-white px-3 rounded-md py-1 mr-2">Edit</button>
				<button class="bg-red-500 text-white px-3 rounded-md py-1">Delete</button>
			</div>
		</Panel>
	</div>
    <div class="w-full sm:w-full md:w-1/2 lg:min-w-[400px] lg:w-1/3 px-4 mb-4">
		<Panel>
			<h2 class="text-2xl">30 Minute Meeting</h2>
			<p class="text-sm">30 mins, One-on-One</p>
			<div class="flex mt-2 justify-between">
				<button class="bg-blue-500 text-white px-3 rounded-md py-1 mr-2">Edit</button>
				<button class="bg-red-500 text-white px-3 rounded-md py-1">Delete</button>
			</div>
		</Panel>
	</div>
</div>
