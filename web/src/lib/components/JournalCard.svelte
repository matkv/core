<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { widgetRegistry } from '$lib/state/widgetState';
	import Card from './Card.svelte';

	export let widgetId: string = 'journal-card';

	type JournalWidgetState = {
		text: string;
	};

	let text = '';
	const scopePath = '/journal';

	const restoreFromRegistry = () => {
		const stored = widgetRegistry.getWidgetState<JournalWidgetState>(scopePath, widgetId);
		if (stored) {
			text = stored.text;
		}
	};

	const persistToRegistry = () => {
		widgetRegistry.setWidgetState<JournalWidgetState>(scopePath, widgetId, { text });
	};

	onMount(() => {
		restoreFromRegistry();
	});

	const handleInput = (event: Event) => {
		const target = event.target as HTMLTextAreaElement;
		text = target.value;
		persistToRegistry();
	};
</script>

<Card title="Journal" subtitle="Write a quick note">
	<div class="space-y-2">
		<label class="block text-xs text-neutral-400" for="journal-text">Note</label>
		<textarea
			id="journal-text"
			class="w-full rounded-md border border-neutral-700 bg-neutral-900 px-2 py-1.5 text-sm text-neutral-100 placeholder:text-neutral-500 focus:outline-none focus:ring-1 focus:ring-neutral-500"
			rows="4"
			placeholder="Type something..."
			value={text}
			on:input={handleInput}
		></textarea>
	</div>
</Card>
