<script lang="ts">
	import Header from './Header.svelte';
	import Sidebar from './Sidebar.svelte';

	let sidebarOpen = false;

	const toggleSidebar = () => {
		sidebarOpen = !sidebarOpen;
	};

	const closeSidebar = () => {
		sidebarOpen = false;
	};
</script>

<div class="min-h-screen flex bg-black">
	<Sidebar open={sidebarOpen} on:close={closeSidebar} />

	<div class="flex-1 flex flex-col">
		<Header onToggleSidebar={toggleSidebar} />
		<main class="p-6">
			<div class="content-area mx-auto">
				<slot />
			</div>
		</main>
	</div>
</div>
