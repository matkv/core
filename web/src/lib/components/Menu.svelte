<script lang="ts">
	import { page } from '$app/stores';
	import { menuItems as defaultItems } from '$lib/menu';

	type MenuItem = { label: string; path: string };

	export let items: MenuItem[] = defaultItems;

	function normalize(p: string) {
		if (p.length > 1 && p.endsWith('/')) return p.slice(0, -1);
		return p;
	}

	function isActive(currentPath: string, path: string) {
		const cur = normalize(currentPath);
		const target = normalize(path);
		if (target === '/') return cur === '/';
		return cur === target || cur.startsWith(target + '/');
	}

	const base = 'block rounded-md px-3 py-2 text-sm font-medium transition-colors duration-150';
	const active = 'bg-neutral-800 text-neutral-100';
	const inactive = 'text-neutral-300 hover:bg-neutral-900 hover:text-neutral-100';
</script>

<nav aria-label="Main navigation">
	<ul class="m-0 list-none p-0 space-y-1">
		{#each items as item}
			<li>
				<a
					href={item.path}
					class={`${base} ${isActive($page.url.pathname, item.path) ? active : inactive}`}
					aria-current={isActive($page.url.pathname, item.path) ? 'page' : undefined}
					data-sveltekit-preload-data
				>
					{item.label}
				</a>
			</li>
		{/each}
	</ul>
</nav>
