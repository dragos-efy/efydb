<section>
{#if theme}
<div id="theme" style="background: url({theme.screenshot});">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
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

    let theme: any;
    let id = parseInt($page.url.searchParams.get('id')!);

    onMount(async () => {
        theme = await fetchJson(`/themes/${id}`, {});
    })
</script>

<style>
    #theme {
        height: 60vh;
        width: 60vw;
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
</style>