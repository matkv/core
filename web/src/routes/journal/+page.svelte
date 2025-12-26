<script lang="ts">
	import { onMount } from 'svelte';
	import { widgetRegistry } from '$lib/state/widgetState';
	import { Card } from '$lib';

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

<section class="space-y-6">
	<div class="space-y-2">
		<h1 class="text-3xl font-semibold tracking-tight text-neutral-100">Journal</h1>
		<p class="text-sm text-neutral-400">Your personal journal entries.</p>
	</div>

	{#if text}
		<Card title="Current note" subtitle="Saved locally in the dashboard state">
			<p class="whitespace-pre-wrap text-sm leading-relaxed text-neutral-200">{text}</p>
		</Card>
	{:else}
		<div class="surface p-5 text-sm text-neutral-400">
			Nothing written yet. Use the journal card to add a note.
		</div>
	{/if}
</section>
