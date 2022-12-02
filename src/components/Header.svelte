<script lang="ts">
	import { getToken } from "$lib/token";
	import { onMount } from "svelte";

    let token: string | null;

    onMount(() => {
        token = getToken();
    })
</script>

<header>
    <a href="/">EfyDb</a>
    <nav>
        <ul>
            <li><a href="/">Themes</a></li>
            <li><a href="/account">Account</a></li>
            {#if token}
            <li><a href="/new">New</a></li>
            {/if}
        </ul>
    </nav>
</header>

<style>
    header {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
</style>