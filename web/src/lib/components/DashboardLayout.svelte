<script lang="ts">
	import Header from './Header.svelte';
	import Sidebar from './Sidebar.svelte';
	import CommandPalette from './CommandPalette.svelte';

	let sidebarOpen = false;
	let commandPaletteOpen = false;

	const toggleSidebar = () => {
		sidebarOpen = !sidebarOpen;
	};

	const closeSidebar = () => {
		sidebarOpen = false;
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

<svelte:window on:keydown={handleKeydown} />

<div class="min-h-screen flex bg-black">
	<Sidebar open={sidebarOpen} on:close={closeSidebar} />

	<div class="flex-1 flex flex-col">
		<Header onToggleSidebar={toggleSidebar} onOpenCommandPalette={openCommandPalette} />
		<main class="p-6">
			<div class="content-area mx-auto">
				<slot />
			</div>
		</main>
	</div>

	<CommandPalette open={commandPaletteOpen} onClose={closeCommandPalette} />
</div>
