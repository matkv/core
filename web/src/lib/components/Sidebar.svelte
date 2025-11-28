<script lang="ts">
	// simple sidebar/navigation
	export let open = false;
	const handleBackdropClick = () => {
		open = false;
	};

	import Menu from '$lib/components/Menu.svelte';
</script>

<!-- Desktop sidebar -->
<aside class="w-64 bg-neutral-950 border-r border-neutral-800 hidden md:flex flex-col">
	<div class="p-4 text-xl font-semibold text-neutral-100">Core</div>
	<div class="px-2 pb-4">
		<Menu />
	</div>
	<div class="mt-auto p-4 text-xs text-neutral-500 border-t border-neutral-800">
		<div>v{import.meta.env.VITE_APP_VERSION ?? '0.0.0'}</div>
	</div>
</aside>

<!-- Mobile sidebar -->
{#if open}
	<div class="fixed inset-0 z-40 flex md:hidden">
		<button
			class="fixed inset-0 bg-black/60 cursor-pointer"
			type="button"
			aria-label="Close sidebar"
			on:click={handleBackdropClick}
		></button>
		<aside class="relative w-64 bg-neutral-950 border-r border-neutral-800 flex flex-col">
			<div class="p-4 text-xl font-semibold text-neutral-100">Core</div>
			<div class="px-2 pb-4">
				<Menu />
			</div>
			<div class="mt-auto p-4 text-xs text-neutral-500 border-t border-neutral-800">
				<div>v{import.meta.env.VITE_APP_VERSION ?? '0.0.0'}</div>
			</div>
		</aside>
	</div>
{/if}
