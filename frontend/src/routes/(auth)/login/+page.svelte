<script lang="ts">
	import { enhance } from '$app/forms';
	import type { ActionData } from './$types';
	import Input from '$components/system/Input.svelte';
	import Button from '$components/system/Button.svelte';
	import Panel from '$components/Panel.svelte';
	export let form: ActionData;

</script>

<div class="h-screen flex items-center justify-center flex flex-col">
	<Panel title="Login">
		<form method="POST" use:enhance>
			<Input
				name="email"
				placeholder="Email"
				element="input"
				variant="outline"
				typeis="email"
				color="red"
				ariaPlaceholder="Enter email"
				ariaLabelledBy="emailLabel"
				label="Email"
				fullSize={true} />
	
			{#if form?.errors?.emailMissing}
				<p class="text-red-500">{form.errors.emailMissing}</p>
			{/if}
	
			<Input
				name="password"
				placeholder="Password"
				element="input"
				variant="outline"
				typeis="password"
				color="red"
				ariaPlaceholder="Enter password"
				ariaLabelledBy="passwordLabel"
				label="Password"
				fullSize={true} />
	
			{#if form?.errors?.passwordMissing}
				<p class="text-red-500">{form.errors.passwordMissing}</p>
			{/if}
	
			<Button
				element="button"
				variant="solid"
				color="green"
				ariaLabel="Press me!"
				ariaPressed={true}
				fullSize={true}
			>
				Submit
			</Button>
		</form>
		<p class="pt-4 text-center">Don't have an account? <a class="text-blue-500" href="/register">Register</a></p>
	</Panel>
	
</div>
