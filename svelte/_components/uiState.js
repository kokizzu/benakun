import { writable } from 'svelte/store';

/** @type {import('svelte/store').Writable<boolean>} */
export const IsShrinkMenu = writable(false);

/** @type {import('svelte/store').Writable<boolean>} */
export const IsShowFilterTable = writable(false);