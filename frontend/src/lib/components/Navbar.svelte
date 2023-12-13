<script lang="ts">
	import { page } from '$app/stores';
	import logo from '../../images/svelte-logo.svg';
	import Menu from '$lib/icons/Menu.svelte';
	import { fly } from 'svelte/transition';
	import type { NavItem } from '../../app';

	const navItems: NavItem[] = [
		{
			name: 'Home',
			href: '/'
		},
		{
			name: 'About',
			href: '/about'
		},
		{
			name: 'Blogs',
			href: '/blog'
		},
		{
			name: 'Pricing',
			href: '/pricing'
		}
	];

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
			{#each navItems as { name, href }}
				<li aria-current={$page.url.pathname === href ? 'page' : undefined} class="inline-block py-2 px-4">
					<a {href} class={($page.url.pathname === href) ? 'text-blue-900 bg-amber-100 p-4 rounded-xl' : 'text-blue-500'}>{name}</a>
				</li>
			{/each}
		</ul>
	</nav>
	<div class="flex items-center space-x-4">
		<a href="login" class="text-blue-500 hover:text-blue-700">Login</a>
		<div class="lg:hidden cursor-pointer text-xl ">
			<Menu {hamburger} {menuOpen} />
		</div>
	</div>
</header>

{#if hamburger}
	<nav transition:fly={{ y: -200, duration: 400 }} class="bg-white p-4">
		<ul class="flex flex-col space-y-4 justify-between ">
			{#each navItems as { name, href }}
				<li aria-current={$page.url.pathname === href ? 'page' : undefined} class="inline-block py-2 px-4">
					<a {href} class={($page.url.pathname === href) ? 'text-blue-900' : 'text-blue-500'}
						 on:click={() => $page.url.pathname === href && menuClose()}>
						{name}
					</a>
				</li>
			{/each}
		</ul>
	</nav>
{/if}

