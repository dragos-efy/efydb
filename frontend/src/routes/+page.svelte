<svelte:head>
    <title>Themes</title>
</svelte:head>

<section efy_select>
{#if role > 0}
<input type="checkbox" id="showUnapproved" on:change={onToggleShowUnapproved} bind:value={showUnapproved} />
<label for="showUnapproved">Unapproved</label>
{/if}
<span>
    {#if token}
    <a class="new-theme efy_color_trans" efy_card href="/new">
        <i efy_icon="plus"></i>
        <p>New Theme</p>
    </a>
    {/if}
    {#if themes}
    {#each themes as theme}
    <ThemePreview theme={theme} />
    {/each}
    {:else}
<div class="spin"></div>
{/if}
</span>
</section>
<script type="ts">
	// import { getRole } from "$lib/token";
	import { onMount } from "svelte";
	import ThemePreview from "../components/ThemePreview.svelte";
	import fetchJson from "../lib/fetchjs";

	import { getRole, getToken } from "$lib/token";

    let themes: any[];
    let role: number = 0;
    let showUnapproved = false;

    let token: string | null;

    const fetchThemes = async () => {
        token = getToken();

        role = getRole();
        themes = await fetchJson(`/themes?unapproved=${showUnapproved}`, {});
    }

    const onToggleShowUnapproved = () => {
        showUnapproved = !showUnapproved;
        fetchThemes();
    }

    onMount(() => {
        fetchThemes();
    })
</script>
<style>
    label {
        align-self: flex-end;
        margin: 0 0 15rem 0;
    }

    section {
        justify-content: start;
    }

    section span {
        display: grid;
        grid-template-columns: repeat(auto-fill,minmax(240rem,1fr));
        flex-wrap: wrap;
        gap: var(--efy_gap);
        width: 100%;
    }

    .new-theme {
        display: flex;
        flex-direction: column;
        gap: 15rem;
        align-items: center;
        place-content: center;
        -webkit-background-clip: text, padding-box!important;
        background: var(--efy_color), var(--efy_color_trans);
        min-height: 215rem;
    }
    .new-theme * {
        margin: 0;
    }
    [efy_icon=plus]:before {
        position: relative;
        font-size: 40rem;
        margin: 0;
        transform: rotate(45deg);
        display: inline-block;
        -webkit-background-clip: text!important;
        background-clip: text!important;
        background: var(--efy_color);
        color: transparent;
    }
</style>