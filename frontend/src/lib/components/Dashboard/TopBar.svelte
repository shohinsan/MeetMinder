<script lang="ts">
	import Magnifier from '$assets/Magnifier.svelte';
	import Bell from '$assets/Bell.svelte';
	import Button from '$components/system/Button.svelte';
	import { config } from '$lib/nav';
	import { page } from '$app/stores';

	let dropdown = false;

	let menuOpen = (): void => {
		dropdown = !dropdown;
	};

	let menuClose = (): void => {
		dropdown = false;
	};
</script>

<div
	class="sticky top-0 z-40 flex h-16 shrink-0 items-center gap-x-4 border-b border-gray-200 bg-white px-4 shadow-sm sm:gap-x-6 sm:px-6 lg:px-8"
>
	<div class="flex flex-1 gap-x-4 self-stretch lg:gap-x-6">
		<form class="relative flex flex-1" action="#" method="GET">
			<label for="search-field" class="sr-only">Search</label>
			<Magnifier />
			<input
				id="search-field"
				class="block h-full w-full border-0 py-0 pl-8 pr-0 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm"
				placeholder="Search..."
				type="search"
				name="search"
			/>
		</form>

		<div class="flex items-center gap-x-4 lg:gap-x-6">
			<button type="button" class="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500">
				<span class="sr-only">View notifications</span>
				<Bell width={24} height={24} />
			</button>

			<!-- Separator -->
			<div class="block h-6 w-px bg-gray-200" aria-hidden="true"></div>

			<!-- Profile dropdown -->
			<div class="relative inline-block text-left">
				<Button
					on:click={menuOpen}
					element="button"
					variant="outline"
					color="green"
					className="px-4 py-2 rounded-md transition-colors duration-300"
				>
					Option
				</Button>

				{#if dropdown}
					<div
						class="absolute right-0 z-10 mt-2 w-56 origin-top-right divide-y divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
						role="menu"
						aria-orientation="vertical"
						aria-labelledby="menu-button"
						tabindex="-1"
					>
						<div class="px-4 py-3" role="none">
							<p class="text-sm" role="none">Signed in as</p>
							<p class="truncate text-sm font-medium text-gray-900" role="none">tom@example.com</p>
						</div>

						<div class="py-1" role="none">
							{#each config.dropdownNavigation as { name, href }, index (index)}
								<a
									on:click={menuClose}
									href={`/app/${href === 'settings' ? '' : 'settings/'}${
										href === 'settings' ? '' : ''
									}${href}`}
									class={$page.url.pathname === `/app/${href}`
										? ' text-indigo-600 bg-gray-50 group flex gap-x-3 px-4 py-2 text-sm leading-6 font-semibold'
										: 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex px-4 py-2 gap-x-3  text-sm leading-6 font-semibold'}
								>
									{name}</a
								>
							{/each}
						</div>
						<div class="py-1" role="none">
							<!-- Sign out item -->
							{#each config.dropdownLogout as { name, href }, index (index)}
								<a
									{href}
									class="text-gray-700 block px-4 py-2 text-sm hover:bg-red-50 hover:text-red-900 font-semibold"
									role="menuitem"
									tabindex={-1}
									id={`menu-item-${index}`}>{name}</a
								>
							{/each}
						</div>
					</div>
				{/if}
			</div>
		</div>
	</div>
</div>
