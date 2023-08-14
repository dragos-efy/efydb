<svelte:head>
    <title>Theme preview</title>
</svelte:head>

<section>
{#if theme}
<img id="theme" src={theme.screenshot} alt={theme.title} />
<div id="btns-container">
    <div>
        <span id="info">
            <h5>{theme.title}</h5>
            <a href={`/user?username=${theme.username}`}><h6>@{theme.username}</h6></a>
        </span>
        <span class="actions">
            <a href={theme.config} download="{theme.title}_efy_config.json" role="button"><i efy_icon="arrow_down"></i>Config</a>
            {#if theme.imageConfig}
            <a href={theme.imageConfig} download="{theme.title}_efy_images.json" role="button"><i efy_icon="arrow_down"></i>Images</a>
            {/if}
        </span>
    </div>
    <div>
        {#if showApproveBtn}
        <button on:click={approve}><i efy_icon="check"></i>Approve</button>
        {/if}
        {#if showDeleteBtn}
        <button on:click={deleteTheme}><i efy_icon="delete"></i>Delete</button>
        {/if}
    </div>
</div>
{:else}
<div class="spin"></div>
{/if}
</section>

<script lang="ts">
	import fetchJson from "$lib/fetchjs";
	import { onMount } from "svelte";
    import { page } from '$app/stores';
	import { getRole, getUsername } from "$lib/token";
	import { goto } from "$app/navigation";

    let theme: any;
    let showApproveBtn = false;
    let showDeleteBtn = false;
    
    const getId = () => parseInt($page.url.searchParams.get('id')!);

    onMount(async () => {
        theme = await fetchJson(`/themes/${getId()}`, {});
        let role = getRole();
        let isAdmin = !isNaN(role) && role != 0;
        showApproveBtn = isAdmin && !theme.approved;
        showDeleteBtn = isAdmin || getUsername() == theme.username;
    })

    const approve = async () => {
        let response = await fetchJson(`/themes/approve?id=${getId()}`, {
            method: 'POST'
        });
        if (response.message) alert(response.message);
        else showApproveBtn = false;
    }

    const deleteTheme = async () => {
        let response = await fetchJson(`/themes/delete?id=${getId()}`, {
            method: 'DELETE'
        });
        if (response.message != "ok") alert(response.message);
        else goto("/");
    }
</script>

<style>
    #theme {
        height: auto;
        max-width: 40vw;
        background-size: contain;
        aspect-ratio: 16/9;
        border-radius: var(--efy_radius);
    }

    #btns-container {
        display: flex;
        width: 100%;
        max-width: 40vw;
        align-items: center;
        justify-content: space-between;
        padding: 20rem 0rem;
    }

    #btns-container > div {
        display: flex;
        align-items: center;
    }

    #info {
        display: flex;
        flex-direction: column;
        padding: 5rem;
    }

    #btns-container a {
        margin: 10rem;
    }

    #btns-container button {
        margin: 0 5rem;
    }

    [efy_icon=arrow_down]:before, [efy_icon=delete]:before {
        position: relative;
        margin: 0 8rem 0 0;
        display: inline-block;
        -webkit-background-clip: text!important;
        background-clip: text!important;
        background: var(--efy_color);
        color: transparent;
    }
</style>