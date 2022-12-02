<section>
{#if themes}
{#each themes as theme}
    <ThemePreview theme={theme} />
{/each}
{:else}
<div class="spin"></div>
{/if}
</section>
<script type="ts">
	import ThemePreview from "../components/ThemePreview.svelte";
	import fetchJson from "../lib/fetchjs";

    let themes: any[];

    const fetchThemes = async () => {
        themes = await fetchJson("/themes?showUnapproved=true", {});
    }

    fetchThemes();
</script>
<style>
    section {
        display: grid;
        grid-template-columns: auto auto auto;
        flex-wrap: wrap;
    }
</style>