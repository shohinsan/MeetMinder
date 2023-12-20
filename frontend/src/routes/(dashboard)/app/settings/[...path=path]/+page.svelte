<script>
	import Button from '$components/system/Button.svelte';
	import Input from '$components/system/Input.svelte';

	const sections = [
		{
			name: 'Personal Information',
			description: 'Use a permanent address where you can receive mail.'
		},
		{
			name: 'Delete Account',
			description:
				'No longer want to use our service? You can delete your account here. This action is not reversible. All information related to this account will be deleted permanently.'
		}
	];
</script>

<div class="divide-y">
	{#each sections as { name, description }}
		<div
			class="grid max-w-7xl grid-cols-1 gap-x-8 gap-y-10 px-4 py-16 sm:px-6 md:grid-cols-3 lg:px-8"
		>
			<div>
				<h2 class="text-base font-semibold leading-7">{name}</h2>
				<p class="mt-1 text-sm leading-6 text-gray-400">{description}</p>
			</div>

			{#if name === 'Personal Information'}
				<form class="md:col-span-2">
					<div class="grid grid-cols-1 gap-x-6 gap-y-8 sm:max-w-xl sm:grid-cols-6">
						<div class="col-span-full flex items-center gap-x-8">
							<img
								src="###"
								alt=""
								class="h-24 w-24 flex-none rounded-lg bg-gray-800 object-cover"
							/>
							<div>
                                
								<button
									type="button"
									class="rounded-md bg-white/10 px-3 py-2 text-sm font-semibold shadow-sm hover:bg-white/20"
									>Change avatar</button
								>
								<p class="mt-2 text-xs leading-5 text-gray-400">JPG, GIF or PNG. 1MB max.</p>
							</div>
						</div>

						<div class="sm:col-span-3">
							<div class="mt-2">
								<Input
									name="first-name"
									placeholder="First name"
									element="input"
									variant="outline"
									typeis="email"
									color="red"
									ariaPlaceholder="Enter first name"
									ariaLabelledBy="firstNameLabel"
									label="First Name"
									fullSize={true}
								/>
							</div>
						</div>

						<div class="sm:col-span-3">
							<div class="mt-2">
								<Input
									name="last-name"
									placeholder="Last name"
									element="input"
									variant="outline"
									typeis="email"
									color="red"
									ariaPlaceholder="Enter last name"
									ariaLabelledBy="lastNameLabel"
									label="Last Name"
									fullSize={true}
								/>
							</div>
						</div>

						<div class="col-span-full">
							<div class="mt-2">
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
							</div>
						</div>

						<div class="col-span-full">
							<div class="mt-2">
								<div class="mt-2">
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
								</div>
							</div>
						</div>
					</div>

					<div class="mt-8 flex">
						<Button element="button" color="green" variant="solid">Save</Button>
					</div>
				</form>
			{/if}

			{#if name === 'Delete Account'}
				<form class="flex items-start md:col-span-2">
					<Button element="button" color="red" variant="solid">Delete Account</Button>
				</form>
			{/if}
		</div>
	{/each}
</div>

<!-- <div class="col-span-full">
							<label for="timezone" class="block text-sm font-medium leading-6">Timezone</label>
							<div class="mt-2">
								<select
									id="timezone"
									name="timezone"
									class="block w-full rounded-md border-0 bg-white/5 py-1.5 shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-indigo-500 sm:text-sm sm:leading-6 [&_*]:text-black"
								>
									<option>Pacific Standard Time</option>
									<option>Eastern Standard Time</option>
									<option>Greenwich Mean Time</option>
								</select>
							</div>
						</div> -->
