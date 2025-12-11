import { get, writable } from 'svelte/store';

export type WidgetState = unknown;

export type PageWidgets = Record<string, WidgetState>;

export type WidgetRegistry = Record<string, PageWidgets>;

function createWidgetRegistry() {
    const store = writable<WidgetRegistry>({});

    const { subscribe, update } = store;

    function getWidgetState<T = unknown>(pathname: string, widgetId: string): T | undefined {
        const registry = get(store);
        return (registry[pathname]?.[widgetId] as T | undefined) ?? undefined;
    }

    function setWidgetState<T = unknown>(pathname: string, widgetId: string, state: T): void {
        update((registry) => {
            const pageWidgets = registry[pathname] ?? {};
            return {
                ...registry,
                [pathname]: {
                    ...pageWidgets,
                    [widgetId]: state
                }
            };
        });
    }

    function updateWidgetState<T extends Record<string, unknown>>(
        pathname: string,
        widgetId: string,
        partial: Partial<T>
    ): void {
        update((registry) => {
            const pageWidgets = registry[pathname] ?? {};
            const current = (pageWidgets[widgetId] as T | undefined) ?? ({} as T);
            return {
                ...registry,
                [pathname]: {
                    ...pageWidgets,
                    [widgetId]: {
                        ...current,
                        ...partial
                    }
                }
            };
        });
    }

    function clearWidgetState(pathname?: string, widgetId?: string): void {
        // No arguments: clear everything
        if (!pathname) {
            update(() => ({}));
            return;
        }

        // Path only: clear all widgets for that page
        if (!widgetId) {
            update((registry) => {
                const { [pathname]: _removed, ...rest } = registry;
                return rest;
            });
            return;
        }

        // Path + widget: clear a single widget state
        update((registry) => {
            const pageWidgets = registry[pathname];
            if (!pageWidgets) return registry;

            const { [widgetId]: _removed, ...restWidgets } = pageWidgets;
            return {
                ...registry,
                [pathname]: restWidgets
            };
        });
    }

    return {
        subscribe,
        getWidgetState,
        setWidgetState,
        updateWidgetState,
        clearWidgetState
    };
}

export const widgetRegistry = createWidgetRegistry();
