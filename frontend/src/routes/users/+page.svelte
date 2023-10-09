<script lang="ts">
	import fetchJson from '$lib/fetchjs';
	import { onMount } from 'svelte';

	let users: any[] = [];

	const fetchUsers = async () => {
		let response = await fetchJson('/users', {});
		if (response.message) {
			alert(response.message);
		} else {
			users = response;
		}
	};

	const changeUserRank = (user: any) => {
		fetchJson('/users/promote', {
			method: 'POST',
			body: JSON.stringify(user)
		}).then((resp) => {
			if (resp.message) alert(resp.message);
		});
		return null;
	};

	onMount(() => {
		fetchUsers();
	});
</script>

<svelte:head>
	<title>Users</title>
</svelte:head>

<section>
	{#if users}
		{#each users as user}
			<div efy_card>
				<p>{user.name}</p>
				<input type="range" min="0" max="2" bind:value={user.role} />
				<button on:click={changeUserRank(user)}>Change Rank</button>
			</div>
		{/each}
	{:else}
		<div class="spin" />
	{/if}
</section>

<style>
section {
	justify-content: start;
	gap: var(--efy_gap);
}
div {
	display: flex;
	align-items: center;
	gap: var(--efy_gap0);
	width: 100%;
	padding: var(--efy_gap0);
	border-radius: var(--efy_radius2);
	& * {
		height: min-content;
		margin: 0;
	}
}
p {
	flex-grow: 1;
	font-weight: bold;
	padding-left: 5rem;
	font-size: 20rem;
}
button { white-space: nowrap }
input { width: 70rem }
</style>
