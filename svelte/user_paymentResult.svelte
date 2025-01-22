<script>
  /** @typedef {import('./_components/types/invoicePayment').InvoicePayment} InvoicePayment */

  import { onMount } from 'svelte';
  import { UserPaymentResult } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier';
  import { isoTime } from './_components/xformatter';

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
      {#if invoicePayment.status === InvoiceStatusPending}
        <p>Please wait...</p>
      {/if}
      {#if invoicePayment.status === InvoiceStatusSuccess}
        <img
          src="/assets/icons/payment-success.png"
          alt="Payment Success"
          class="icon"
        />
      {/if}
      {#if invoicePayment.status === InvoiceStatusFailed}
        <img
          src="/assets/icons/payment-failed.png"
          alt="Payment Failed"
          class="icon"
        />
      {/if}
      <div class="content">
        {#if invoicePayment.status === InvoiceStatusSuccess}
          <h1>Payment Success</h1>
          <p>Your payment has been successfully processed. Thank you for completing the transaction.</p>
        {/if}
        {#if invoicePayment.status === InvoiceStatusFailed}
          <h1>Payment Failed</h1>
          <p>Unfortunately, your payment could not be processed. Please review the details below to understand the issue and take corrective action.</p>
        {/if}
        <div class="details">
          <table>
            <tr>
              <th>Invoice Number</th>
              <td class="sep">:</td>
              <td>{invoicePayment.invoiceNumber}</td>
            </tr>
            <tr>
              <th>Amount</th>
              <td class="sep">:</td>
              <td>{invoicePayment.amount}</td>
            </tr>
            <tr>
              <th>Currency</th>
              <td class="sep">:</td>
              <td>{invoicePayment.currency}</td>
            </tr>
            <tr>
              <th>Payment Method</th>
              <td class="sep">:</td>
              <td>{invoicePayment.paymentMethod}</td>
            </tr>
            <tr>
              <th>Payment Status</th>
              <td class="sep">:</td>
              <td>{invoicePayment.status}</td>
            </tr>
            <tr>
              <th>Payment Date</th>
              <td class="sep">:</td>
              <td>{isoTime(invoicePayment.createdAt)}</td>
            </tr>
          </table>
        </div>
      </div>
      <div class="footer">
        <button class="back" on:click={() => window.location.href = '/'}>
          Back to home
        </button>
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

  .payment-container .box .content .details {
    width: 100%;
  }

  .payment-container .box .content .details table {
    width: 100%;
    text-align: left;
  }

  .payment-container .box .content .details table .sep {
    padding: 0 10px;
  }

  .payment-container .box .content h1 {
    margin: 0;
    font-size: var(--font-xl);
  }

  .payment-container .box .content p {
    margin: 0;
    text-align: center;
  }

  .payment-container .box .footer {
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    width: 100%;
  }

  .payment-container .box .footer .back {
    width: 100%;
    background-color: var(--gray-002);
    border: none;
    padding: 5px 15px;
    border-radius: 999px;
    cursor: pointer;
    width: fit-content;
  }

  .payment-container .box .footer .back:hover {
    background-color: var(--gray-003);
    border-radius: 999px;
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