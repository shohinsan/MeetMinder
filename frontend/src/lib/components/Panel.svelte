<script lang="ts">
	export let title: string = '';
	export let width: string = 'w-full';
</script>

<div class="flex bg-opacity-50 backdrop-blur-md">
	<div class="bg-white bg-opacity-30 backdrop-blur-md {width} p-4 border rounded-lg shadow-lg flex flex-col gap-2">
		<div class="text-2xl text-center font-bold">
			{title}
		</div>
		<slot />
	</div>
</div>
