<script lang="ts">
	import Card from './Card.svelte';

	let randomValue: number | null = null;
	let loading = false;
	let error: string | null = null;

	const fetchRandom = async () => {
		loading = true;
		error = null;
		try {
			const res = await fetch('/api/random');
			if (!res.ok) {
				throw new Error(`Request failed with status ${res.status}`);
			}
			const data = await res.json();
			randomValue = data.value ?? null;
		} catch (e) {
			error = e instanceof Error ? e.message : 'Unknown error';
			randomValue = null;
		} finally {
			loading = false;
		}
	};
</script>

<Card title="Random number demo" subtitle="Calls Go /api/random">
	<div class="space-y-3">
		<button
			class="rounded bg-neutral-800 px-3 py-1.5 text-sm font-medium text-neutral-100 hover:bg-neutral-700 disabled:opacity-60"
			type="button"
			on:click={fetchRandom}
			disabled={loading}
		>
			{#if loading}
				Fetching...
			{:else}
				Get random number
			{/if}
		</button>

		{#if error}
			<p class="text-xs text-red-400">Error: {error}</p>
		{:else if randomValue !== null}
			<p class="text-2xl font-semibold">{randomValue}</p>
		{:else}
			<p class="text-xs text-neutral-500">Click the button to fetch a random value.</p>
		{/if}
	</div>
</Card>
