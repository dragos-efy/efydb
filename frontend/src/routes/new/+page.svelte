<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import fetchJson, { fetchFormJson } from '$lib/fetchjs';
	import { onMount } from 'svelte';

	const SCREENSHOT_SIZE_LIMIT = 512 * 1024;
	const CONFIG_SIZE_LIMIT = 1024 * 1024;
	const DATABASE_SIZE_LIMIT = 50 * 1024 * 1024;

	let id: string | null;
	let title: string;
	let description: string;
	let screenshot: any;
	let config: any;
	let database: any;

	const loadTheme = async (id: string) => {
		let theme = await fetchJson(`/themes/${id}`, {});
		if (theme.message) {
			alert(theme.message);
			return;
		}
		title = theme.title;
		description = theme.description;
	};

	onMount(() => {
		id = $page.url.searchParams.get('id');
		if (id != null) {
			loadTheme(id);
		}
	});

	const exceedsLimit = (fileInput: HTMLInputElement, limit: number, name: string) => {
		if (!fileInput.files?.length) return false;
		if (fileInput.files[0].size > limit) {
			alert(`${name} is too large. Max size: ${limit / 1024 / 1024}KB`);
			return true;
		}
		return false;
	}

	const create = async () => {
		if (!title || !description || !screenshot.files[0] || !config.files[0]) {
			alert('Please provide all the info!');
			return;
		}

		if (exceedsLimit(screenshot, SCREENSHOT_SIZE_LIMIT, "Screenshot") || exceedsLimit(config, CONFIG_SIZE_LIMIT, "Config") || exceedsLimit(database, DATABASE_SIZE_LIMIT, "Database")) {
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
		if (database.files.length > 0) formData.append('database', database.files[0]);

		const response =
			id == null
				? await fetchFormJson('/themes/create', {
						method: 'POST',
						body: formData
				  })
				: await fetchFormJson(`/themes/edit?id=${id}`, {
						method: 'PATCH',
						body: formData
				  });
		if (response.ID) goto('/');
	};
</script>

<svelte:head>
	<title>{id ? 'Edit theme' : 'New theme'}</title>
</svelte:head>

<section>
	<h3>{id ? 'Edit theme' : 'New theme'}</h3>
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
			<input type="file" id="database" accept="application/json" bind:this={database} />
			<button><label for="database">Database</label></button>
		</div>
		<button on:click={create}>Create</button>
	</div>
</section>

<style>
h3 { width: 100% }
#buttons {
	display: flex;
	width: 100%;
	justify-content: space-between;
}
#files { display: flex;
	& button {
		margin-right: 5rem;
		cursor: pointer;
	}
}
input[type='file'] { display: none }
</style>
