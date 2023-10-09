<script lang="ts">
	// import { getRole } from "$lib/token";
	import { onMount } from 'svelte';
	import ThemePreview from '../components/ThemePreview.svelte';
	import fetchJson from '../lib/fetchjs';

	import { getRole, getToken } from '$lib/token';

	let themes: any[] = [];
	let role: number = 0;
	let showUnapproved = false;

	let token: string | null;
	let nextPage = 1;
	let loading = false;

	const fetchThemes = async () => {
		if (loading) return;

		loading = true;
		const newThemes = await fetchJson(
			`/themes?unapproved=${showUnapproved}&page=${nextPage}&limit=20&sort=score`,
			{}
		);
		themes.push(...newThemes);
		themes = themes;

		nextPage += 1;
		loading = false;
	};

	const onToggleShowUnapproved = () => {
		showUnapproved = !showUnapproved;
		themes = [];
		nextPage = 1;
		fetchThemes();
	};

	const initScrollListener = () => {
		document.addEventListener('scroll', () => {
			if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
				fetchThemes();
			}
		});
	};

	onMount(() => {
		token = getToken();
		role = getRole();
		fetchThemes();
		initScrollListener();
	});

onMount(async () => {
	let input = document.querySelector('header .search input'),
	button = document.querySelector('header .search button'),
	grid = document.querySelector('#themes-grid'),
	links = document.querySelectorAll('header [href="/"]');

	const update =(value = input.value)=>{
		fetchJson(`/themes?q=${value}`, {}).then(function(result) {
			document.querySelectorAll('#themes-grid > .theme')
				.forEach(theme =>{ theme.remove() });
			result.forEach(theme =>{
				new ThemePreview({ target: grid, props: { theme }});
			})
		})
	}

	button.addEventListener('click', async ()=>{
		update()
	})
	input.addEventListener('keyup', ({key})=>{
		if (key === 'Enter'){ update()}
	})
	links.forEach(a =>{
		a.addEventListener('click', ()=> update(''))
	})
});
</script>

<svelte:head>
	<title>Themes</title>
</svelte:head>

<section efy_select>
	{#if role > 0}
		<input
			type="checkbox"
			id="showUnapproved"
			on:change={onToggleShowUnapproved}
			bind:value={showUnapproved}
		/>
		<label for="showUnapproved">Unapproved</label>
	{/if}
	<span id="themes-grid">
		{#if token}
			<a class="new-theme efy_color_trans" efy_card href="/new">
				<i efy_icon="plus" />
				<p>New Theme</p>
			</a>
		{/if}
		{#if themes.length}
			{#each themes as theme}
				<ThemePreview {theme} />
			{/each}
		{:else}
			<div class="spin" />
		{/if}

	</span>
</section>

<style>
	label {
		align-self: flex-end;
		margin: 0 0 15rem 0;
	}
	section {
		justify-content: start;
	}
	#themes-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(240rem, 1fr));
		flex-wrap: wrap;
		gap: var(--efy_gap);
		width: 100%;
	}
	.new-theme {
		display: flex;
		flex-direction: column;
		gap: 15rem;
		align-items: center;
		place-content: center;
		-webkit-background-clip: text, padding-box !important;
		background-clip: text, padding-box !important;
		background: var(--efy_color), var(--efy_color_trans);
		min-height: 215rem;
		& i:before {
			font-size: 40rem;
			margin: 0;
			-webkit-background-clip: text !important;
			background-clip: text !important;
			background: var(--efy_color);
		}
	}
	.new-theme * {margin: 0}
</style>
