<script lang="ts">
	// lightweight header with title, avatar, and actions
	export let onOpenCommandPalette: () => void;
</script>

<header
	class="sticky top-0 z-30 flex items-center justify-between border-b border-neutral-800/80 bg-neutral-950/60 px-4 py-3 backdrop-blur supports-backdrop-filter:bg-neutral-950/40 sm:px-6 lg:px-8"
>
	<div class="flex items-center gap-4">
		<div class="min-w-0">
			<h1 class="truncate text-base font-semibold tracking-tight text-neutral-100 sm:text-lg">
				Dashboard
			</h1>
			<p class="hidden text-xs text-neutral-400 sm:block">Quick access to core tools</p>
		</div>
	</div>

	<div class="flex items-center gap-3">
		<button
			class="hidden md:inline-flex items-center gap-2 rounded-lg border border-neutral-800 bg-neutral-950/30 px-3 py-2 text-xs font-medium text-neutral-200 shadow-sm shadow-black/20 transition hover:bg-neutral-900/60 motion-reduce:transition-none focus-ring"
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
			<span class="whitespace-nowrap">Command palette</span>
			<span class="kbd">Ctrl+Shift+Alt+P</span>
		</button>
	</div>
</header>
