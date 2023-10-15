<script lang="ts">
	import fetchJson from '$lib/fetchjs';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { getRole, getToken, getUsername } from '$lib/token';
	import { goto } from '$app/navigation';

	let theme: any;
	let showApproveBtn = false;
	let showEditBtn = false;
	let showDeleteBtn = false;
	let showVoteBtn = false;

	const getId = () => parseInt($page.url.searchParams.get('id')!);

	onMount(async () => {
		theme = await fetchJson(`/themes/${getId()}`, {});
		let role = getRole();
		let isAdmin = !isNaN(role) && role != 0;
		let isCreator = getUsername() == theme.username;
		let isLoggedIn = getToken() != null;

		showApproveBtn = isAdmin && !theme.approved;
		showEditBtn = isCreator;
		showDeleteBtn = isAdmin || isCreator;
		showVoteBtn = isLoggedIn;
	});

	const vote = async () => {
		let newVote = theme.user_score == 1 ? 0 : 1;
		let response = await fetchJson(`/themes/${getId()}/vote?score=${newVote}`, {
			method: 'POST',
		});
		if (response.message) alert(response.message);
		else theme = response;
	}

	const approve = async () => {
		let response = await fetchJson(`/themes/approve?id=${getId()}`, {
			method: 'POST'
		});
		if (response.message) alert(response.message);
		else showApproveBtn = false;
	};

	const editTheme = () => {
		goto(`/new?id=${getId()}`);
	};

	const deleteTheme = async () => {
		let response = await fetchJson(`/themes/delete?id=${getId()}`, {
			method: 'DELETE'
		});
		if (response.message != 'ok') alert(response.message);
		else goto('/');
	};

/*Get the current theme's Config*/
let themeConfig = [];
const theme_config = async (file)=>{
	const result = await fetch(file);
	themeConfig = await result.json();

	document.querySelector(`style[efy_theme="${getId()}"]`).textContent = `html.efy_theme_preview {
		--efy_color: ${themeConfig.gradient}!important;
		--efy_color_trans: ${themeConfig.gradient_trans}!important;
		--efy_color_text_var: hsl(${themeConfig.colorText})!important;
		--efy_text: hsl(${themeConfig.text})!important;
		--efy_radius: ${themeConfig.radius}!important;
		--efy_radius0: calc(${themeConfig.radius} / 1.25)!important;
		--efy_radius00: calc(${themeConfig.radius} / 3)!important;
		--efy_border: ${themeConfig.border_size} solid hsla(var(--efy_bg_var) / var(--efy_border_alpha))!important;
		--efy_bg_var: 0 0% 100%!important;
		--efy_color_bgcol_var: 170 100% 10%;
		--efy_bg: hsl(var(--efy_color_bgcol_var))!important;
		--efy_bg1: hsla(var(--efy_bg_var) / .07)!important;
	}`;
	document.querySelector('.preview').addEventListener('click', ()=>{
		document.querySelector('html').classList.toggle('efy_theme_preview')
	})
}

// download File
const downloadFile =(url, filename)=>{
    fetch(url)
        .then(response => response.blob())
        .then(data => {
            const blobURL = URL.createObjectURL(data);
            const link = document.createElement('a');
            link.href = blobURL;
            link.download = filename;
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        })
        .catch(error => console.error("Failed to download file:", error));
}

</script>

<svelte:head>
	<title>Theme preview</title>
</svelte:head>

<section>
	{#if theme}
		<div id="theme" class="efy_trans_filter efy_shadow_trans">
			<img src={theme.screenshot} alt={theme.title} loading="lazy" />
			<div class="info">
				<span class="summary">
					<h5>{theme.title}</h5>
				</span>
				<span class="actions">
					<a class="username efy_trans_filter efy_shadow_trans" href={`/user?username=${theme.username}`}>
						<i efy_icon="user" />{theme.username}
					</a>
					{#if themeConfig.version}
					<a class="version efy_trans_filter efy_shadow_trans">{themeConfig.version}</a>
					{/if}
					{#if showVoteBtn}
					<button on:click={vote}><i efy_icon="heart" />{theme.score}</button>
					{/if}
					<button class="preview">
						<i efy_icon="play" />Preview
					</button>
					<button on:click={downloadFile(theme.config, `${theme.title}_efy_config.json`)}>
						<i efy_icon="arrow_down" />Config
					</button>
					{#if theme.database}
					<button on:click={downloadFile(theme.config, `${theme.title}_efy_database.json`)}>
						<i efy_icon="arrow_down" />Database
					</button>
					{/if}
					{#if showApproveBtn}
					<button on:click={approve}><i efy_icon="check" />Approve</button>
					{/if}
					{#if showEditBtn}
						<button on:click={editTheme}><i efy_icon="edit" />Edit</button>
					{/if}
					{#if showDeleteBtn}
						<button on:click={deleteTheme}><i efy_icon="remove" />Delete</button>
					{/if}
				</span>
				<p class="description">{theme.description}</p>
				<p class="efy_hide_i">{theme_config(theme.config)}</p>
				<style efy_theme={getId()} class="efy_hide_i" />
			</div>
		</div>
	{:else}
		<div class="spin" />
	{/if}
</section>

<style>
#theme {
	display: flex;
	flex-direction: column;
	width: fit-content;
	justify-content: space-between;
	padding: 0;
	background: var(--efy_bg1);
	border-radius: var(--efy_radius);
	border: var(--efy_border);
	width: 100%;
	max-width: calc(var(--efy_100vh) * 0.7);
	& img {
		aspect-ratio: 16/9;
		object-fit: contain;
		border-radius: var(--efy_radius) var(--efy_radius) 0 0;
		border-bottom: var(--efy_border);
	}
	& .actions {
		display: flex;
		flex-wrap: wrap;
		gap: var(--efy_gap0);
		padding: var(--efy_gap0);
		border-top: var(--efy_border);
		border-bottom: var(--efy_border);
		& :is(a, button, i:not([efy_icon=check])) { margin: 0 }
	}
	& .actions :is(a, button), & :is(.username, .version, .preview) {
		display: flex;
		align-items: center;
	}
	& :is(.description, .summary) { padding: 10rem }
}
#theme i:before {
	position: relative;
	margin: 0 8rem 0 0;
	display: inline-block;
	-webkit-background-clip: text !important;
	background-clip: text !important;
	background: var(--efy_color);
	color: transparent;
}

.info {
	display: flex;
	flex-direction: column;
}
.username, .version {
	background: var(--efy_bg1);
	border-radius: var(--efy_radius);
	border: var(--efy_border);
	background-clip: none;
	-webkit-background-clip: none;
	-webkit-text-fill-color: var(--efy_text);
	padding: var(--efy_padding);
	gap: 8rem;
}
.preview { gap: 8rem }

/*Theme Preview*/
[efy_theme] { display: none!important;
	display: flex;
	width: 168rem;
	border: var(--efy_border);
	background: var(--efy_bg);
	border-radius: var(--efy_radius);
	flex-wrap: wrap;
	gap: var(--efy_gap0);
	align-items: center;
	padding: 10rem;
	& > * { pointer-events: none }
	& * { backdrop-filter: none!important }
	& progress { width: calc(100% - 26rem) }
	& :is(button, i, [efy_select] label) { margin: 0 }
	& mark { line-height: 1.5 }
	& :is(i, p) { line-height: 1 }
	& [efy_card] { padding: 2rem 7rem; width: fit-content }
	& [efy_card].bg { background: var(--efy_bg) }
}
</style>
