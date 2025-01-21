<script>
  /** @typedef {import('./_components/types/invoicePayment').InvoicePayment} InvoicePayment */

  import { onMount } from 'svelte';
  import { UserPaymentResult } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';

  const InvoiceStatusPending  = `pending`;
	const InvoiceStatusSuccess  = `success`;
	const InvoiceStatusFailed   = `failed`;
	const InvoiceStatusCanceled = `canceled`;

  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);

  let invoicePayment = /** @type {InvoicePayment} */ ({});

  async function fetchPaymentResult() {
    await UserPaymentResult(
    /** @type {import('./jsApi.GEN').UserPaymentResultIn | any} */ ({
      invoiceNumber: urlParams.get('invoiceNumber'),
    }), /** @type {import('./jsApi.GEN').UserPaymentResultCallback | any} */ async (/** @type {any} */ res) => {
      if (res.error) {
        notifier.showError(res.error || 'failed to fetch payment result');
        return;
      }

      invoicePayment = res.invoicePayment;
      if (invoicePayment.status === InvoiceStatusPending) {
        setTimeout(() => {
          fetchPaymentResult();
        }, 1300);
      }

      return;
    });
  }

  onMount(async () => {
    await fetchPaymentResult();
  });
</script>

<div class="root-layout">
  <section class="payment-container">
    <div class="box shadow">
      <img
        src="/assets/icons/payment-success.png"
        alt="Payment Success"
        class="icon"
      />
      <div class="content">
        <h1>Payment Success</h1>
        {#if invoicePayment.status === InvoiceStatusPending}
          <p>Please wait...</p>
        {/if}

        {#if invoicePayment.status !== InvoiceStatusPending}
          <p>Invoice status : {invoicePayment.status}</p>
        {/if}

        <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Tempore officia sed illo neque quae non blanditiis, voluptate impedit, numquam quis totam iste provident perspiciatis sapiente, et expedita. Doloribus, animi consequuntur.</p>
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
    width: 500px;
    height: fit-content;
    background-color: #FFF;
    border-radius: 10px;
    padding: 20px 50px 30px;
  }

  .payment-container .box .icon {
    width: 150px;
    height: auto;
  }

  .payment-container .box .content {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 10px;
  }

  .payment-container .box .content h1 {
    margin: 0;
    font-size: var(--font-xl);
  }

  .payment-container .box .content p {
    margin: 0;
    text-align: center;
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