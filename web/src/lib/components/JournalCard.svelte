<script lang="ts">
	import { onMount } from 'svelte';
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
		<label class="field-label" for="journal-text">Note</label>
		<textarea
			id="journal-text"
			class="field min-h-30 resize-none"
			rows="4"
			placeholder="Type something..."
			value={text}
			on:input={handleInput}
		></textarea>
	</div>
</Card>
