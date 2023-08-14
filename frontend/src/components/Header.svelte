<script lang="ts">
	import { getRole, getToken } from '$lib/token';
	import { onMount } from 'svelte';

	let token: string | null;
	let role: number | null;

	onMount(() => {
		token = getToken();
		role = getRole();
	});
</script>

<header>
	<a href="/">efyDB</a>
	<nav>
		<a href="/">Themes</a>
		<a href="/account">Account</a>
		{#if role && role > 0}
			<a href="/users">Users</a>
		{/if}
		<a href="https://efy.ooo">Apps</a>
		<button efy_sidebar_btn class="efy_square_btn"><i efy_icon="menu" /></button>
	</nav>
</header>

<style>
	header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 15rem;
		border-bottom: var(--efy_border);
	}
	header > nav {
		padding: 0;
		display: flex;
		align-items: center;
		gap: 15rem;
	}
	header [efy_sidebar_btn] {
		background: transparent;
		padding: 0;
		margin: 0;
		border: 0;
	}
	header :is(a, button) {
		-webkit-text-fill-color: var(--efy_text);
		color: var(--efy_text);
		background: transparent;
		margin: 0;
	}
</style>
