<script lang="ts">
    import { enhance } from '$app/forms';
</script>

<div class="flex flex-wrap h-screen p-4 items-center justify-center">
    <div class="p-4 border bg-white rounded-lg items-center w-[1000px] shadow-lg flex flex-col gap-2 bg-opacity-30 rounded relative">
        <form method="POST" use:enhance>
            <div class="mb-96">
                This is your username page
            </div>
        </form>
        <!-- Powered by text in the bottom-right corner -->
        <div class="absolute bottom-4 right-4 text-gray-500">
            Powered by <a href="http://localhost:5173/">meetminder</a>
        </div>
    </div>
</div>
