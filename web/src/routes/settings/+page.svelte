<script lang="ts">
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { page } from '$app/stores';
	import { widgetRegistry } from '$lib/state/widgetState';
	import { Card } from '$lib';

	type Settings = {
		Paths?: {
			ObsidianVault?: string;
			Website?: string;
		};
		Device?: string;
	};

	type SettingsPageState = {
		settings: Settings | null;
		error: string | null;
	};

	let settings: Settings | null = null;
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

<section class="space-y-6">
	<div class="space-y-2">
		<h1 class="text-3xl font-semibold tracking-tight text-neutral-100">Settings</h1>
		<p class="text-sm text-neutral-400">App configuration loaded from server.</p>
	</div>

	{#if error}
		<div class="surface-solid p-5">
			<p class="text-sm font-medium text-red-400">Failed to load settings</p>
			<p class="mt-1 text-sm text-neutral-300">{error}</p>
		</div>
	{:else if settings}
		<div class="grid gap-4 lg:grid-cols-2">
			<Card title="Key settings" subtitle="High-signal configuration values">
				<div class="grid gap-3">
					<div class="surface-solid p-4">
						<div class="text-xs font-medium text-neutral-400">Obsidian vault</div>
						<div class="mt-1 break-all text-sm text-neutral-100">
							{settings?.Paths?.ObsidianVault ?? '—'}
						</div>
					</div>
				</div>
			</Card>

			<Card title="Raw JSON" subtitle="Full settings payload">
				<pre
					class="surface-solid max-h-112 overflow-auto p-4 text-xs text-neutral-100">{JSON.stringify(
						settings,
						null,
						2
					)}</pre>
			</Card>
		</div>
	{:else}
		<div class="surface p-5 text-sm text-neutral-400">Loading settings…</div>
	{/if}
</section>
