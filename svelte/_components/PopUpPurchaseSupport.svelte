<script>
  export let isShow = false;
	export let isPurchasing = false;
  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;
	export let OnSubmit = async () => {};
</script>

<div class={`popup-container ${isShow ? 'show' : ''}`}>
  <div class="popup">
    <header>
      <img
				src="/assets/img/banner-support-plus.png"
				alt="Banner Support Plus"
				class="banner"
			/>
    </header>
    <div class="content">
			<div class="description">
				<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sapiente ad quidem magni incidunt accusantium illum amet repellat ea distinctio voluptatum! Similique molestias vel quos in omnis aut vero debitis velit.</p>
			</div>
			<div class="buttons">
				<button class="btn buy" on:click={OnSubmit} disabled={isPurchasing}>
					{#if isPurchasing}
						<span>Purchasing...</span>
					{/if}
					{#if !isPurchasing}
						<span>Purchase</span>
					{/if}
				</button>
				<button class="btn skip" on:click={Hide} disabled={isPurchasing}>
					<span>Skip for now</span>
				</button>
			</div>
    </div>
  </div>
</div>

<style>
  .popup-container {
    display: none;
		position: fixed;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		z-index: 2000;
		background-color: rgba(0 0 0 / 40%);
		backdrop-filter: blur(1px);
		justify-content: center;
		padding: 50px;
    overflow: auto;
	}

  .popup-container.show {
    display: flex;
  }

	.popup-container .popup {
		border-radius: 14px;
		background-color: #FFF;
		height: fit-content;
		width: 500px;
		display: flex;
		flex-direction: column;
		overflow: hidden;
	}

  .popup-container .popup header {
		display: flex;
		flex-direction: row;
		align-items: center;
	}

	.popup-container .popup header img {
		width: 100%;
		height: auto;
	}

	.popup-container .popup .content {
		display: flex;
		flex-direction: column;
		gap: 20px;
		padding: 20px;
	}

	.popup-container .popup .content .description p {
		margin: 0;
	}

	.popup-container .popup .content .buttons {
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.popup-container .popup .content .buttons .btn {
		width: 100%;
		padding: 10px 20px;
		border-radius: 9999px;
		border: none;
		cursor: pointer;
		font-weight: 600;
	}

	.popup-container .popup .content .buttons .btn:disabled {
		cursor: not-allowed;
	}

	.popup-container .popup .content .buttons .btn.buy {
		background-color: var(--sky-006);
		color: #FFF;
	}

	.popup-container .popup .content .buttons .btn.buy:hover {
		background-color: var(--sky-005);
	}

	.popup-container .popup .content .buttons .btn.buy:disabled {
		background-color: var(--sky-005);
	}

	.popup-container .popup .content .buttons .btn.skip {
		background-color: transparent;
		color: var(--gray-008);
	}

	.popup-container .popup .content .buttons .btn.skip:hover {
		background-color: var(--gray-002);
	}
</style>