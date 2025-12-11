<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { widgetRegistry } from '$lib/state/widgetState';

	type JournalWidgetState = {
		text: string;
	};

	let text = '';
	const widgetId = 'journal-card';
	const scopePath = '/journal';

	const restoreFromRegistry = () => {
		const stored = widgetRegistry.getWidgetState<JournalWidgetState>(scopePath, widgetId);
		if (stored) {
			text = stored.text;
		}
	};

	onMount(() => {
		restoreFromRegistry();
	});
</script>

<section class="space-y-4">
	<h1 class="text-2xl font-bold text-neutral-100">Journal</h1>
	<p class="text-neutral-400">Your personal journal entries.</p>

	{#if text}
		<div class="p-4 border border-neutral-700 rounded bg-neutral-900">
			<h2 class="text-lg font-semibold text-neutral-100">Current note</h2>
			<p class="text-neutral-300 mt-2 whitespace-pre-wrap">{text}</p>
		</div>
	{:else}
		<div class="text-neutral-400">Nothing written yet. Use the journal card to add a note.</div>
	{/if}
</section>
