<footer>
    <span>Licensed under <a class="license" href="https://github.com/Bnyro/efydb/blob/main/LICENSE.md">AGPL-3.0</a></span>
    <a href="https://github.com/Bnyro/efydb">
        <img src="/github.svg" alt="GitHub">
    </a>
</footer>

<style>
    footer {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 5rem 10rem 5rem 15rem;
        margin: 0;
    }
    footer a {
        line-height: 1;
    }
    footer .license {
        background: transparent;
        -webkit-text-fill-color: var(--efy_text_trans);
        text-decoration: underline;
    }
    footer a img {
        height: 40rem;
        width: 40rem;
        padding: 10rem;
        filter: invert(0.5);
    }
    footer a:hover img {
        background: var(--efy_color);
        filter: invert(0);
    }
</style>