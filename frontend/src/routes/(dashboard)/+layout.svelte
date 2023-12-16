<script>
	import '$styles/app.css';
	import { onNavigate } from '$app/navigation';
	import Desktop from '$components/Dashboard/Desktop.svelte';
	import Mobile from '$components/Dashboard/Mobile.svelte';
	import TopBar from '$components/Dashboard/TopBar.svelte';
	
	onNavigate((navigation) => {
		if (!document.startViewTransition) return;
		return new Promise((resolve) => {
			document.startViewTransition(async () => {
				resolve();
				await navigation.complete;
			});
		});
	});

</script>

<Desktop />

<div class="lg:pl-72">
	<TopBar />
	<main class="py-10">
		<div class="px-4 sm:px-6 lg:px-8">
			<slot />

		</div>
	</main>
</div>

<Mobile />

<style lang="postcss">
    main {
        min-height: 100vh;
    }
</style>