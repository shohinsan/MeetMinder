<script lang="ts">
	import { page } from '$app/stores'
	import { dev } from '$app/environment'

	export let title: string

	$: title = $page?.status
		? `${$page.status} - ${
			$page.status === 404 ? 'Not Found' : $page.status === 500 ? 'Server Error' : ''
		}`
		: ''
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
	<main class="max-w-2xl p-6 rounded-md text-center">
		<button
			class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mx-auto"
			on:click={() => history.back()}
		>
			Go Back
		</button>

		{#if $page?.status === 404}
			<h1 class="mt-4 text-4xl font-bold text-red-600">Page Not Found</h1>
		{:else if $page?.status === 500}
			<h1 class="mt-4 text-4xl font-bold text-red-600">
				Whoops. We weren't expecting that. We're investigating the issue.
			</h1>
			{#if dev}
				<div class="mt-6">
					<h2 class="text-2xl font-semibold">Debugging Tips:</h2>
					<ul class="list-disc list-inside">
						<li>Remove unused import</li>
						<li>Run your database and server</li>
					</ul>
				</div>
			{/if}
		{/if}
	</main>
</div>
