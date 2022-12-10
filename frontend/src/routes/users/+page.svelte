<svelte:head>
    <title>Users</title>
</svelte:head>

<section>
{#if users}
{#each users as user}
<div>
    <span>{user.name}</span>
    <input type="range" min="0" max="2" bind:value={user.role}>
    <button on:click={changeUserRank(user)}>Change Rank</button>
</div>
{/each}
{:else}
<div class="spin"></div>
{/if}
</section>

<script lang="ts">
	import fetchJson from "$lib/fetchjs";
	import { onMount } from "svelte";

    let users: any[] = [];

    const fetchUsers = async () => {
        let response = await fetchJson("/users", {});
        if (response.message) {
            alert(response.message)
        } else {
            users = response;
        }
    }

    const changeUserRank = (user: any) => {
        fetchJson("/users/promote", {
            method: 'POST',
            body: JSON.stringify(user),
        }).then((resp) => {
            if (resp.message) alert(resp.message);
        })
        return null;
    }

    onMount(()=> {
        fetchUsers();
    })
</script>

<style>
    div {
        display: flex;
        align-items: center;
        width: 100%;
    }

    div * {
        height: min-content;
        margin: 5rem 10rem;
    }

    span {
        flex-grow: 1;
    }

    button {
        white-space: nowrap;
    }

    input {
        width: 70rem;
    }
</style>