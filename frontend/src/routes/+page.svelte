<svelte:head>
    <title>Themes</title>
</svelte:head>

<section>
{#if role > 0}
<label for="showUnapproved">Show unapproved
<input type="checkbox" id="showUnapproved" on:change={onToggleShowUnapproved} bind:value={showUnapproved}>
</label>
{/if}
<span>
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
	import { getRole } from "$lib/token";
	import { onMount } from "svelte";
	import ThemePreview from "../components/ThemePreview.svelte";
	import fetchJson from "../lib/fetchjs";

    let themes: any[];
    let role: number = 0;
    let showUnapproved = false;

    const fetchThemes = async () => {
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
        margin: 10rem;
    }

    section {
        justify-content: start;
    }

    section span {
        display: grid;
        grid-template-columns: auto auto auto;
        flex-wrap: wrap;
    }
</style>