<script lang="ts">
	export let data;

	const username = 'ina';
</script>

<div class="flex h-screen p-4 items-center sm:justify-center lg:justify-center">
	<div
		class="p-4 border rounded-lg w-full md:w-[800px] shadow-lg flex flex-col gap-2 bg-opacity-30 rounded relative"
	>
		<div class="w-36 aspect-square absolute -top-2 -right-2 rounded-sm overflow-hidden">
			<div class="absolute top-0 left-0 h-2 w-2 bg-amber-500"></div>
			<div class="absolute right-0 bottom-0 h-2 w-2 bg-amber-500"></div>
			<a
				href="###"
				class="py-1.5 w-[250px] bg-amber-300 text-amber-800 font-semibold uppercase text-xs tracking-wider block text-center absolute bottom-0 right-0 rotate-45 origin-bottom-right shadow-sm"
			>
				<p class="ml-12">Powered by <br /> Feature Link Here</p>
			</a>
		</div>

		<div class="divide-y md:divide-x">
			<div class=" w-full flex flex-col gap-2">
				<img src="###" alt="" class="h-16 w-16 flex-none rounded-full bg-gray-800 object-cover" />
				<p>Name</p>
				<p class="text-gray-400 w-1/2">
					Welcome to my scheduling page. Please follow the instructions to add an event to my
					calendar.
				</p>
			</div>
		</div>
        <div class="mt-1">Pass title here</div>

        
        <div class="border mt-5"></div>


		<div class="mt-5">Calendar component goes here</div>
        <!-- My guess is that if month changes url must be like this: http://localhost:5175/username/meeting-type?month=2023-12 and etc -->
	</div>
</div>
