<script lang="ts">
	import { enhance } from '$app/forms';
	import type { ActionData } from './$types';
	import Panel from '$components/Panel.svelte';
	import Input from '$components/system/Input.svelte';
	import Button from '$components/system/Button.svelte';

	export let form: ActionData


	
</script>

<div class="h-screen flex items-center justify-center flex flex-col">
	<Panel title="Register">
		<form method="POST" use:enhance>
			<!-- Step 0 Starts -->
			<Input
				name="username"
				placeholder="Username"
				element="input"
				variant="outline"
				typeis="text"
				color="red"
				ariaPlaceholder="Enter username"
				ariaLabelledBy="usernameLabel"
				label="Username"
				fullSize={true}
			/>

			{#if form?.errors?.usernameMissing}
				<p class="text-red-500">{form.errors.usernameMissing}</p>
			{/if}

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
				fullSize={true}
			/>

			{#if form && form.errors}
				{#if form.errors.emailMissing}
					<p class="text-red-500">{form.errors.emailMissing}</p>
				{/if}
				{#if form.errors.invalidEmail}
					<p class="text-red-500">{form.errors.invalidEmail}</p>
				{/if}
			{/if}
			<!-- Step 0 Ends -->

			<!-- Step 1 Starts -->
			<Input
				name="name"
				placeholder="Name"
				element="input"
				variant="outline"
				typeis="text"
				color="red"
				ariaPlaceholder="Enter name"
				ariaLabelledBy="nameLabel"
				label="Name"
				fullSize={true}
			/>

			{#if form?.errors?.nameMissing}
				<p class="text-red-500">{form.errors.nameMissing}</p>
			{/if}
			<!-- Step 1 Ends -->

			<!-- Step 2 Starts -->
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
				fullSize={true}
			/>

			{#if form?.errors?.passwordMissing}
				<p class="text-red-500">{form.errors.passwordMissing}</p>
			{/if}

			<Input
				name="confirm_password"
				placeholder="Confirm Password"
				element="input"
				variant="outline"
				typeis="password"
				color="red"
				ariaPlaceholder="Enter confirm password"
				ariaLabelledBy="confirmPasswordLabel"
				label="Confirm Password"
				fullSize={true}
			/>

			{#if form && form.errors}
				{#if form.errors.confirmPasswordMissing}
					<p class="text-red-500">{form.errors.confirmPasswordMissing}</p>
				{/if}
				{#if form.errors.passwordMismatch}
					<p class="text-red-500">{form.errors.passwordMismatch}</p>
				{/if}
			{/if}

			<Button
				className="mt-4"
				element="button"
				variant="solid"
				color="green"
				ariaLabel="Press me!"
				ariaPressed={true}
				fullSize={true}
			>
				Submit
			</Button>
			<!-- Step 2 Ends -->
		</form>

		<p class="pt-4 text-center">
			Already have an account? <a class="text-blue-500" href="/login">Login</a>
		</p>
	</Panel>
</div>
