<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="theme" on:click={onClick}>
    <img src={theme.screenshot} alt={theme.name}>
    <span class="theme-info">
        <h4>{theme.title}</h4>
        <h6>{theme.username}</h6>
    </span>
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

    h6 {
        opacity: 0.7;
        margin-top: -7rem;
    }
</style>