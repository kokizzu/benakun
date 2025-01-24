<script>
	let segment2 = /** @type {string} */ (window.location.pathname.split('/').pop() || '');
</script>

<div class="profile-submenu">
	<nav>
		<a href="/user/profile" class:active={segment2==='profile'}>
			<span class="title">Profile</span>
			<span class="block"></span>
		</a>
    <a href="/user/profileSessions" class:active={segment2==='profileSessions'}>
			<span class="title">Sessions</span>
			<span class="block"></span>
		</a>
    <a href="/user/profileInvoices" class:active={segment2==='profileInvoices'}>
			<span class="title">Invoices</span>
			<span class="block"></span>
		</a>
	</nav>
</div>

<style>
	.profile-submenu {
		display: flex;
		flex-direction: row;
		width: 100%;
		height: fit-content;
    margin-bottom: 20px;
	}

	.profile-submenu nav {
		display: flex;
		flex-direction: row;
		width: 100%;
		height: fit-content;
		gap: 10px;
	}

	.profile-submenu nav a {
		height: fit-content;
		display: flex;
		flex-direction: column;
		text-decoration: none;
		color: var(--gray-008);
	}

	.profile-submenu nav a:hover {
		background-color: var(--gray-001);
	}

	.profile-submenu nav a.active {
		color: var(--blue-005);
	}

	.profile-submenu nav a .title {
		flex-grow: 1;
		padding: 15px 10px;
		display: flex;
		align-items: center;
	}

	.profile-submenu nav a .block {
		height: 3px;
		background-color: transparent;
		width: 100%;
	}

	.profile-submenu nav a:hover .block {
		background-color: var(--gray-005);
	}
	.profile-submenu nav a.active .block {
		background-color: var(--blue-005);
	}
</style>
