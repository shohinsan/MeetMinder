<script>
	import Navbar from '$components/Navbar.svelte'
	import Metadata from '$components/Metadata.svelte';
	import '$styles/app.css';

	import { onNavigate } from '$app/navigation';

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

<Metadata />
<Navbar />

<main>
	<slot />
</main>

<style lang="postcss">
    main {
        min-height: 100vh;
    }
</style>

