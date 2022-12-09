<svelte:head>
    <title>Theme preview</title>
</svelte:head>

<section>
{#if theme}
<div id="theme" style="background: url({theme.screenshot});">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    {#if showApproveBtn}
    <button id="approve" on:click={approve}>Approve</button>
    {/if}
    <div>
        <span id="theme-info">
            <h5>{theme.title}</h5>
            <h6>{theme.username}</h6>
        </span>
        <a href={theme.config} download="{theme.title}-config.json">Config</a>
        {#if theme.imageConfig}
        <a href={theme.imageConfig} download="{theme.title}-image-config.json">Image config</a>
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
	import { getRole } from "$lib/token";

    let theme: any;
    let showApproveBtn = false;
    
    onMount(async () => {
        let id = parseInt($page.url.searchParams.get('id')!);
        theme = await fetchJson(`/themes/${id}`, {});
        let role = getRole();
        showApproveBtn = !isNaN(role) && role != 0 && !theme.approved;
    })

    const approve = async () => {
        let id = parseInt($page.url.searchParams.get('id')!);
        let response = await fetchJson(`/themes/approve?id=${id}`, {
            method: 'POST'
        });
        if (response.message) alert(response.message)
        else showApproveBtn = false;
    }
</script>

<style>
    #theme {
        height: 60vh;
        width: auto;
        max-width: 100%;
        background-size: contain;
        aspect-ratio: 16/9;
        border-radius: var(--efy_radius);
        position: relative;
    }

    #theme>div {
        display: flex;
        position: absolute;
        bottom: 0;
        width: 100%;
        align-items: center;
        justify-content: space-between;
        padding: 20rem;
    }

    #theme-info {
        display: flex;
        flex-direction: column;
        color: white;
        mix-blend-mode: difference;
    }

    #theme>div>a {
        margin: 0 10rem;
    }

    #approve {
        float: right;
        margin-right: 10rem;
    }
</style>