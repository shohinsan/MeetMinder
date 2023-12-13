<script lang="ts">
	export let title: string = '';

</script>

<div class="flex items-center justify-center bg-opacity-50 backdrop-blur-md">
	<div class="bg-white bg-opacity-30 backdrop-blur-md w-[400px] p-4 border rounded-lg shadow-lg flex flex-col gap-2">
		<div class="text-2xl text-center font-bold">
			{title}
		</div>
		<slot />
	</div>
</div>