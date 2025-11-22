<script lang="ts">
	// lightweight header with title, avatar, and actions
	export let onToggleSidebar: () => void;
	export let onOpenCommandPalette: () => void;
</script>

<header class="bg-neutral-900 border-b border-neutral-800 p-4 flex items-center justify-between">
	<div class="flex items-center gap-4">
		<button
			class="md:hidden p-2 rounded bg-neutral-800"
			aria-label="open menu"
			on:click={onToggleSidebar}
		>
			<svg class="w-5 h-5 text-neutral-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M4 6h16M4 12h16M4 18h16"
				></path>
			</svg>
		</button>
		<h1 class="text-lg font-semibold">Dashboard</h1>
	</div>

	<div class="flex items-center gap-3">
		<button
			class="hidden md:inline-flex items-center gap-1 rounded border border-neutral-700 bg-neutral-900 px-3 py-1.5 text-xs font-medium text-neutral-200 hover:bg-neutral-800"
			type="button"
			on:click={onOpenCommandPalette}
		>
			<svg
				class="h-3.5 w-3.5 text-neutral-400"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			>
				<circle cx="11" cy="11" r="8" />
				<path d="m21 21-4.3-4.3" />
			</svg>
			<span>Command palette</span>
			<span class="rounded border border-neutral-700 px-1 py-0.5 text-[10px] text-neutral-400">
				Ctrl+Shift+Alt+P
			</span>
		</button>
	</div>
</header>
