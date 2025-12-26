<script lang="ts">
	// simple sidebar/navigation
	export let collapsed = false;

	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher<{ toggleCollapsed: void }>();

	const handleToggleCollapsed = () => {
		dispatch('toggleCollapsed');
	};

	import Menu from '$lib/components/Menu.svelte';
</script>

<aside
	class={`flex overflow-hidden border-r border-neutral-800/80 bg-neutral-950/50 backdrop-blur supports-backdrop-filter:bg-neutral-950/40 ${collapsed ? 'w-16' : 'w-64'} flex-col transition-[width] duration-200 motion-reduce:transition-none`}
>
	<div
		class={`flex items-center ${collapsed ? 'justify-center px-2 py-2' : 'justify-between px-4 pb-3 pt-4'}`}
	>
		{#if !collapsed}
			<div class="min-w-0">
				<div class="truncate text-sm font-semibold tracking-wide text-neutral-100">Core</div>
				<div class="mt-1 text-xs text-neutral-500">Personal CLI + dashboard</div>
			</div>
		{/if}
		<button
			class="btn-ghost"
			type="button"
			aria-label={collapsed ? 'Expand sidebar' : 'Collapse sidebar'}
			on:click={handleToggleCollapsed}
			title={collapsed ? 'Expand sidebar' : 'Collapse sidebar'}
		>
			<svg
				class="h-5 w-5"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
				aria-hidden="true"
			>
				{#if collapsed}
					<path d="M9 6l6 6-6 6" />
				{:else}
					<path d="M15 18l-6-6 6-6" />
				{/if}
			</svg>
		</button>
	</div>

	<div class={`px-2 pb-4 ${collapsed ? 'pt-2' : ''}`}>
		<Menu {collapsed} />
	</div>

	<div
		class={`mt-auto border-t border-neutral-800 p-4 text-xs text-neutral-500 ${collapsed ? 'text-center' : ''}`}
	>
		<div title={`v${import.meta.env.VITE_APP_VERSION ?? '0.0.0'}`}>
			{#if collapsed}
				v
			{:else}
				v{import.meta.env.VITE_APP_VERSION ?? '0.0.0'}
			{/if}
		</div>
	</div>
</aside>
