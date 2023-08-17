<script lang="ts">
	import fetchJson from '$lib/fetchjs';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { getRole, getUsername } from '$lib/token';
	import { goto } from '$app/navigation';

	let theme: any;
	let showApproveBtn = false;
	let showEditBtn = false;
	let showDeleteBtn = false;

	const getId = () => parseInt($page.url.searchParams.get('id')!);

	onMount(async () => {
		theme = await fetchJson(`/themes/${getId()}`, {});
		let role = getRole();
		let isAdmin = !isNaN(role) && role != 0;
		let isCreator = getUsername() == theme.username;

		showApproveBtn = isAdmin && !theme.approved;
		showEditBtn = isCreator;
		showDeleteBtn = isAdmin || isCreator;
	});

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

const theme_config = async (a) => {
	const b = await fetch(a),
	c = await b.json(),
	i = getId();
	console.log(c);

	document.querySelector(`style[efy_theme="${i}"]`).textContent = `html.efy_theme_preview {
		--efy_color: ${c.gradient}!important;
		--efy_color_trans: ${c.gradient_trans}!important;
		--efy_color_text_var: hsl(${c.colorText})!important;
		--efy_text: hsl(${c.text})!important;
		--efy_radius: ${c.radius}!important; --efy_radius0: calc(${c.radius} / 1.25)!important; --efy_radius00: calc(${c.radius} / 3)!important;
		--efy_border: ${c.border_size} solid hsla(var(--efy_bg_var) / var(--efy_border_alpha))!important;
		--efy_bg_var: 0 0% 100%!important;
		--efy_color_bgcol_var: 170 100% 10%; --efy_bg: hsl(var(--efy_color_bgcol_var))!important;
		--efy_bg1: hsla(var(--efy_bg_var) / .07)!important;
	}`;
	document.querySelector('.preview').addEventListener('click', ()=>{
		document.querySelector('html').classList.toggle('efy_theme_preview')
	})
}
</script>

<svelte:head>
	<title>Theme preview</title>
</svelte:head>

<section>
	{#if theme}
		<div id="theme">
			<img src={theme.screenshot} alt={theme.title} />
			<div class="info">
				<span class="summary">
					<h5>{theme.title}</h5>
				</span>
				<span class="actions">
					<a class="username" href={`/user?username=${theme.username}`}>
						<i efy_icon="user" />{theme.username}
					</a>
					<button class="preview">
						<i efy_icon="play" />Preview
					</button>
					<a href={theme.config} download="{theme.title}_efy_config.json" role="button">
						<i efy_icon="arrow_down" />Config
					</a>
					{#if theme.imageConfig}
					<a href={theme.imageConfig} download="{theme.title}_efy_images.json" role="button">
						<i efy_icon="arrow_down" />Images
					</a	>
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
	}

	#theme img {
		aspect-ratio: 16/9;
		border-radius: var(--efy_radius) var(--efy_radius) 0 0;
	}

	#theme .actions {
		display: flex;
		flex-wrap: wrap;
		gap: var(--efy_gap0);
		padding: var(--efy_gap0);
		border-top: var(--efy_border);
		border-bottom: var(--efy_border);
	}
	#theme .actions :is(a, button, i:not([efy_icon=check])) {
		margin: 0;
	}

	#theme .actions :is(a, button), #theme :is(.username, .preview) {
		display: flex;
		align-items: center;
	}

	#theme :is(.description, .summary) {
		padding: 10rem;
	}

	.info {
		display: flex;
		flex-direction: column;
	}

	.username {
		background: var(--efy_bg1);
		border-radius: var(--efy_radius);
		border: var(--efy_border);
		background-clip: none;
		-webkit-background-clip: none;
		-webkit-text-fill-color: var(--efy_text);
		padding: var(--efy_padding);
		gap: 8rem;
	}

	.preview {
		gap: 8rem;
	}

	#theme [efy_icon='arrow_down']::before,
	#theme [efy_icon='edit']::before,
	#theme [efy_icon='remove']::before {
		position: relative;
		margin: 0 8rem 0 0;
		display: inline-block;
		-webkit-background-clip: text !important;
		background-clip: text !important;
		background: var(--efy_color);
		color: transparent;
	}

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
}
[efy_theme] > * {
	pointer-events: none
}
[efy_theme] * {
	backdrop-filter: none!important
}
[efy_theme] progress {
	width: calc(100% - 26rem)
}
[efy_theme] :is(button, i, [efy_select] label) {
	margin: 0
}
[efy_theme] mark {
	line-height: 1.5
}
[efy_theme] :is(i, p) {
	line-height: 1
}
[efy_theme] [efy_card] {
	padding: 2rem 7rem; width: fit-content
}
[efy_theme] [efy_card].bg {
	background: var(--efy_bg)
}
</style>
