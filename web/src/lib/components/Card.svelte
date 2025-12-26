<script lang="ts">
	export let title: string = '';
	export let subtitle: string = '';
</script>

<div class="surface surface-hover p-5">
	{#if title || subtitle || $$slots.actions}
		<div class="flex items-start justify-between gap-4">
			<div class="min-w-0">
				{#if title}
					<h3 class="truncate text-sm font-semibold text-neutral-100">{title}</h3>
				{/if}
				{#if subtitle}
					<p class="mt-1 text-xs text-neutral-400">{subtitle}</p>
				{/if}
			</div>
			{#if $$slots.actions}
				<div class="shrink-0"><slot name="actions" /></div>
			{/if}
		</div>
	{/if}
	<div class="mt-4"><slot /></div>
</div>
