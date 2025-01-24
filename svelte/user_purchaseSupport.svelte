<script>
    import { notifier } from './_components/xNotifier';
    import { UserPurchaseSupport } from './jsApi.GEN';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
	import { AiFillTags } from './node_modules/svelte-icons-pack/dist/ai';

  let isPurchasing = false;

  /**
	 * @type {'yearly' | 'monthly' | 'quarterly'}
	 */
	let supportDuration = 'quarterly';

  /**
	 * @description Submit purchase
	 * @returns {Promise<void>}
	 */
	async function submitPurchase() {
    isPurchasing = true;
		await UserPurchaseSupport(
    /** @type {import('./jsApi.GEN').UserPurchaseSupportIn | any} */ ({
      state: 'paymentRequest',
      supportDuration: supportDuration
    }), /** @type {import('./jsApi.GEN').UserPurchaseSupportCallback} */ async (res) => {
      isPurchasing = false;
      
			console.log(res); // @ts-ignore
			if (res.error) { // @ts-ignore
				notifier.showError(res.error || 'failed to purchase support+');
				return;
			}

      const paymentURL = res.paymentResponse.response.payment.url;

      if (!paymentURL) {
        notifier.showError('failed to purchase support+');
        return;
      }

      
      // @ts-ignore
      loadJokulCheckout(paymentURL);

			return;
    });
	}
</script>

<div class="root-layout">
  <section class="payment-container">
    <div class="box shadow">
      <div class="img-container">
        <img
          src="/assets/img/banner-support-plus.png"
          alt="Payment Success"
          class="icon"
        />
      </div>
      <div class="content">
        <div class="description">
          <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Sapiente ad quidem magni incidunt accusantium illum amet repellat ea distinctio voluptatum! Similique molestias vel quos in omnis aut vero debitis velit.</p>
        </div>
        <div class="options">
          <div class="support-duration">
            <label for="monthly" class="shadow" class:monthly={supportDuration === 'monthly'}>
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
              bind:group={supportDuration}
            />
          </div>
          <div class="support-duration">
            <label for="quarterly" class="shadow" class:quarterly={supportDuration === 'quarterly'}>
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
              bind:group={supportDuration}
            />
          </div>
          <div class="support-duration">
            <label for="yearly" class="shadow" class:yearly={supportDuration === 'yearly'}>
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
              bind:group={supportDuration}
            />
          </div>
        </div>
        <div class="buttons">
          <button class="btn buy" on:click={submitPurchase} disabled={isPurchasing}>
            {#if isPurchasing}
              <span>Purchasing...</span>
            {/if}
            {#if !isPurchasing}
              <span>Purchase</span>
            {/if}
          </button>
          <button class="btn skip" on:click={() => window.location.href = '/'} disabled={isPurchasing}>
            <span>Skip for now</span>
          </button>
        </div>
      </div>
    </div>
  </section>
</div>

<style>
  .root-layout {
    display: block;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		height: 100vh;
		width: 100vw;
    color: var(--gray-009);
  }

  .payment-container {
    display: flex;
    height: 100%;
    width: 100%;
    align-items: center;
    justify-content: center;
    background-color: var(--gray-001);
  }

  .payment-container .box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 20px;
    width: 700px;
    height: fit-content;
    background-color: #FFF;
    border-radius: 10px;
    padding: 0 0 20px 0;
    overflow: hidden;
  }

  .payment-container .box .img-container {
    height: auto;
    width: 100%;
  }

  .payment-container .box .img-container .icon {
    width: 100%;
    height: auto;
  }

  .payment-container .box .content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 20px;
    padding: 20px
  }

	.payment-container .box .content .description p {
		margin: 0;
	}

	.payment-container .box .content .options {
		display: grid;
    grid-template-columns: repeat(3, 1fr);
    align-items: stretch;
    justify-items: center;
    width: 100%;
    gap: 10px;
	}

	.payment-container .box .content .options .support-duration {
    display: flex;
    flex-direction: row;
    height: 120px;
    width: 100%;
	}

	.payment-container .box .content .options .support-duration label {
		display: flex;
    flex-direction: column;
    justify-content: space-between;
    border: 2px solid #FFF;
    border-radius: 10px;
    padding: 20px;
    width: 100%;
    height: 100%;
    background-color: #FFF;
    cursor: pointer;
	}

	.payment-container .box .content .options .support-duration label.monthly {
		border: 2px solid var(--orange-003);
	}

	.payment-container .box .content .options .support-duration label.monthly h2 {
		color: var(--orange-004);
	}

	.payment-container .box .content .options .support-duration label.quarterly {
		border: 2px solid var(--blue-003);
	}

	.payment-container .box .content .options .support-duration label.quarterly h2 {
		color: var(--blue-004);
	}

	.payment-container .box .content .options .support-duration label.yearly {
		border: 2px solid var(--violet-003);
	}

	.payment-container .box .content .options .support-duration label.yearly h2 {
		color: var(--violet-004);
	}

	.payment-container .box .content .options .support-duration label header {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 10px;
		margin: 0;
		font-size: 16px;
	}

	.payment-container .box .content .options .support-duration label header .icon {
		display: flex;
    align-items: center;
    justify-content: center;
    width: fit-content;
    height: fit-content;
    padding: 5px;
    border-radius: 5px;
	}

	.payment-container .box .content .options .support-duration label header .icon.monthly {
		background-image: var(--orange-gradient);
    color: var(--orange-006);
    border: 1px solid var(--orange-003);
	}

	.payment-container .box .content .options .support-duration label header .icon.quarterly {
		background-image: var(--blue-gradient);
    color: var(--blue-006);
    border: 1px solid var(--blue-003);
	}

	.payment-container .box .content .options .support-duration label header .icon.yearly {
		background-image: var(--violet-gradient);
    color: var(--violet-006);
    border: 1px solid var(--violet-003);
	}

	.payment-container .box .content .options .support-duration input {
		display: none;
		visibility: hidden;
	}

	.payment-container .box .content .options .support-duration h2 {
		margin: 0;
		font-size: 24px;
		font-weight: 600;
	}

	.payment-container .box .content .buttons {
		display: flex;
    flex-direction: column;
    gap: 10px;
    width: 50%;
    margin-top: 20px;
	}

	.payment-container .box .content .buttons .btn {
		width: 100%;
		padding: 10px 20px;
		border-radius: 9999px;
		border: none;
		cursor: pointer;
		font-weight: 600;
	}

	.payment-container .box .content .buttons .btn:disabled {
		cursor: not-allowed;
	}

	.payment-container .box .content .buttons .btn.buy {
		background-color: var(--sky-006);
		color: #FFF;
	}

	.payment-container .box .content .buttons .btn.buy:hover {
		background-color: var(--sky-005);
	}

	.payment-container .box .content .buttons .btn.buy:disabled {
		background-color: var(--sky-005);
	}

	.payment-container .box .content .buttons .btn.skip {
		background-color: transparent;
		color: var(--gray-008);
	}

	.payment-container .box .content .buttons .btn.skip:hover {
		text-decoration: underline;
		color: var(--sky-006);
	}

  @media only screen and (max-width : 768px) {
    .payment-container {
      align-items: baseline;
    }

    .payment-container .box {
      width: 100%;
      margin: 20px 10px;
    }
  }
</style>