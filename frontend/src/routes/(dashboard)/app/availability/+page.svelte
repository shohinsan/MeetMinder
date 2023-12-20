<script lang="ts">
	import Settings from '$assets/Settings.svelte';
	import Input from '$components/system/Input.svelte';
	import Panel from '$lib/components/Panel.svelte';
	import Tabs from '$lib/components/Tabs.svelte';
  
	let days: Record<string, boolean> = {
	  Monday: false,
	  Tuesday: false,
	  Wednesday: false,
	  Thursday: false,
	  Friday: false,
	  Saturday: false,
	  Sunday: false
	};
  
	let items = ['Schedules', 'Holidays'];
	let activeItem = 'Schedules';
  
	const tabChange = (e: CustomEvent<string>) => {
	  activeItem = e.detail;
	};
  </script>
  
  <h1 class="text-3xl md:text-5xl sm:text-md">Availability</h1>
  
  <Tabs {activeItem} {items} on:tabChange={tabChange} />
  
  {#if activeItem === 'Schedules'}
	<h1 class="text-2xl md:text-4xl my-10">Work Schedule</h1>
  
	<div class="my-10">
	  <Panel>
		<div class="flex items-start justify-between">
		  <div class="flex flex-wrap">
			<div class="mr-5 bg-red-100 min-w-[230px] rounded-md p-4">
			  <h1 class="text-lg font-semibold mb-2">Active On</h1>
			  <h1 class="text-lg font-semibold">Event Type Dropdown Here</h1>
			</div>
			<div class="bg-red-100 rounded-md min-w-[230px] p-4">
			  <h1 class="text-lg font-semibold mb-2">Time Zone</h1>
			  <h1 class="text-lg font-semibold">Time Zone Dropdown Here</h1>
			</div>
		  </div>
  
		  <div class="ml-5">
			<Settings width={24} height={24} />
		  </div>
		</div>
  
		<div class="border-t border-gray-300 mt-5"></div>
  
		<h1 class="text-lg md:text-xl my-10">Weekly hours</h1>
		<!-- Monday -->
		<div class="flex flex-wrap flex items-center border border-gray-200 rounded">
		  <div class="mr-5 min-w-[150px] rounded-md p-4">
			<div class="mr-4">
			  <input bind:checked={days.Monday} type="checkbox" value="" id="monday" name="monday" class="mr-2"/>
			  <label for="monday" class="text-sm font-medium text-gray-900">Monday</label>
			</div>
		  </div>
		  <div class="rounded-md min-w-[230px] p-4">
			{#if days.Monday}
			  <input type="time" name="start" class="border rounded px-2" aria-labelledby="monday" />
			  <span class="mx-1">to</span>
			  <input type="time" name="end" class="border rounded px-2" aria-labelledby="monday" />
			{:else}
			  <span>Unavailable</span>
			{/if}
		  </div>
		</div>

		<!-- Tuesday -->
	<div class="flex flex-wrap flex items-center border border-gray-200 rounded">
		<div class="mr-5 min-w-[150px] rounded-md p-4">
			<div class="mr-4">
				<input bind:checked={days.Tuesday} type="checkbox" value="" id="tuesday" name="tuesday" class="mr-2"/>
				<label for="tuesday" class="text-sm font-medium text-gray-900">Tuesday</label>
			</div>
		</div>
		<div class="rounded-md min-w-[230px] p-4">
			{#if days.Tuesday}
				<input type="time" name="start" class="border rounded px-2" aria-labelledby="tuesday" />
				<span class="mx-1">to</span>
				<input type="time" name="end" class="border rounded px-2" aria-labelledby="tuesday" />
			{:else}
				<span>Unavailable</span>
			{/if}
		</div>
	</div>

	<!-- Wednesday -->
	<div class="flex flex-wrap flex items-center border border-gray-200 rounded">
		<div class="mr-5 min-w-[150px] rounded-md p-4">
			<div class="mr-4">
				<input bind:checked={days.Wednesday} type="checkbox" value="" id="wednesday" name="wednesday" class="mr-2"/>
				<label for="wednesday" class="text-sm font-medium text-gray-900">Wednesday</label>
			</div>
		</div>
		<div class="rounded-md min-w-[230px] p-4">
			{#if days.Wednesday}
				<input type="time" name="start" class="border rounded px-2" aria-labelledby="wednesday" />
				<span class="mx-1">to</span>
				<input type="time" name="end" class="border rounded px-2" aria-labelledby="wednesday" />
			{:else}
				<span>Unavailable</span>
			{/if}
		</div>
	</div>

	<!-- Thursday -->
	<div class="flex flex-wrap flex items-center border border-gray-200 rounded">
		<div class="mr-5 min-w-[150px] rounded-md p-4">
			<div class="mr-4">
				<input bind:checked={days.Thursday} type="checkbox" value="" id="thursday" name="thursday" class="mr-2"/>
				<label for="thursday" class="text-sm font-medium text-gray-900">Thursday</label>
			</div>
		</div>
		<div class="rounded-md min-w-[230px] p-4">
			{#if days.Thursday}
				<input type="time" name="start" class="border rounded px-2" aria-labelledby="thursday" />
				<span class="mx-1">to</span>
				<input type="time" name="end" class="border rounded px-2" aria-labelledby="thursday" />
			{:else}
				<span>Unavailable</span>
			{/if}
		</div>
	</div>

	<!-- Friday -->
	<div class="flex flex-wrap flex items-center border border-gray-200 rounded">
		<div class="mr-5 min-w-[150px] rounded-md p-4">
			<div class="mr-4">
				<input bind:checked={days.Friday} type="checkbox" value="" id="friday" name="friday" class="mr-2"/>
				<label for="friday" class="text-sm font-medium text-gray-900">Friday</label>
			</div>
		</div>
		<div class="rounded-md min-w-[230px] p-4">
			{#if days.Friday}
				<input type="time" name="start" class="border rounded px-2" aria-labelledby="friday" />
				<span class="mx-1">to</span>
				<input type="time" name="end" class="border rounded px-2" aria-labelledby="friday" />
			{:else}
				<span>Unavailable</span>
			{/if}
		</div>
	</div>

	<!-- Saturday -->
	<div class="flex flex-wrap flex items-center border border-gray-200 rounded">
		<div class="mr-5 min-w-[150px] rounded-md p-4">
			<div class="mr-4">
				<input bind:checked={days.Saturday} type="checkbox" value="" id="saturday" name="saturday" class="mr-2"/>
				<label for="saturday" class="text-sm font-medium text-gray-900">Saturday</label>
			</div>
		</div>
		<div class="rounded-md min-w-[230px] p-4">
			{#if days.Saturday}
				<input type="time" name="start" class="border rounded px-2" aria-labelledby="saturday" />
				<span class="mx-1">to</span>
				<input type="time" name="end" class="border rounded px-2" aria-labelledby="saturday" />
			{:else}
				<span>Unavailable</span>
			{/if}
		</div>
	</div>

	<!-- Sunday -->
	<div class="flex flex-wrap flex items-center border border-gray-200 rounded">
		<div class="mr-5 min-w-[150px] rounded-md p-4">
			<div class="mr-4">
				<input bind:checked={days.Sunday} type="checkbox" value="" id="sunday" name="sunday" class="mr-2"/>
				<label for="sunday" class="text-sm font-medium text-gray-900">Sunday</label>
			</div>
		</div>
		<div class="rounded-md min-w-[230px] p-4">
			{#if days.Sunday}
				<input type="time" name="start" class="border rounded px-2" aria-labelledby="sunday" />
				<span class="mx-1">to</span>
				<input type="time" name="end" class="border rounded px-2" aria-labelledby="sunday" />
			{:else}
				<span>Unavailable</span>
			{/if}
		</div>
	</div>


  
	  </Panel>
	</div>
  {:else if activeItem === 'Holidays'}
	<!-- Holidays section -->
	<p class="my-10">Holidays Tab</p>
  {/if}
  

  <!-- Need later -->
<!-- <h1 class="text-md md:text-lg my-10">Third Dropdown</h1>

<select>
	<option>Option 1</option>
	<option>Option 2</option>
	<option>Option 3</option>
</select>

<select>
	<option>Option A</option>
	<option>Option B</option>
	<option>Option C</option>
</select>

<select>
	<option>Option X</option>
	<option>Option Y</option>
	<option>Option Z</option>
</select> -->



	
	