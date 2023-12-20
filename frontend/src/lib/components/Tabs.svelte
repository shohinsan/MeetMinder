<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    
    let dispatch = createEventDispatcher();
    export let items: string[] = [];
    export let activeItem: string = '';
  </script>
  
  <div class="sm:flex list-none">
    <div class="border-b border-gray-200 w-full flex">
      {#each items as item}
        <li class="-mb-px">
          <button
            on:click={() => dispatch('tabChange', item)}
            class={`
              border-transparent text-gray-500 mr-5 hover:text-gray-700
              whitespace-nowrap border-b-2 py-4 px-1 text-sm font-medium
              ${item === activeItem ? 'border-b-0' : ''}
            `}
          >
            {item}
          </button>
          {#if item === activeItem}
            <span class="block h-1 bg-blue-500" />
          {/if}
        </li>
      {/each}
    </div>
  </div>
  