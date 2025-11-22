<script lang="ts">
	export let title: string = '';
	export let subtitle: string = '';
</script>

<div class="bg-neutral-900 border border-neutral-800 rounded-lg shadow-sm p-4">
	{#if title}
		<div class="flex items-baseline justify-between">
			<h3 class="text-sm font-medium">{title}</h3>
		</div>
	{/if}
	{#if subtitle}
		<p class="text-xs text-neutral-400">{subtitle}</p>
	{/if}
	<div class="mt-3"><slot /></div>
</div>
