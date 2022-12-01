<section>
{#each themes as theme}
    <div class="theme">
        <img src={theme.screenshot} alt={theme.name}>
        <h4>{theme.title}</h4>
        <h5>{theme.username}</h5>
    </div>
{/each}
</section>
<script type="ts">
	import fetchJson from "../lib/fetchjs";

    let themes: any[] = [];

    const fetchThemes = async () => {
        themes = await fetchJson("/themes?showUnapproved=true", {});
        console.log(themes)
    }

    fetchThemes();
</script>
<style>
    section {
        display: grid;
        grid-template-columns: auto auto auto;
        flex-wrap: wrap;
    }

    .theme {
        background: var(--efy_bg1);
        border-radius: var(--efy_radius);
        padding: var(--efy_padding);
        margin: var(--efy_padding);
    }

    .theme img {
        aspect-ratio: 16 / 9;
    }
</style>