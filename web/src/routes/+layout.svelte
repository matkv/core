<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import Header from '$lib/components/Header.svelte';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import CommandPalette from '$lib/components/CommandPalette.svelte';
	import { onMount } from 'svelte';

	let sidebarCollapsed = false;
	let commandPaletteOpen = false;

	onMount(() => {
		const media = window.matchMedia('(min-width: 768px)');
		// Mobile: collapsed by default. Desktop: expanded by default.
		sidebarCollapsed = !media.matches;
	});

	const toggleSidebarCollapsed = () => {
		sidebarCollapsed = !sidebarCollapsed;
	};

	const openCommandPalette = () => {
		commandPaletteOpen = true;
	};

	const closeCommandPalette = () => {
		commandPaletteOpen = false;
	};

	const handleKeydown = (event: KeyboardEvent) => {
		// Use Ctrl+Shift+Alt+P (or Cmd+Shift+Alt+P on macOS)
		const isMac = navigator.platform.toLowerCase().includes('mac');
		const metaPressed = isMac ? event.metaKey : event.ctrlKey;
		if (metaPressed && event.shiftKey && event.altKey && event.key.toLowerCase() === 'p') {
			event.preventDefault();
			commandPaletteOpen = !commandPaletteOpen;
		}
	};
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</svelte:head>

<svelte:window on:keydown={handleKeydown} />

<div class="app-shell flex">
	<Sidebar collapsed={sidebarCollapsed} on:toggleCollapsed={toggleSidebarCollapsed} />

	<div class="flex-1 flex flex-col">
		<Header onOpenCommandPalette={openCommandPalette} />
		<main id="main" class="px-4 py-6 sm:px-6 lg:px-8">
			<div class="content-area mx-auto">
				<slot />
			</div>
		</main>
	</div>

	<CommandPalette open={commandPaletteOpen} onClose={closeCommandPalette} />
</div>
