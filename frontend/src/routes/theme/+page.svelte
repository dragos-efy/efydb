<svelte:head>
    <title>Theme preview</title>
</svelte:head>

<section>
{#if theme}
<div id="theme" class="efy_trans_filter">
    <img src="{theme.screenshot}" alt="Theme thumbnail" />
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div>
        <span id="info">
            <h5>{theme.title}</h5>
            <h6>@{theme.username}</h6>
        </span>
        <span class="actions">
            <a href={theme.config} download="{theme.title}_efy_config.json" role="button"><i efy_icon="arrow_down"></i>Config</a>
            {#if theme.imageConfig}
            <a href={theme.imageConfig} download="{theme.title}_efy_images.json" role="button"><i efy_icon="arrow_down"></i>Images</a>
            {/if}
            {#if showApproveBtn}
            <button id="approve" on:click={approve}><i efy_icon="check"></i>Approve</button>
            {/if}
        </span>
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
        background: var(--efy_bg1);
        width: 100%;
        max-width: calc(var(--efy_100vh) * 0.7);
        border-radius: var(--efy_radius);
        border: var(--efy_border);
    }

    #theme img {
        border-radius: var(--efy_radius) var(--efy_radius) 0 0;
        border-bottom: var(--efy_border);
    }

    #theme>div {
        display: flex;
        flex-direction: column;
        width: 100%;
    }

    #info {
        display: flex;
        flex-direction: column;
        padding: 10rem 15rem 15rem 15rem;
    }

    #theme .actions {
        display: flex;
        gap: var(--efy_gap0);
        border-top: var(--efy_border);
        padding: var(--efy_gap0);
    }

    #theme .actions :is(a, button, i:not([efy_icon=check])) {
        margin: 0;
}

    #approve {
        float: right;
        margin-right: 0 0 10rem 0;
    }

    [efy_icon=arrow_down]:before {
        position: relative;
        margin: 0 8rem 0 0;
        display: inline-block;
        -webkit-background-clip: text!important;
        background-clip: text!important;
        background: var(--efy_color);
        color: transparent;
    }
</style>