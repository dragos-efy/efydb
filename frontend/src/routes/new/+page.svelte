<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import fetchJson, { fetchFormJson } from '$lib/fetchjs';
	import { onMount } from 'svelte';

	let id: string | null;
	let title: string;
	let description: string;
	let screenshot: any;
	let config: any;
	let imageConfig: any;

	const loadTheme = async (id: string) => {
		let theme = await fetchJson(`/themes/${id}`, {});
		if (theme.message) {
			alert(theme.message);
			return;
		}
		title = theme.title;
		description = theme.description;
	}

	onMount(() => {
		id = $page.url.searchParams.get("id");
		if (id != null) {
			loadTheme(id);
		}
	});

	const create = async () => {
		if (!title || !description || !screenshot.files[0] || !config.files[0]) {
			alert('Please provide all the info!');
			return;
		}

		let formData = new FormData();
		formData.append(
			'data',
			JSON.stringify({
				title: title,
				description: description
			})
		);
		formData.append('screenshot', screenshot.files[0]);
		formData.append('config', config.files[0]);
		if (imageConfig.files.length > 0) formData.append('imageConfig', imageConfig.files[0]);

		const response = (id == null) ? await fetchFormJson('/themes/create', {
			method: 'POST',
			body: formData
		}) : await fetchFormJson(`/themes/edit?id=${id}`, {
			method: 'PATCH',
			body: formData,
		});
		if (response.ID) goto('/');
	};
</script>

<svelte:head>
	<title>New theme</title>
</svelte:head>

<section>
	<h3>Create Theme</h3>
	<input type="text" bind:value={title} placeholder="Title" />
	<input type="text" bind:value={description} placeholder="Description" />
	<div id="buttons">
		{#if id}
		<p>Please note that all files must be re-uploaded when editing a theme!</p>
		{/if}
		<div id="files">
			<input type="file" id="screenshot" accept="image/*" bind:this={screenshot} />
			<button><label for="screenshot">Screenshot</label></button>
			<input type="file" id="config" accept="application/json" bind:this={config} />
			<button><label for="config">Config</label></button>
			<input type="file" id="imageConfig" accept="application/json" bind:this={imageConfig} />
			<button><label for="imageConfig">Image config</label></button>
		</div>
		<button on:click={create}>Create</button>
	</div>
</section>

<style>
	h3 {
		width: 100%;
	}

	#buttons {
		display: flex;
		width: 100%;
		justify-content: space-between;
	}

	#files {
		display: flex;
	}

	#files button {
		margin-right: 5rem;
		cursor: pointer;
	}

	input[type='file'] {
		display: none;
	}
</style>
