<section>
{#each themes as theme}
    <ThemePreview theme={theme} />
{/each}
</section>
<script type="ts">
	import ThemePreview from "../components/ThemePreview.svelte";
	import fetchJson from "../lib/fetchjs";

    let themes: any[] = [];

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