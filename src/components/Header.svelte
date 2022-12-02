<script>
	import { getToken } from "$lib/token";
</script>

<header>
    <a href="/">EfyDb</a>
    <nav>
        <ul>
            <li><a href="/">Themes</a></li>
            <li><a href="/account">Account</a></li>
            {#if getToken()}
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