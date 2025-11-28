<script context="module" lang="ts">
	export type MenuItem = { label: string; path: string };
</script>

<script lang="ts">
	import { page } from '$app/stores';
	import { menuItems as defaultItems } from '$lib/menu';

	export let items: MenuItem[] = defaultItems;
	export let exact = false;

	$: currentPath = $page.url.pathname;

	function isActive(path: string) {
		if (exact) return currentPath === path;
		return currentPath === path || (path !== '/' && currentPath.startsWith(path + '/'));
	}

	const base = 'block rounded-md px-3 py-2 text-sm font-medium transition-colors';
	const active = 'bg-white/10 text-white';
	const inactive = 'text-gray-300 hover:bg-white/5 hover:text-white';
</script>

<nav aria-label="Main navigation">
	<ul class="m-0 list-none p-0 space-y-1">
		{#each items as item}
			<li>
				<a
					href={item.path}
					class={`${base} ${isActive(item.path) ? active : inactive}`}
					aria-current={isActive(item.path) ? 'page' : undefined}
					data-sveltekit-preload-data
				>
					{item.label}
				</a>
			</li>
		{/each}
	</ul>
</nav>
