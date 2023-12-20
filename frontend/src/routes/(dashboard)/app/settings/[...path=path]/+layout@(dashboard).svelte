<script>
	import { page } from '$app/stores';
	import { config } from '$lib/nav';

	console.log($page.url.pathname);
</script>

<nav class="flex overflow-x-auto">
	<ul
		role="list"
		class="flex min-w-full flex-none gap-x-6 px-4 text-sm font-semibold leading-6 text-gray-400 sm:px-6 lg:px-8"
	>
		{#each config.settingsNavigation as { name, href }}
			<li>
				<a
					href={`/app/${href === 'settings' ? '' : 'settings/'}${
						href === 'settings' ? '' : ''
					}${href}`}
					class={$page.url.pathname === `/app/${href}` ||
					($page.url.pathname === '/app/settings' && href === 'settings')
						? ' text-indigo-600 bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
						: 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'}
				>
					<span class="truncate">{name}</span>
				</a>
			</li>
		{/each}
	</ul>
</nav>

<slot />

<!-- <script>
  import { config } from '$lib/nav';
  import { page } from '$app/stores';

  console.log($page.url.pathname);
</script>

<nav class="flex overflow-x-auto">
  <ul
      role="list"
      class="flex min-w-full flex-none gap-x-6 px-4 text-sm font-semibold leading-6 text-gray-400 sm:px-6 lg:px-8"
  >
      {#each config.settingsNavigation as { name, href }}
          <li aria-current={$page.url.pathname === `/app/${href}` || ($page.url.pathname === '/app/settings' && href === 'settings') ? 'page' : undefined}>
              <a
              href={`/app/${href === 'settings' ? '' : 'settings/'}${href === 'settings' ? '' : ''}${href}`}
              class={$page.url.pathname === `/app/${href}` || ($page.url.pathname === '/app/settings' && href === 'settings')
                  ? ' text-indigo-600 bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
                  : 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'}
          >
              <span class="truncate">{name}</span>
          </a>
          </li>
      {/each}
  </ul>
</nav>

<slot /> -->
