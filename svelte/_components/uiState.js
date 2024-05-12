import { writable } from 'svelte/store';

export let isSideMenuOpen = writable(false); // Side Menu

export let isShowFilterTable = writable(false); // Filter Table