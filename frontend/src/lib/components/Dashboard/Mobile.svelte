<script lang="ts">
    import { page } from '$app/stores';
    import Settings from '$assets/Settings.svelte';
    import { config } from '$lib/nav';
    import Clock from '$assets/Clock.svelte';
    import Calendar from '$assets/Calendar.svelte';
    import Event from '$assets/Event.svelte';
    import CreateEvent from '$assets/CreateEvent.svelte';


    let basePath = "/app/";


</script>

<div class="fixed lg:hidden z-50 w-full h-16 -translate-x-1/2 bg-white border border-gray-200 bottom-0 left-1/2">
    <div class="grid h-full grid-cols-5 mx-auto">

        {#each config.mobileNavigation as { name, href }, i}
            {#if i % 4 === 2}
                <div class="inline-flex flex-col items-center justify-center group">
                    <div class="flex items-center justify-center">
                        <a href={href} class="inline-flex items-center justify-center font-medium rounded-full hover:bg-blue-600 group focus:ring-4 focus:ring-blue-300 focus:outline-none">
                            <CreateEvent width={64} height={64}/>
                            <span class="sr-only">New item</span>
                        </a>
                    </div>
                </div>
            {/if}

            <a href={`${basePath}${href}`}
               data-tooltip-target={'tooltip-' + `${basePath}${href}`}
               class={`
                   inline-flex flex-col items-center justify-center hover:bg-gray-50 group
                   ${
                       $page.url.pathname === `${basePath}${href}`
                           ? ' bg-gray-200 group flex gap-x-3 p-2 text-sm leading-6 font-semibold'
                           : 'text-gray-700 hover:text-indigo-600 hover:bg-gray-50 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold'
                   }
               `}
            >
                {#if href === 'events'}
                    <Event width={24} height={24}/>
                {:else if href === 'scheduled'}
                    <Calendar width={24} height={24}/>
                {:else if href === 'availability'}
                    <Clock width={24} height={24}/>
                {:else if href === 'settings'}
                    <Settings width={24} height={24}/>
                {/if}
                
                <span class="sr-only">{name}</span>
            </a>

            <div id={'tooltip-' + `${basePath}${href}`} role="tooltip" class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 shadow-sm opacity-0 tooltip">
                {name}
                <div class="tooltip-arrow" data-popper-arrow></div>
            </div>
        {/each}
    </div>
</div>
