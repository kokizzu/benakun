<script>
  /** @typedef {import('../_components/types/access').Access} Access */
  /** @typedef {import('../_components/types/user').User} User*/

  import SideMenu from '../_components/partials/SideMenu.svelte';
  import Navbar from '../_components/partials/Navbar.svelte';
  import Footer from '../_components/partials/Footer.svelte';
  import PopUpPurchaseSupport from '../_components/PopUpPurchaseSupport.svelte';
  import { IsShrinkMenu } from '../_components/uiState';
  import { onMount } from 'svelte';
  import { UserPurchaseSupport } from '../jsApi.GEN';
  import { notifier } from '../_components/xNotifier';
  import { IsUnixTimeExpired } from '../_components/xHelper';
  
  let segments = /** @type Access */ ({/* segments */});
  let user = /** @type User */ ({/* user */});

  let viewportWidth = window.innerWidth;
  let isPopUpReady = false;
  let isPopUpShowing = false;

  function setPopUpIntervalTime() {
    const timestamp = new Date().getTime();
    localStorage.setItem('popup-date', timestamp.toString());
  }

  onMount(() => {
    isPopUpReady = true;
    if (viewportWidth > 768) {
      const isShrink = localStorage.getItem('IsShrinkMenu');
      IsShrinkMenu.set(JSON.parse(isShrink));
    } else {
      localStorage.setItem('IsShrinkMenu', 'true');
      IsShrinkMenu.set(true);
      document.addEventListener('click', (e) => {
        const elm = /** @type HTMLElement */ (e.target);
        if (elm.tagName == 'a') {
          console.log('TRIGGERED')
          $IsShrinkMenu = true;
        }
      })
    }

    // TODO: fix it to show once a day if user not yet buy support+
    if (IsUnixTimeExpired(user.supportExpiredAt)) {
      const storedPopupDate = localStorage.getItem('popup-date');
      if (storedPopupDate) {
        const savedTime = parseInt(storedPopupDate, 10);
        const currentTime = new Date().getTime();
        const threeDaysInMs = 3 * 24 * 60 * 60 * 1000;
        if (((currentTime / 1000)-savedTime) > threeDaysInMs) {
          setTimeout(() => {
            popUpPurchaseSupport.Show();
          }, 3000);
        }
      } else {
        localStorage.setItem('popup-date', new Date().getTime().toString());
        setTimeout(() => {
          popUpPurchaseSupport.Show();
          setPopUpIntervalTime();
        }, 3000);
      }
    }
  });

  let popUpPurchaseSupport;
  let isPurchasing = false;

  /**
	 * @type {'yearly' | 'monthly' | 'quarterly' | any}
	 */
  let supportDuration = 'monthly';

  async function purchaseSupport() {
		isPurchasing = true;
		await UserPurchaseSupport(
    /** @type {import('../jsApi.GEN').UserPurchaseSupportIn | any} */ ({
      state: 'paymentRequest',
      supportDuration: supportDuration
    }), /** @type {import('../jsApi.GEN').UserPurchaseSupportCallback} */ async (res) => {
			console.log(res); // @ts-ignore
			if (res.error) { // @ts-ignore
				notifier.showError(res.error || 'failed to purchase support+');
				return;
			}
      isPurchasing = false;

      const paymentURL = res.paymentResponse.response.payment.url;

      if (!paymentURL) {
        notifier.showError('failed to purchase support+');
        return;
      }
      
      isPopUpShowing = false;
      
      // @ts-ignore
      loadJokulCheckout(paymentURL);

			return;
		});

		popUpPurchaseSupport.Hide();
  }
</script>

{#if isPopUpReady}
  <PopUpPurchaseSupport
    bind:this={popUpPurchaseSupport}
    bind:isPurchasing
    bind:supportDuration
    bind:isShow={isPopUpShowing}
    OnSubmit={purchaseSupport}
  />
{/if}

<div class="root-layout">
  <div class="root-container">
    <SideMenu access={segments} />
    <div class="root-content {$IsShrinkMenu ? 'shrink' : 'expand'}">
      <Navbar {user} />
      <div class="content">
        <main><slot /></main>
        <Footer />
      </div>
    </div>
  </div>
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
  }

  .root-layout .root-container {
    height: 100%;
		width: 100%;
		display: flex;
  }

  .root-layout .root-container .root-content {
		display: flex;
		flex-direction: column;
		-webkit-box-orient: vertical;
		-webkit-box-direction: normal;
		overflow-y: auto;
		min-height: calc(100vh - var(--navbar-height));
		transition: 0.3s;
		width: 100%;
  }

  .root-layout .root-container .root-content.shrink {
    margin-left: var(--sidemenu-width-sm);
  }

  .root-layout .root-container .root-content.expand {
		margin-left: var(--sidemenu-width);
  }

  .root-layout .root-container .root-content .content {
    width: 100%;
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    gap: 30px;
    min-height: calc(100vh - var(--navbar-height));
    max-height: fit-content;
    overflow: inherit;
  }

  @media only screen and (max-width : 768px) {
    .root-layout .root-container .root-content.shrink {
      margin-left: 0;
    }

    .root-layout .root-container .root-content.expand {
      margin-left: 0;
    }

    .root-layout .root-container .root-content .content {
      padding: 15px;
    }
  }
</style>