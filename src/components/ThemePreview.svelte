<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="theme" on:click={onClick}>
    <img src={theme.screenshot} alt={theme.name}>
    <h4>{theme.title}</h4>
    <h5>{theme.username}</h5>
</div>

<script lang="ts">
    export let theme: any;

    const onClick = () => {
        window.location.href = "/theme?id=" + theme.ID;
    }
</script>

<style>
    .theme {
        background: var(--efy_bg1);
        border-radius: var(--efy_radius);
        padding: var(--efy_padding);
        margin: var(--efy_padding);
    }

    .theme:hover {
        cursor: pointer;
    }

    .theme img {
        aspect-ratio: 16 / 9;
    }
</style>