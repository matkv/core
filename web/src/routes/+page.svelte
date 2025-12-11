<script lang="ts">
	import { Card, RandomCard, JournalCard } from '$lib';
</script>

<section class="space-y-4">
	<h1 class="text-2xl font-semibold">Core</h1>
	<p class="text-sm text-neutral-400">This is the home page of the core dashboard.</p>

	<p class="text-sm text-neutral-400">
		Currently I'm just playing around with the layout & design before starting to work on the actual
		features.
	</p>

	<div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
		<RandomCard widgetId="home-random" />
		<JournalCard widgetId="journal-card" />
		<Card title="Example widget" subtitle="Example widget description">
			<p class="text-2xl font-semibold"></p>
		</Card>
	</div>
</section>
