<script lang="ts">
	type Command = {
		id: string;
		name: string;
		group?: string;
		description?: string;
	};

	export let open = false;
	export let onClose: () => void;

	const commands: Command[] = [
		{
			id: 'open-settings',
			name: 'Open settings',
			group: 'Navigation',
			description: 'Jump to the settings dashboard'
		},
		{
			id: 'toggle-dark-mode',
			name: 'Toggle dark mode',
			group: 'Appearance',
			description: 'Switch between light and dark themes'
		}
	];

	let query = '';
	let highlightedIndex = 0;

	const filtered = () => {
		const q = query.toLowerCase().trim();
		if (!q) return commands;
		return commands.filter(
			(cmd) => cmd.name.toLowerCase().includes(q) || cmd.group?.toLowerCase().includes(q)
		);
	};

	const handleKeydown = (event: KeyboardEvent) => {
		if (!open) return;

		if (event.key === 'Escape') {
			event.preventDefault();
			onClose();
			return;
		}

		const items = filtered();
		if (!items.length) return;

		if (event.key === 'ArrowDown') {
			event.preventDefault();
			highlightedIndex = (highlightedIndex + 1) % items.length;
		} else if (event.key === 'ArrowUp') {
			event.preventDefault();
			highlightedIndex = (highlightedIndex - 1 + items.length) % items.length;
		} else if (event.key === 'Enter') {
			event.preventDefault();
			const cmd = items[highlightedIndex];
			if (cmd) {
				// For now, just close on "run"
				onClose();
			}
		}
	};

	const handleBackdropClick = () => {
		onClose();
	};
</script>

{#if open}
	<div
		class="fixed inset-0 z-50 flex items-start justify-center pt-24"
		on:keydown={handleKeydown}
		role="dialog"
		aria-modal="true"
		tabindex="-1"
	>
		<button
			class="fixed inset-0 bg-black/50 cursor-default"
			type="button"
			aria-label="Close command palette"
			on:click={handleBackdropClick}
		></button>

		<div
			class="relative w-full max-w-xl rounded-lg border border-neutral-800 bg-neutral-900 shadow-xl"
		>
			<div class="flex items-center gap-2 border-b border-neutral-800 px-3 py-2">
				<svg
					class="h-4 w-4 text-neutral-500"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
				>
					<circle cx="11" cy="11" r="8" />
					<path d="m21 21-4.3-4.3" />
				</svg>
				<input
					class="flex-1 bg-transparent text-sm text-neutral-100 placeholder:text-neutral-500 outline-none"
					placeholder="Run a command..."
					bind:value={query}
				/>
				<span
					class="rounded border border-neutral-700 px-1.5 py-0.5 text-[10px] font-medium text-neutral-400"
				>
					Ctrl+Shift+Alt+P
				</span>
			</div>

			<ul class="max-h-64 overflow-y-auto py-1 text-sm">
				{#if filtered().length === 0}
					<li class="px-3 py-2 text-neutral-500">No commands match "{query}"</li>
				{:else}
					{#each filtered() as cmd, index}
						<li
							class="flex cursor-pointer flex-col gap-0.5 px-3 py-2 {index === highlightedIndex
								? 'bg-neutral-800 text-neutral-50'
								: 'text-neutral-200 hover:bg-neutral-800/60'}"
						>
							<div class="flex items-center justify-between">
								<span>{cmd.name}</span>
								{#if cmd.group}
									<span class="text-[10px] uppercase tracking-wide text-neutral-500"
										>{cmd.group}</span
									>
								{/if}
							</div>
							{#if cmd.description}
								<p class="text-xs text-neutral-400">{cmd.description}</p>
							{/if}
						</li>
					{/each}
				{/if}
			</ul>
		</div>
	</div>
{/if}
