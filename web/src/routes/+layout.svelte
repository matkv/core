<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import DashboardLayout from '$lib/components/DashboardLayout.svelte';
	export let data;
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</svelte:head>

<DashboardLayout>
	<slot />
</DashboardLayout>
