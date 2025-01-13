<script>
	import { Icon } from '../node_modules/svelte-icons-pack/dist';
	import { AiFillTags } from '../node_modules/svelte-icons-pack/dist/ai';

  export let isShow = true;
	export let isPurchasing = false;
  export const Show = () => isShow = true;
  export const Hide = () => isShow = false;

	/**
	 * @type {'yearly' | 'monthly' | 'quarterly'}
	 */
	let duration = 'quarterly';

	/**
	 * @description Submit purchase
	 * @type {Function}
	 * @param duration {'yearly' | 'monthly' | 'quarterly'}
	 * @returns {Promise<void>}
	 */
	export let OnSubmit = async (duration) => {
		console.log('OnSubmit duration:', duration);
	}
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
			<div class="options">
				<div class="support-duration">
					<label for="monthly" class="shadow" class:monthly={duration === 'monthly'}>
						<header>
							<div class="icon monthly">
								<Icon
									src={AiFillTags}
									color="var(--orange-004)"
								/>
							</div>
							<span>Monthly</span>
						</header>
						<h2>IDR 50000</h2>
					</label>
					<input
						id="monthly"
						type="radio"
						name="duration"
						value="monthly"
						bind:group={duration}
					/>
				</div>
				<div class="support-duration">
					<label for="quarterly" class="shadow" class:quarterly={duration === 'quarterly'}>
						<header>
							<div class="icon quarterly">
								<Icon
									src={AiFillTags}
									color="var(--blue-004)"
								/>
							</div>
							<span>Quarterly</span>
						</header>
						<h2>IDR 120000</h2>
					</label>
					<input
						id="quarterly"
						type="radio"
						name="duration"
						value="quarterly"
						bind:group={duration}
					/>
				</div>
				<div class="support-duration">
					<label for="yearly" class="shadow" class:yearly={duration === 'yearly'}>
						<header>
							<div class="icon yearly">
								<Icon
									src={AiFillTags}
									color="var(--violet-004)"
								/>
							</div>
							<span>Yearly</span>
						</header>
						<h2>IDR 450000</h2>
					</label>
					<input
						id="yearly"
						type="radio"
						name="duration"
						value="yearly"
						bind:group={duration}
					/>
				</div>
			</div>
			<div class="buttons">
				<button class="btn buy" on:click={()=>OnSubmit(duration)} disabled={isPurchasing}>
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
		background-color: var(--gray-001);
		height: fit-content;
		width: 600px;
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

	.popup-container .popup .content .options {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		align-items: stretch;
		justify-items: center;
		gap: 10px;
	}

	.popup-container .popup .content .options .support-duration {
		display: flex;
		height: 120px;
		width: 100%;
		position: relative;
	}

	.popup-container .popup .content .options .support-duration label {
		display: flex;
    flex-direction: column;
    justify-content: space-between;
    border: 2px solid #FFF;
    border-radius: 10px;
    padding: 20px;
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: #FFF;
    cursor: pointer;
	}

	.popup-container .popup .content .options .support-duration label.monthly {
		border: 2px solid var(--orange-003);
	}

	.popup-container .popup .content .options .support-duration label.monthly h2 {
		color: var(--orange-004);
	}

	.popup-container .popup .content .options .support-duration label.quarterly {
		border: 2px solid var(--blue-003);
	}

	.popup-container .popup .content .options .support-duration label.quarterly h2 {
		color: var(--blue-004);
	}

	.popup-container .popup .content .options .support-duration label.yearly {
		border: 2px solid var(--violet-003);
	}

	.popup-container .popup .content .options .support-duration label.yearly h2 {
		color: var(--violet-004);
	}

	.popup-container .popup .content .options .support-duration label header {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 10px;
		margin: 0;
		font-size: 16px;
	}

	.popup-container .popup .content .options .support-duration label header .icon {
		display: flex;
    align-items: center;
    justify-content: center;
    width: fit-content;
    height: fit-content;
    padding: 5px;
    border-radius: 5px;
	}

	.popup-container .popup .content .options .support-duration label header .icon.monthly {
		background-image: var(--orange-gradient);
    color: var(--orange-006);
    border: 1px solid var(--orange-003);
	}

	.popup-container .popup .content .options .support-duration label header .icon.quarterly {
		background-image: var(--blue-gradient);
    color: var(--blue-006);
    border: 1px solid var(--blue-003);
	}

	.popup-container .popup .content .options .support-duration label header .icon.yearly {
		background-image: var(--violet-gradient);
    color: var(--violet-006);
    border: 1px solid var(--violet-003);
	}

	.popup-container .popup .content .options .support-duration input {
		display: none;
		visibility: hidden;
	}

	.popup-container .popup .content .options .support-duration h2 {
		margin: 0;
		font-size: 24px;
		font-weight: 600;
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
		text-decoration: underline;
		color: var(--sky-006);
	}
</style>