<script lang="ts">
	import { Card, RandomCard, JournalCard } from '$lib';
</script>

<section class="space-y-6">
	<div class="space-y-2">
		<h1 class="text-3xl font-semibold tracking-tight text-neutral-100">Core</h1>
		<p class="text-sm text-neutral-400">A minimal dashboard for your personal CLI.</p>
	</div>

	<div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
		<RandomCard widgetId="home-random" />
		<JournalCard widgetId="journal-card" />
		<Card title="Example widget" subtitle="Placeholder card for future features">
			<p class="text-sm text-neutral-400">
				This space can hold upcoming widgets and quick actions.
			</p>
		</Card>
	</div>
</section>
