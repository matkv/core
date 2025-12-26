<script lang="ts">
	import { page } from '$app/stores';
	import { resolve } from '$app/paths';
	import { menuItems as defaultItems } from '$lib/menu';

	type MenuItem = (typeof defaultItems)[number];

	export let items: readonly MenuItem[] = defaultItems;
	export let collapsed = false;

	const badgeClass =
		'inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-md border border-neutral-800 bg-neutral-950/40 text-xs font-semibold text-neutral-200';

	const badgeFor = (label: string) => label.trim().slice(0, 1).toUpperCase();

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

	const base =
		'block rounded-md px-3 py-2 text-sm font-medium transition-colors duration-150 motion-reduce:transition-none';
	const active =
		'bg-neutral-900/60 text-neutral-50 ring-1 ring-neutral-800/80 shadow-sm shadow-black/20';
	const inactive = 'text-neutral-300 hover:bg-neutral-900/50 hover:text-neutral-100';
	const focus = 'focus-ring';
</script>

<nav aria-label="Main navigation">
	<ul class="m-0 list-none p-0 space-y-1">
		{#each items as item (item.path)}
			<li>
				<a
					href={resolve(item.path)}
					class={`${base} ${focus} ${collapsed ? 'flex h-10 items-center justify-center px-2' : 'flex items-center gap-2'} ${
						isActive($page.url.pathname, item.path) ? active : inactive
					}`}
					aria-current={isActive($page.url.pathname, item.path) ? 'page' : undefined}
					aria-label={collapsed ? item.label : undefined}
					title={collapsed ? item.label : undefined}
					data-sveltekit-preload-data
				>
					<span class={badgeClass} aria-hidden="true">{badgeFor(item.label)}</span>
					<span class={collapsed ? 'sr-only' : 'truncate'}>{item.label}</span>
				</a>
			</li>
		{/each}
	</ul>
</nav>
