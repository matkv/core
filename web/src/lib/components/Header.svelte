<script lang="ts">
	// lightweight header with title and avatar
	export let onToggleSidebar: () => void;
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
		<div class="w-8 h-8 bg-neutral-700 rounded-full" title="User"></div>
	</div>
</header>
