<script lang="ts">
	import { page } from '$app/stores';
	import Logo from '$assets/Logo.svelte';
	import { config } from '$lib/nav';

	let dropdown = false;

	let menuOpen = (): void => {
		dropdown = !dropdown;
	};
	
	let menuClose = (): void => {
		dropdown = false;
	};

	let basePath = "/app/";

</script>

<div class="hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
	<div
		class="flex grow flex-col gap-y-5 overflow-y-auto border-r border-gray-200 bg-white px-6 pb-4"
	>
		<div class="flex h-16 shrink-0 items-center">
			<Logo width={32} height={32} />
		</div>
		<nav class="flex flex-1 flex-col">
			<ul role="list" class="flex flex-1 flex-col gap-y-7">
				<li>
					<ul role="list" class="-mx-2 space-y-1">
						{#each config.sidebarTopNavigation as { name, href } (name)}
							<li aria-current={$page.url.pathname === `${basePath}${href}` ? 'page' : undefined}>
								<a
									href={`${basePath}${href}`}
									class={$page.url.pathname === `${basePath}${href}`
										? ' text-indigo-600 bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
										: 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'}
								>
									<span class="truncate">{name}</span>
								</a>
							</li>
						{/each}
					</ul>
				</li>
				<li>
					{#each Object.entries(config.sidebarNavigation) as [header, items]}
						<div class="text-xs font-semibold leading-6 text-gray-400">{header}</div>
						<ul role="list" class="-mx-2 mt-2 space-y-1">
							{#each items as { name, href } (name)}
								<li aria-current={$page.url.pathname === `${basePath}${href}` ? 'page' : undefined}>
									<a
										href={`${basePath}${href}`}
										class={$page.url.pathname === `${basePath}${href}`
											? ' text-indigo-600 bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
											: 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'}
									>
										<span class="truncate">{name}</span>
									</a>
								</li>
							{/each}
						</ul>
					{/each}
				</li>
				<li class="mt-auto">
					<ul role="list" class="-mx-2 mt-2 space-y-1">
						{#each config.sidebarBottomNavigation as { name, href } (name)}
							<li aria-current={$page.url.pathname === `${basePath}${href}` ? 'page' : undefined}>
								<a
									href={`${basePath}${href}`}
									class={$page.url.pathname === `${basePath}${href}`
										? ' text-indigo-600 bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
										: 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'}
								>
									<span class="truncate">{name}</span>
								</a>
							</li>
						{/each}
					</ul>
				</li>
			</ul>
		</nav>
	</div>
</div>
