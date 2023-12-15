<script lang="ts">

	import logo from '../../images/Svelte.svg';
	import Menu from '$lib/icons/Menu.svelte';
	import Button from './system/Button.svelte';
	import { config } from '$lib/nav';
	import { fly } from 'svelte/transition';
	import { page } from '$app/stores';
	
	let hamburger = false;

	export let menuOpen = (): void => {
		hamburger = !hamburger;
	};

	let menuClose = (): void => {
		hamburger = false;
	};

</script>

<header class="bg-white p-4 shadow flex justify-between items-center">
	<div class="flex justify-between items-center">
		<a href="/" class="flex items-center">
			<img src={logo} alt="SvelteKit" class="w-12 h-12" />
		</a>
	</div>
	<nav class="ml-4 hidden sm:flex items-center">
		<ul class="flex text-center items-center space-x-4">
			{#each config.headerNavigation as { name, href }}
				<li aria-current={$page.url.pathname === href ? 'page' : undefined} class="inline-block py-2 px-4">
					<a {href} class={($page.url.pathname === href) ? 'text-blue-900 bg-amber-100 p-4 rounded-xl' : 'text-blue-500'}>{name}</a>
				</li>
			{/each}
		</ul>
	</nav>
	<div class="flex items-center space-x-4">
		<Button element="a" variant="solid" color="green" href="login" className="px-4 py-2 rounded-md text-white bg-green-600 hover:bg-green-700 transition-colors duration-300 shadow-md
		">Login</Button>
		<div class="lg:hidden cursor-pointer text-xl ">
			<Menu {hamburger} {menuOpen} />
		</div>
	</div>
</header>

{#if hamburger}
  <nav transition:fly={{ y: -200, duration: 400 }} class="bg-white p-4">
    <ul class="flex flex-col space-y-4 justify-between">
      {#each config.headerNavigation as { name, href }}
        <li aria-current={$page.url.pathname === href ? 'page' : undefined} class="inline-block py-2 px-4">
          <a {href} class:active={$page.url.pathname === href} class={($page.url.pathname === href) ? 'block w-full text-blue-900 bg-amber-100 p-4 rounded-md' : 'block w-full text-blue-500'} on:click={() => $page.url.pathname === href && menuClose()}>
            {name}
          </a>
        </li>
      {/each}
    </ul>
  </nav>
{/if}