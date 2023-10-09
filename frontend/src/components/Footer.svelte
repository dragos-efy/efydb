<footer>
	<span
		>Licensed under <a class="license" href="https://github.com/Bnyro/efydb/blob/main/LICENSE.md"
			>AGPL-3.0</a
		></span
	>
	<a href="https://github.com/Bnyro/efydb" title="Github">
		<i efy_icon="github" />
	</a>
</footer>

<style>
footer {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 5rem 5rem 5rem 15rem;
	margin: 0;
	& a { line-height: 1 }
	& .license {
		background: transparent;
		-webkit-text-fill-color: var(--efy_text_trans);
		text-decoration: underline;
		&:hover {
			background: var(--efy_color);
			-webkit-text-fill-color: transparent;
			background-clip: text;
			-webkit-background-clip: text;
		}
	}
	& [title=Github] {
		display: flex;
		align-items: center;
		place-content: center;
		border-radius: var(--efy_radius);
		background: transparent;
		-webkit-text-fill-color: unset;
		color: var(--efy_text_trans);
		height: var(--efy_ratio_width);
		width: var(--efy_ratio_width);
		& i {margin: 0}
		&:hover {
			background: var(--efy_color);
			& i {color: var(--efy_bg)}
		}
	}
}
</style>
