<script lang="ts">
	import { page } from '$app/stores';
	import fetchJson from '$lib/fetchjs';
	import ThemePreview from '../../components/ThemePreview.svelte';
	import { onMount } from 'svelte';

	let user: any;
	let themes: any[];

	const fetchThemes = async () => {
		let username = $page.url.searchParams.get('username');
		let response = await fetchJson(`/users/${username}`, {});
		if (response.message) {
			alert(response.message);
			return;
		}
		user = response.user;
		themes = response.themes;
	};

	onMount(() => {
		fetchThemes();
	});
</script>

<svelte:head>
	<title>{user?.name ?? 'User profile'}</title>
</svelte:head>

<section>
	{#if user}
		<div id="user-info">
			<h5>{user.name}</h5>
			<p>{user.bio}</p>
			<p>Uploaded themes: {themes.length}</p>
		</div>
		<span id="themes-grid">
			{#each themes as theme}
				<ThemePreview {theme} />
			{/each}
		</span>
	{:else}
		<div class="spin" />
	{/if}
</section>

<style>
	section {
		justify-content: start;
		align-items: start;
	}

	#user-info {
		margin: 50rem 5rem;
	}

	#themes-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(240rem, 1fr));
		flex-wrap: wrap;
		gap: var(--efy_gap);
		width: 100%;
	}
</style>
