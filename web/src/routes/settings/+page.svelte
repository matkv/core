<script lang="ts">
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { page } from '$app/stores';
	import { widgetRegistry } from '$lib/state/widgetState';

	type SettingsPageState = {
		settings: Record<string, any> | null;
		error: string | null;
	};

	let settings: Record<string, any> | null = null;
	let error: string | null = null;
	const widgetId = 'settings-page';

	const restoreFromRegistry = () => {
		const pathname = get(page).url.pathname;
		const stored = widgetRegistry.getWidgetState<SettingsPageState>(pathname, widgetId);
		if (stored) {
			settings = stored.settings;
			error = stored.error;
		}
	};

	const persistToRegistry = () => {
		const pathname = get(page).url.pathname;
		widgetRegistry.setWidgetState<SettingsPageState>(pathname, widgetId, {
			settings,
			error
		});
	};

	onMount(async () => {
		restoreFromRegistry();
		// If we already have data or an error, don't refetch
		if (settings !== null || error !== null) {
			return;
		}

		try {
			const res = await fetch('/api/settings');
			if (!res.ok) throw new Error(`HTTP ${res.status}`);
			settings = await res.json();
		} catch (e) {
			error = (e as Error).message;
		} finally {
			persistToRegistry();
		}
	});
</script>

<section class="space-y-4">
	<h1 class="text-2xl font-bold text-neutral-100">Settings</h1>
	<p class="text-neutral-400">App configuration loaded from server.</p>

	{#if error}
		<div class="text-red-500">Failed to load settings: {error}</div>
	{:else if settings}
		<div class="space-y-4">
			<div class="space-y-2">
				<h2 class="text-lg font-semibold text-neutral-100">Key settings</h2>
				<div class="grid grid-cols-1 md:grid-cols-2 gap-3">
					<div class="border border-neutral-700 rounded p-3">
						<div class="text-sm text-neutral-400">Obsidian vault</div>
						<div class="text-neutral-100 break-all">{settings?.Paths?.ObsidianVault ?? '—'}</div>
					</div>
				</div>
			</div>

			<div class="space-y-2">
				<h3 class="text-md font-semibold text-neutral-100">Raw JSON</h3>
				<pre
					class="bg-neutral-800 text-neutral-100 p-3 rounded text-sm overflow-auto">{JSON.stringify(
						settings,
						null,
						2
					)}</pre>
			</div>
		</div>
	{:else}
		<div class="text-neutral-400">Loading settings…</div>
	{/if}
</section>
