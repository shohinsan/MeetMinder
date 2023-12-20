<script lang="ts">
    const cards = [
        {
            title: '30 min meeting',
            description: 'Card description or any other content you want to include for Card 1.'
        },
        {
            title: '45 min meeting',
            description: 'Card description or any other content you want to include for Card 2.'
        },
        {
            title: '1 hour meeting',
            description: 'Card description or any other content you want to include for Card 2.'
        }
        // Add more cards as needed
    ];

    const username = 'ina'; // Replace with your actual username
</script>

<div class="flex h-screen p-4 items-center sm:justify-center lg:justify-center">
    <div class="p-4 border rounded-lg w-full md:w-[800px] shadow-lg flex flex-col gap-2 bg-opacity-30 rounded relative">
        <div class="w-36 aspect-square absolute -top-2 -right-2 rounded-sm overflow-hidden">
            <div class="absolute top-0 left-0 h-2 w-2 bg-amber-500"></div>
            <div class="absolute right-0 bottom-0 h-2 w-2 bg-amber-500"></div>
            <a href="###" class="py-1.5 w-[250px] bg-amber-300 text-amber-800 font-semibold uppercase text-xs tracking-wider block text-center absolute bottom-0 right-0 rotate-45 origin-bottom-right shadow-sm">
                <p class="ml-12">Powered by <br /> Feature Link Here</p>
            </a>
        </div>

        <div class="text-center items-center w-full mb-20 flex flex-col gap-2">
            <img src="###" alt="" class="h-16 w-16 flex-none rounded-full bg-gray-800 object-cover" />
            <p>Name</p>
            <p class="text-gray-400 w-1/2">
                Welcome to my scheduling page. Please follow the instructions to add an event to my
                calendar.
            </p>
        </div>

        <div class="flex flex-wrap gap-4 lg:ml-10 justify-center md:justify-start">
            {#each cards as card (card.title)}
                <div class="p-4 w-[300px] bg-gray-100 shadow-md rounded-md m-4">
                    <h2 class="text-2xl font-semibold mb-4 flex items-center justify-between">
                        <div class="flex items-center"> <!-- Wrap the icon and title in a flex container -->
                            <div class="rounded-full bg-amber-500 h-4 w-4 mr-2"></div> <!-- Add margin to the icon -->
                            <a href={`/${username}/${encodeURIComponent(card.title.replace(/\s+/g, '-').toLowerCase())}?`}>
                                <span>{card.title}</span>
                            </a>
                            <!-- Use a span for the title to keep it inline with the icon -->
                        </div>

                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                        </svg>
                    </h2>

                    <p class="text-gray-600">{card.description}</p>
                </div>
            {/each}
        </div>
    </div>
</div>
